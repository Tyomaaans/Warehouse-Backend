package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"Backend-Warehouse/config"
	"Backend-Warehouse/infrastructure/postgres"
	"Backend-Warehouse/infrastructure/redis"
	"Backend-Warehouse/interface/handler"
	"Backend-Warehouse/interface/httpclient"
	"Backend-Warehouse/interface/middleware"
	"Backend-Warehouse/interface/router"
	"Backend-Warehouse/interface/usecase"
	"Backend-Warehouse/validator"
)

func main() {
	cfg := config.NewConfig()
	validate := validator.New()
	redisClient := redis.NewRedisClient(cfg.REDISaddr, cfg.REDISpassword)
	db, _ := postgres.NewPostgresDB(cfg.DSN)

	empRepo    := postgres.NewEmployeeRepository(db)
    itemRepo   := postgres.NewItemRepository(db)
    spplrRepo  := postgres.NewSupplierRepository(db)
    goodsRepo  := postgres.NewGoodsRepository(db)
	jwtService := redis.NewJWTService(cfg.JWTSecretKey, cfg.JWTExpiry, cfg.JWTRefreshExpiry, redisClient)
    signatureRepo := postgres.NewSignatureRepository(db)

    fileClient := httpclient.NewPythonClient(cfg.FileServiceURL)

	empUseCase   := usecase.NewEmployeeUseCase(empRepo, jwtService, validate)
    itemUseCase  := usecase.NewItemUseCase(itemRepo, validate)
    spplrUseCase := usecase.NewSupplierUsecase(spplrRepo, validate)
    goodsUseCase := usecase.NewGoodsUseCase(goodsRepo, validate)
    signatureUseCase := usecase.NewSignatureUseCase(signatureRepo, empRepo, spplrRepo, fileClient)

	empHandler     := handler.NewEmployeeHandler(empUseCase, fileClient, cfg.JWTRefreshExpiry)
    itemHandler    := handler.NewItemHandler(itemUseCase)
    spplrHandler   := handler.NewSupplierHandler(spplrUseCase)
    goodsHandler   := handler.NewGoodsHandler(goodsUseCase)
    signatureHandler := handler.NewSignatureHandler(signatureUseCase)
    authMiddleware := middleware.NewAuthMiddleware(jwtService)

	r := router.NewRouter(empHandler, itemHandler, spplrHandler, goodsHandler, signatureHandler, authMiddleware)

	srv := &http.Server{
        Addr:         ":" + cfg.APPport,
        Handler:      r,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  30 * time.Second,
    }

    go func() {
        log.Printf("Server running on :%s", cfg.APPport)
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %s\n", err)
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    log.Println("Shutting down server...")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("Server forced to shutdown: %v", err)
    }

    if err := redisClient.Close(); err != nil {
        log.Printf("Redis close error: %v", err)
    }

    sqlDB, err := db.DB()
    if err == nil {
        if err := sqlDB.Close(); err != nil {
            log.Printf("DB close error: %v", err)
        }
    }

    log.Println("Server exited")

}
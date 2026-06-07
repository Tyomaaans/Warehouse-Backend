package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"Backend-Warehouse/domain/entity"
	"Backend-Warehouse/interface/handler"
	"Backend-Warehouse/interface/middleware"
)

func NewRouter(
	empHandler     *handler.EmployeeHandler,
	itemHandler    *handler.ItemHandler,
	spplrHandler   *handler.SupplierHandler,
	goodsHandler   *handler.GoodsHandler,
	signatureHandler *handler.SignatureHandler,
	authMiddleware *middleware.AuthMiddleware,

) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:4173"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", empHandler.LoginEmployee)
			auth.POST("/refresh", empHandler.RefreshToken)
			auth.POST("/logout", authMiddleware.Authenticate(), empHandler.LogoutEmployee)
		}

		employee := v1.Group("/employees", authMiddleware.Authenticate())
		{
			// Semua role yang sudah login bisa lihat & update profile sendiri
			employee.GET("/me", empHandler.GetEmployee)
			employee.PATCH("/me/:employeeID", empHandler.UpdateEmployeeForUser)
			employee.PATCH("/me/:employeeID/avatar", empHandler.UpdateAvatarForUser)

			// Hanya Admin yang bisa create & manage employee lain
			employee.POST("/admin", authMiddleware.RequireRole(entity.RoleWHAdmin), empHandler.RegisterEmployee)
			employee.GET("/admin", authMiddleware.RequireRole(entity.RoleWHAdmin), empHandler.GetEmployees)
			employee.PATCH("/admin/:employeeID", authMiddleware.RequireRole(entity.RoleWHAdmin), empHandler.UpdateEmployeeForAdmin)
			employee.PATCH("/admin/:employeeID/avatar", authMiddleware.RequireRole(entity.RoleWHAdmin), empHandler.UpdateAvatarForAdmin)

			employee.POST("/signature", signatureHandler.RegisterEmployeeSignature)
			employee.GET("/me/signature", signatureHandler.GetEmployeeActiveSignature)
			employee.GET("/me/signature/history", signatureHandler.GetEmployeeSignatureHistory)
		}

		item := v1.Group("/items", authMiddleware.Authenticate())
		{
			// Hanya Admin yang bisa create & update item
			item.POST("", authMiddleware.RequireRole(entity.RoleWHAdmin), itemHandler.AddItem)
			item.PATCH("/:itemID", authMiddleware.RequireRole(entity.RoleWHAdmin), itemHandler.UpdateItem)

			// Admin, Spv, Staff, Purchasing bisa lihat item (operasional butuh lihat item)
			item.GET("", authMiddleware.RequireRole(entity.RoleWHAdmin, entity.RoleWHSpv, entity.RoleWHStaff, entity.RolePurchasing), itemHandler.GetItems)
			item.GET("/:itemID", authMiddleware.RequireRole(entity.RoleWHAdmin, entity.RoleWHSpv, entity.RoleWHStaff, entity.RolePurchasing), itemHandler.GetItem)
		}

		category := v1.Group("/categories", authMiddleware.Authenticate())
		{
			// Hanya Admin yang bisa create & update category
			category.POST("", authMiddleware.RequireRole(entity.RoleWHAdmin), itemHandler.AddCategory)
			category.PATCH("/:categoryID", authMiddleware.RequireRole(entity.RoleWHAdmin), itemHandler.UpdateCategory)

			// Admin, Spv, Staff, Purchasing bisa lihat category
			category.GET("", authMiddleware.RequireRole(entity.RoleWHAdmin, entity.RoleWHSpv, entity.RoleWHStaff, entity.RolePurchasing), itemHandler.GetCategories)
			category.GET("/:categoryID", authMiddleware.RequireRole(entity.RoleWHAdmin, entity.RoleWHSpv, entity.RoleWHStaff, entity.RolePurchasing), itemHandler.GetCategory)
		}

		// Supplier: hanya Admin yang bisa akses (master data sensitif)
		supplier := v1.Group("/suppliers", authMiddleware.Authenticate(), authMiddleware.RequireRole(entity.RoleWHAdmin))
		{
			supplier.POST("", spplrHandler.RegisterSupplier)
			supplier.GET("", spplrHandler.GetSuppliers)
			supplier.GET("/:supplierID", spplrHandler.GetSupplier)
			supplier.PATCH("/:supplierID", spplrHandler.UpdateSupplier)
		}

		goods := v1.Group("/goods-receipts", authMiddleware.Authenticate())
		{
			// Admin, Staff, Purchasing bisa create goods receipt
			goods.POST("", authMiddleware.RequireRole(entity.RoleWHAdmin, entity.RoleWHStaff, entity.RolePurchasing), goodsHandler.AddGoodsReceipt)

			// Admin, Spv, Purchasing, Management bisa lihat goods receipt
			goods.GET("", authMiddleware.RequireRole(entity.RoleWHAdmin, entity.RoleWHSpv, entity.RolePurchasing, entity.RoleMangement), goodsHandler.GetGoodsReceipts)
			goods.GET("/:goodsReceiptID", authMiddleware.RequireRole(entity.RoleWHAdmin, entity.RoleWHSpv, entity.RolePurchasing, entity.RoleMangement), goodsHandler.GetGoodsReceipt)

			// Hanya Admin & Spv yang bisa approve
			goods.PATCH("/:goodsReceiptID/approve", authMiddleware.RequireRole(entity.RoleWHAdmin, entity.RoleWHSpv), goodsHandler.ApproveGoodsReceipt)
		}

		issue := v1.Group("/goods-issues", authMiddleware.Authenticate())
		{
			// Admin, Staff, Purchasing bisa create goods issue
			issue.POST("", authMiddleware.RequireRole(entity.RoleWHAdmin, entity.RoleWHStaff, entity.RolePurchasing), goodsHandler.AddGoodsIssue)

			// Admin, Spv, Purchasing, Management bisa lihat goods issue
			issue.GET("", authMiddleware.RequireRole(entity.RoleWHAdmin, entity.RoleWHSpv, entity.RolePurchasing, entity.RoleMangement), goodsHandler.GetGoodsIssues)
			issue.GET("/:goodsIssueID", authMiddleware.RequireRole(entity.RoleWHAdmin, entity.RoleWHSpv, entity.RolePurchasing, entity.RoleMangement), goodsHandler.GetGoodsIssue)

			// Hanya Admin & Spv yang bisa approve
			issue.PATCH("/:goodsIssueID/approve", authMiddleware.RequireRole(entity.RoleWHAdmin, entity.RoleWHSpv), goodsHandler.ApproveGoodsIssue)
		}
	}
	
	return r
}
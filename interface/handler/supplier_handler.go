package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"Backend-Warehouse/interface/dto"
	"Backend-Warehouse/interface/usecase"
)

type SupplierHandler struct {
	spplrUC *usecase.SupplierUsecase
}

func NewSupplierHandler(spplrUC *usecase.SupplierUsecase) *SupplierHandler {
	return &SupplierHandler{
		spplrUC: spplrUC,
	}
}

func (h *SupplierHandler) RegisterSupplier(c *gin.Context) {
	var req dto.RegisterSupplierRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if err := h.spplrUC.RegisterSupplier(c.Request.Context(), req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Register Supplier Success!",
	})
}

func (h *SupplierHandler) GetSuppliers(c *gin.Context) {
	res, err := h.spplrUC.GetSuppliers(c.Request.Context())
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message":  "Get Supplier Success!",
		"response": res,
	})
}

func (h *SupplierHandler) GetSupplier(c *gin.Context) {
	var req dto.GetSupplierRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	res, err := h.spplrUC.GetSupplier(c.Request.Context(), req)
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message":  "Get Supplier Success!",
		"response": res,
	})
}

func (h *SupplierHandler) UpdateSupplier(c *gin.Context) {
	var req dto.UpdateSupplierRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if err := h.spplrUC.UpdateSupplier(c.Request.Context(), req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Update Supplier Success!",
	})
}
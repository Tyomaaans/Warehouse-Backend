package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"Backend-Warehouse/interface/dto"
	"Backend-Warehouse/interface/usecase"
)

type GoodsHandler struct {
	goodsUC *usecase.GoodsUseCase
}

func NewGoodsHandler(goodsUC *usecase.GoodsUseCase) *GoodsHandler {
	return &GoodsHandler{
		goodsUC: goodsUC,
	}
}

func (h *GoodsHandler) AddGoodsReceipt(c *gin.Context) {
	employeeID := c.GetString("employeeID")

	var req dto.AddGoodsReceiptRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	req.EmployeeID = employeeID

	if err := h.goodsUC.AddGoodsReceipt(c.Request.Context(), req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Add Goods Receipt Success!",
	})
}

func (h *GoodsHandler) GetGoodsReceipts(c *gin.Context) {
	res, err := h.goodsUC.GetGoodsReceipts(c.Request.Context())
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Get Goods Receipts Success!",
		"response": res,
	})
}

func (h *GoodsHandler) GetGoodsReceipt(c *gin.Context) {
	var req dto.GetOrApproveGoodsReceiptOrIssueRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	res, err := h.goodsUC.GetGoodsReceipt(c.Request.Context(), req)
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Get Goods Receipts Success!",
		"response": res,
	})
}

func (h *GoodsHandler) ApproveGoodsReceipt(c *gin.Context) {
	var req dto.GetOrApproveGoodsReceiptOrIssueRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if err := h.goodsUC.ApproveGoodsReceipt(c.Request.Context(), req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Approve Goods Receipt Success!",
	})
}

func (h *GoodsHandler) AddGoodsIssue(c *gin.Context) {
	employeeID := c.GetString("employeeID")

	var req dto.AddGoodsIssueRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	req.EmployeeID = employeeID

	if err := h.goodsUC.AddGoodsIssue(c.Request.Context(), req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Add Goods Issue Success!",
	})
}

func (h *GoodsHandler) GetGoodsIssues(c *gin.Context) {
	res, err := h.goodsUC.GetGoodsIssues(c.Request.Context())
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Get Goods Issues Success!",
		"response": res,
	})
}

func (h *GoodsHandler) GetGoodsIssue(c *gin.Context) {
	var req dto.GetOrApproveGoodsReceiptOrIssueRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	res, err := h.goodsUC.GetGoodsIssue(c.Request.Context(), req)
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Get Goods Issues Success!",
		"response": res,
	})
}

func (h *GoodsHandler) ApproveGoodsIssue(c *gin.Context) {
	var req dto.GetOrApproveGoodsReceiptOrIssueRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if err := h.goodsUC.ApproveGoodsIssue(c.Request.Context(), req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Approve Goods Issue Success!",
	})
}
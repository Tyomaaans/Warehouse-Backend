package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"Backend-Warehouse/interface/dto"
	"Backend-Warehouse/interface/usecase"
)

type ItemHandler struct {
	itemUC *usecase.ItemUseCase
}

func NewItemHandler(itemUC *usecase.ItemUseCase) *ItemHandler {
	return &ItemHandler{
		itemUC: itemUC,
	}
}

func (h *ItemHandler) AddItem(c *gin.Context) {
	var req dto.AddItemRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if err := h.itemUC.AddItem(c.Request.Context(), req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Add Item Success!",
	})
}

func (h *ItemHandler) AddCategory(c *gin.Context) {
	var req dto.AddCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if err := h.itemUC.AddCategory(c.Request.Context(), req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Add Category Success!",
	})
}

func (h *ItemHandler) GetItems(c *gin.Context) {
	res, err := h.itemUC.GetItems(c.Request.Context())
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Get Items Success!",
		"response": res,
	})
}

func (h *ItemHandler) GetItem(c *gin.Context) {
	var req dto.GetItemOrCategoryByID

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	res, err := h.itemUC.GetItem(c.Request.Context(), req)
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Get Item Success!",
		"response": res,
	})
}

func (h *ItemHandler) GetCategories(c *gin.Context) {
	res, err := h.itemUC.GetCategories(c.Request.Context())
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Get Categories Success!",
		"response": res,
	})
}

func (h *ItemHandler) GetCategory(c *gin.Context) {
	var req dto.GetItemOrCategoryByID

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	res, err := h.itemUC.GetCategory(c.Request.Context(), req)
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Get Category Success!",
		"response": res,
	})
}

func (h *ItemHandler) UpdateItem(c *gin.Context) {
	var req dto.UpdateItemRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if err := h.itemUC.UpdateItem(c.Request.Context(), req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Update Item Success!",
	})
}

func (h *ItemHandler) UpdateCategory(c *gin.Context) {
	var req dto.UpdateCategoryRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if err := h.itemUC.UpdateCategory(c.Request.Context(), req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(200, gin.H{
		"message": "Update Category Success!",
	})
}

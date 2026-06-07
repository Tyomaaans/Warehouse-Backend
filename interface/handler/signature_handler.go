package handler

import (
	"Backend-Warehouse/interface/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignatureHandler struct {
	signatureUsecase *usecase.SignatureUseCase
}

func NewSignatureHandler(
	signatureUsecase *usecase.SignatureUseCase,
) *SignatureHandler {
	return &SignatureHandler{
		signatureUsecase: signatureUsecase,
	}
}

// ── Employee Signature ────────────────────────────────────────

// POST /employees/:id/signature
func (h *SignatureHandler) RegisterEmployeeSignature(c *gin.Context) {
	employeeID := c.GetString("employeeID")

	file, header, err := c.Request.FormFile("signature")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "signature file not found"})
		return
	}
	defer file.Close()

	log.Println("Uploud File Hit")

	if err := h.signatureUsecase.Register(c.Request.Context(), employeeID, file, header); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "signature registered successfully"})
}

// GET /employees/:id/signature
func (h *SignatureHandler) GetEmployeeActiveSignature(c *gin.Context) {
	employeeID := c.GetString("employeeID")

	signature, err := h.signatureUsecase.GetActive(c.Request.Context(), employeeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "active signature not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": signature})
}

// GET /employees/:id/signature/history
func (h *SignatureHandler) GetEmployeeSignatureHistory(c *gin.Context) {
	employeeID := c.GetString("employeeID")

	signatures, err := h.signatureUsecase.GetAll(c.Request.Context(), employeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": signatures})
}

// ── Supplier Signature ────────────────────────────────────────

// POST /suppliers/:id/signature
/*func (h *SignatureHandler) RegisterSupplierSignature(c *gin.Context) {
	supplierID := c.Param("id")

	file, header, err := c.Request.FormFile("signature")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "signature file not found"})
		return
	}
	defer file.Close()

	if err := h.supplierSignatureUsecase.Register(c.Request.Context(), supplierID, file, header); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "signature registered successfully"})
}

// GET /suppliers/:id/signature
func (h *SignatureHandler) GetSupplierActiveSignature(c *gin.Context) {
	supplierID := c.Param("id")

	signature, err := h.supplierSignatureUsecase.GetActive(c.Request.Context(), supplierID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "active signature not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": signature})
}

// GET /suppliers/:id/signature/history
func (h *SignatureHandler) GetSupplierSignatureHistory(c *gin.Context) {
	supplierID := c.Param("id")

	signatures, err := h.supplierSignatureUsecase.GetAll(c.Request.Context(), supplierID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": signatures})
}*/
package controllers

import (
	"inventory-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InventarisController struct {
	DB *gorm.DB
}

func (ctrl *InventarisController) GetInventory(c *gin.Context) {
	id := c.Param("id")
	var inventaris models.Inventaris
	if err := ctrl.DB.Where("id_produk = ?", id).First(&inventaris).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventaris tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, inventaris)
}

func (ctrl *InventarisController) UpdateInventory(c *gin.Context) {
	id := c.Param("id")
	var inventaris models.Inventaris
	if err := ctrl.DB.Where("id_produk = ?", id).First(&inventaris).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventaris tidak ditemukan"})
		return
	}

	var updateData struct {
		Jumlah int `json:"jumlah"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctrl.DB.Save(&inventaris)
	c.JSON(http.StatusOK, inventaris)
}

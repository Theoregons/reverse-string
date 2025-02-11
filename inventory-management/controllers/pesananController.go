package controllers

import (
	"inventory-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PesananController struct {
	DB *gorm.DB
}

func (ctrl *PesananController) CreateOrder(c *gin.Context) {
	var pesanan models.Pesanan

	if err := c.ShouldBindJSON(&pesanan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := ctrl.DB.Begin()

	var inventaris models.Inventaris
	if err := tx.Where("id_produk = ?", pesanan.ID_produk).First(&inventaris).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan di inventaris"})
		return
	}

	if inventaris.Jumlah < pesanan.Jumlah {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Stok tidak mencukupi"})
		return
	}

	inventaris.Jumlah -= pesanan.Jumlah
	if err := tx.Save(&inventaris).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui inventaris"})
		return
	}

	if err := tx.Create(&pesanan).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat pesanan"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{
		"pesanan":    pesanan,
		"inventaris": inventaris,
	})
}

func (ctrl *PesananController) GetOrders(c *gin.Context) {
	var pesanan []models.Pesanan
	ctrl.DB.Find(&pesanan)
	c.JSON(http.StatusOK, pesanan)
}

func (ctrl *PesananController) GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	var pesanan models.Pesanan
	if err := ctrl.DB.First(&pesanan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pesanan tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, pesanan)
}

func (ctrl *PesananController) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.DB.Delete(&models.Pesanan{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pesanan tidak ditemukan"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

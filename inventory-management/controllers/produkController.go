package controllers

import (
	"inventory-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProdukController struct {
	DB *gorm.DB
}

func (ctrl *ProdukController) AddProduct(c *gin.Context) {
	var produk models.Produk

	if err := c.ShouldBindJSON(&produk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.DB.Create(&produk).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat produk"})
		return
	}

	inventaris := models.Inventaris{
		ID_Produk: produk.ID,
		Jumlah:    0,
		Lokasi:    "Gudang Utama",
	}

	if err := ctrl.DB.Create(&inventaris).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat inventaris untuk produk"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"produk":     produk,
		"inventaris": inventaris,
	})
}

func (ctrl *ProdukController) GetProducts(c *gin.Context) {
	var produk []models.Produk
	ctrl.DB.Find(&produk)
	c.JSON(http.StatusOK, produk)
}

func (ctrl *ProdukController) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var produk models.Produk
	if err := ctrl.DB.First(&produk, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, produk)
}

func (ctrl *ProdukController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var produk models.Produk
	if err := ctrl.DB.First(&produk, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}
	if err := c.ShouldBindJSON(&produk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctrl.DB.Save(&produk)
	c.JSON(http.StatusOK, produk)
}

func (ctrl *ProdukController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.DB.Delete(&models.Produk{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

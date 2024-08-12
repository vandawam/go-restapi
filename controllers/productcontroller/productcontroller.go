package productcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vandawam/go-restapi/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []models.Product
	
	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func Show(c *gin.Context) {
	var  product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Data gagal di update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil di update"})
	c.JSON(http.StatusOK, gin.H{"data": product})

}

func Delete(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data gagal di delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil di hapus"})
}
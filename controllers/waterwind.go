package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iqbal13/tugas3hactiv/config"
	"github.com/iqbal13/tugas3hactiv/helpers"
	"github.com/iqbal13/tugas3hactiv/models"
)

func CreateWaterWind(c *gin.Context) {
	var waterwind models.WaterWind

	config.DB.Delete(&waterwind)
	if err := c.ShouldBindJSON(&waterwind); err != nil {
		helpers.RespondJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	waterwind.CreatedAt = time.Now()
	waterwind.UpdatedAt = time.Now()

	if err := config.DB.Create(&waterwind).Error; err != nil {
		helpers.RespondJSON(c, http.StatusInternalServerError, "Gagal Create Data", nil)
		return
	}

	helpers.RespondJSON(c, http.StatusOK, "Berhasil", waterwind)
}

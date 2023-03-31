package handler

import (
	"dans/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Login(c *gin.Context) {
	var form model.Login

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.usecase.User.Login(form)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    result,
	})
}

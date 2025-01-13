package handlers

import (
	"github.com/gin-gonic/gin"
	"sm.com/m/src/app/utils"
)

type UpdatePasswordRequest struct {
	Password    string `json:"password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

func UpdatePasswordHandler(c *gin.Context) {
	var requestBody UpdatePasswordRequest

	err := c.ShouldBindBodyWithJSON(&requestBody)

	if err != nil {
		utils.FormatAndSendRequiredFieldsError(err, c)
		return
	}
}

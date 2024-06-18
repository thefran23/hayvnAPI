package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"github.com/thefran23/hayvnAPI/models"
)

var validate = validator.New()

func Message() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var message models.Message
		if err := c.BindJSON(&message); err != nil {
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		if validationErr := validate.Struct(&message); validationErr != nil {
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		c.JSON(http.StatusCreated, ctx)
	}
}

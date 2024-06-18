package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"github.com/thefran23/hayvnAPI/models"
)

var validate = validator.New()

var messages []models.Message

func Message() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

		messages = append(messages, message)
		fmt.Println(messages)

		c.JSON(http.StatusCreated, message)
	}
}

func Scheduler() {
	for range time.Tick(1 * time.Second) {
		fmt.Println(messages)
	}

}

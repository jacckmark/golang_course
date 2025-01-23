package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rest.com/api/models"
)

func registerForAnEvent(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	err = event.Register(userID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}

func cancelRegistration(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, _ := strconv.ParseInt(context.Param("id"), 10, 64)

	event, err := models.GetByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	err = event.CancelRegistration(userID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not deregister user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Deregistered!"})
}

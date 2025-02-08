package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID."})
		return
	}

	event, err := models.GetEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not fetch event.",
		})
		return
	}

	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not fetch events.",
		})
		return
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data."})
		return
	}

	userId := context.MustGet("userId").(int64)
	event.UserID = userId
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not save event.",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"status": "Event created.", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID."})
		return
	}

	event, err := models.GetEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not fetch event.",
		})
		return
	}

	userId := context.MustGet("userId").(int64)

	if event.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update this event."})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data."})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not update event.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully."})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID."})
		return
	}

	event, err := models.GetEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not fetch event.",
		})
		return
	}

	userId := context.MustGet("userId").(int64)

	if event.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to delete this event."})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not delete event.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully."})
}
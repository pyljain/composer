package routes

import (
	"net/http"

	"composer/internal/db"
	"composer/internal/models"

	"github.com/labstack/echo/v4"
)

func RegisterChatSessionRoutes(e *echo.Echo, database *db.Db) {
	e.POST("/api/chat-sessions", createChatSession(database))
	e.GET("/api/chat-sessions", listChatSessions(database))
	e.GET("/api/chat-sessions/:id", getChatSession(database))
	e.PUT("/api/chat-sessions/:id", updateChatSession(database))
	e.DELETE("/api/chat-sessions/:id", deleteChatSession(database))
}

func createChatSession(database *db.Db) echo.HandlerFunc {
	return func(c echo.Context) error {
		session := models.ChatSession{}

		if err := database.InsertChatSession(&session); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, session)
	}
}

func listChatSessions(database *db.Db) echo.HandlerFunc {
	return func(c echo.Context) error {
		chatSessions, err := database.ListChatSessions()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, chatSessions)
	}
}

func getChatSession(database *db.Db) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		chatSession, err := database.GetChatSession(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, err)
		}

		return c.JSON(http.StatusOK, chatSession)
	}
}

func updateChatSession(database *db.Db) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		var chatSession models.ChatSession
		if err := c.Bind(&chatSession); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		chatSession.ID = id
		if err := database.UpdateChatSession(&chatSession); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, chatSession)
	}
}

func deleteChatSession(database *db.Db) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		if err := database.DeleteChatSession(id); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.NoContent(http.StatusNoContent)
	}
}

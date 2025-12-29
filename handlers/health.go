package handlers

import (
	"net/http"
	"pastebin-lite/store"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	if err := store.RDB.Ping(store.Ctx).Err(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"ok": false, "error": "redis unavailable"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}

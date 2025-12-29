package handlers

import (
	"encoding/json"
	"net/http"
	"pastebin-lite/store"
	"pastebin-lite/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPaste(c *gin.Context) {
	id := c.Param("id")

	val, err := store.RDB.Get(store.Ctx, "paste:"+id).Result()
	if err != nil {
		pasteNotFound(c)
		return
	}

	var paste store.Paste
	if err := json.Unmarshal([]byte(val), &paste); err != nil {
		pasteNotFound(c)
		return
	}

	now := utils.Now(c)

	// TTL check
	if paste.ExpiresAt != nil && now > *paste.ExpiresAt {
		pasteNotFound(c)
		return
	}

	// view limit check
	if paste.MaxViews != nil && paste.Views >= *paste.MaxViews {
		pasteNotFound(c)
		return
	}

	// on  success view increment
	paste.Views++

	data, _ := json.Marshal(paste)
	if err := store.RDB.Set(store.Ctx, "paste:"+id, data, 0).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal redis error"})
		return
	}

	// remaining views calculate
	var remainingViews *int
	if paste.MaxViews != nil {
		r := *paste.MaxViews - paste.Views
		if r < 0 {
			r = 0
		}

		remainingViews = &r
	}

	// expires at format
	var expiresAt *string
	if paste.ExpiresAt != nil {
		t := time.UnixMilli(*paste.ExpiresAt).UTC().Format(time.RFC3339Nano)
		expiresAt = &t
	}

	c.JSON(http.StatusOK, gin.H{"content": paste.Content, "remaining_views": remainingViews, "expires_at": expiresAt})
}

func pasteNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "paste not found"})
}

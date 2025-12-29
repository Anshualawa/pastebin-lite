package handlers

import (
	"encoding/json"
	"net/http"
	"pastebin-lite/store"
	"pastebin-lite/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreatePasteRequest struct {
	Content    string `json:"content"`
	TTLSeconds *int64 `json:"ttl_seconds"`
	MaxViews   *int   `json:"max_views"`
}

func CreatePaste(c *gin.Context) {
	var req CreatePasteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
		return
	}

	if req.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "content is required"})
		return
	}

	if req.TTLSeconds != nil && *req.TTLSeconds < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ttl_seconds must be >= 1"})
		return
	}

	if req.MaxViews != nil && *req.MaxViews < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "max_views must be >= 1"})
		return
	}

	now := utils.Now(c)

	var expiresAt *int64

	if req.TTLSeconds != nil {
		v := now + (*req.TTLSeconds * 1000)
		expiresAt = &v
	}

	id := uuid.New().String()

	paste := store.Paste{
		ID:        id,
		Content:   req.Content,
		CreatedAt: now,
		ExpiresAt: expiresAt,
		MaxViews:  req.MaxViews,
		Views:     0,
	}

	data, _ := json.Marshal(paste)

	if err := store.RDB.Set(store.Ctx, "paste:"+id, data, 0).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save paste"})
		return
	}

	baseURL := c.Request.Host
	schema := "http"

	if c.Request.TLS != nil {
		schema = "https"
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":  id,
		"url": schema + "://" + baseURL + "/p/" + id,
	})
}

package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"pastebin-lite/store"
	"pastebin-lite/utils"

	"github.com/gin-gonic/gin"
)

func ViewPaste(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("param id : ", id)
	// fetch from redis
	val, err := store.RDB.Get(store.Ctx, "paste:"+id).Result()
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	var paste store.Paste
	if err := json.Unmarshal([]byte(val), &paste); err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	now := utils.Now(c)

	// TTL check
	if paste.ExpiresAt != nil && now > *paste.ExpiresAt {
		c.Status(http.StatusNotFound)
		return
	}

	// view limit check
	if paste.MaxViews != nil && paste.Views >= *paste.MaxViews {
		c.Status(http.StatusNotFound)
		return
	}

	// increment views
	paste.Views++

	data, _ := json.Marshal(paste)
	if err := store.RDB.Set(store.Ctx, "paste:"+id, data, 0).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal redis view error"})
		return
	}

	// render HTML
	tmpl, err := template.ParseFiles("templates/paste.html")

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Header("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(c.Writer, paste)
}

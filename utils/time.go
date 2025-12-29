package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Now(c *gin.Context) int64 {
	if os.Getenv("TEST_MODE") == "1" {
		h := c.GetHeader("x-test-now-ms")

		if h != "" {
			if v, err := strconv.ParseInt(h, 10, 64); err == nil {
				return v
			}
		}
	}

	return time.Now().UnixMilli()
}

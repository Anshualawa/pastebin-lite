package store

type Paste struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
	ExpiresAt *int64 `json:"expires_at,omitempty"`
	MaxViews  *int   `json:"max_views,omitempty"`
	Views     int    `json:"views"`
}

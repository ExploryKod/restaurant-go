package restaurantHTTP

import (
	"embed"
)

//go:embed src/templates/*
var EmbedTemplates embed.FS

type TemplateData struct {
	Title   string `json:"title"`
	Content any    `json:"content"`
	Success string `json:"success"`
	Error   string `json:"error"`
	Token   string `json:"token"`
}

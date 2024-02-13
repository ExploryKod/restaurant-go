package restaurantHTTP

import (
	"embed"
)

//go:embed src/templates/*
var EmbedTemplates embed.FS

type TemplateData struct {
	Title   string
	Titre   string
	Content any
	Success string
	Error   string
	Token   string
}

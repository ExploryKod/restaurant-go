package restaurantHTTP

import (
	"embed"
)

//go:embed template/*
var EmbedTemplates embed.FS

type TemplateData struct {
	Titre   string
	Content any
	Success string
	Error   string
}

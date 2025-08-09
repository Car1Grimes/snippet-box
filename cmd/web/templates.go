package main

import (
	"github.com/Car1Grimes/snippet-box/internal/models"
)

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}

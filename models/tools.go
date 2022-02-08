package models

import "time"

type Tools struct {
	Name     string    `json:"name"`
	Version  float32   `json:"version"`
	User     string    `json:"user"`
	LastUsed time.Time `json:"last_used"`
}

package models

import "time"

type Tool struct {
	Id               string    `gorm:"primary_key"`
	Name             string    `json:"name"`
	Version          Version   `json:"version"`
	PreviousVersions []Version `json:"previous"`
	Description      string    `json:"description"`
	Status           string    `json:"status"`
	Values           []Value   `json:"values"`
}

type Value struct {
}

type Version struct {
	Latest        float32   `json:"latest"`
	DateOfRelease time.Time `json:"date"`
}

type User struct {
	Name     string    `json:"name"`
	LastUsed time.Time `json:"last_used"`
	Tools    []Tool    `json:"tools"`
}

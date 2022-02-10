package models

type Tool struct {
	Id      string  `gorm:"id"`
	Name    string  `json:"name"`
	Version float64 `json:"version"`
	// PreviousVersions []float64 `json:"previous"`
	Description string `json:"description"`
	Status      string `json:"status"`
	// Values           []Value   `json:"values"`
}

// type Value struct {
// }

// type Version struct {
// 	Latest        float32   `json:"latest"`
// 	DateOfRelease time.Time `json:"date"`
// }

// type User struct {
// 	Name     string    `json:"name"`
// 	LastUsed time.Time `json:"last_used"`
// 	Tools    []Tool    `json:"tools"`
// }

package models

// Todos instance Object
type Todos struct {
	BaseModel
	ID          uint   `gorm:"primaryKey" json:"todoID"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

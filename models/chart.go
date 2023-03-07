package models

// Chart chart models
type Chart struct {
	Model
	Name     string   `json:"name"`
	ChartID  uint     `json:"-"`
	Children []*Chart `json:"children" gorm:"foreignkey:ChartID;references:ID;->" walkrec:"true"`
}

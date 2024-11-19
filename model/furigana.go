package model

import "time"

type Furigana struct {
	ID        int `json:"id" gorm:"primaryKey"` // integer or string
	CreatedAt time.Time
	Q         string `json:"q"` // The text to be add with furigana
}

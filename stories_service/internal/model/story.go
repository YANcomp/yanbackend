package model

import (
	"database/sql"
	"time"
)

type Story struct {
	ID             int64        `json:"ID"`
	IsActive       bool         `json:"isActive,omitempty"`
	IsActiveMobile bool         `json:"IsActiveMobile,omitempty"`
	Preview        string       `json:"preview,omitempty"`
	Slides         []*Slide     `json:"slides,omitempty"`
	Title          string       `json:"title,omitempty"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      sql.NullTime `json:"updated_at"`
}

type Slide struct {
	ID                 int64  `json:"ID"`
	BackgroundImage    string `json:"backgroundImage,omitempty"`
	Caption            string `json:"caption,omitempty"`
	Content            string `json:"content,omitempty"`
	Delay              int64  `json:"delay,omitempty"`
	IsHideShadowBottom bool   `json:"isHideShadowBottom,omitempty"`
	TextPosition       string `json:"textPosition,omitempty"`
}

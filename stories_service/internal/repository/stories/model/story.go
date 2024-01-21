package model

import (
	"database/sql"
	"time"
)

type Story struct {
	ID             int64
	IsActive       bool
	IsActiveMobile bool
	Preview        string
	Slides         []*Slide
	Title          string
	CreatedAt      time.Time
	UpdatedAt      sql.NullTime
}

type Slide struct {
	ID                 int64
	BackgroundImage    string
	Caption            string
	Content            string
	Delay              int64
	IsHideShadowBottom bool
	TextPosition       string
	CreatedAt          time.Time
	UpdatedAt          sql.NullTime
}

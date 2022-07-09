package models

import "time"

type Item struct {
	Key      string        `json:"key"`
	Value    int64         `json:"value"`
	Duration time.Duration `json:"-"`
}

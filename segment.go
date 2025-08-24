package main

import "time"

type Segment struct {
	Key         string            `json:"key"`
	Description string            `json:"description,omitempty"`
	Percent     int               `json:"percent"`
	StartDate   *time.Time        `json:"start_date,omitempty"`
	EndDate     *time.Time        `json:"end_date,omitempty"`
	Attributes  map[string]string `json:"attributes,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

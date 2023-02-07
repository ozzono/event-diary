package models

import (
	"github.com/pkg/errors"
)

type Record struct {
	ID          uint   `gorm:"primaryKey"`
	Description string `json:"description"`
	EventTime   string `json:"event_time"`
	Reporter    string `json:"user"`
	RecordTime  string `json:"regtime"`
}

var (
	EmptyEventTimeErr = errors.New("invalid event time")
	EmptyReporterErr  = errors.New("invalid reporter")
)

func (r Record) Valid() error {
	if r.EventTime == "" {
		return errors.Wrap(EmptyEventTimeErr, "cannot be empty")
	}

	if r.Reporter == "" {
		return errors.Wrap(EmptyReporterErr, "cannot be empty")
	}
	return nil
}

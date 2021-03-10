package models

import (
	"time"
)

// errors
const (
	ErrStartDateIsRequired = ValidationErr("start_date is required query param")
	ErrEndDateToEearly     = ValidationErr("end_date must be earlier than start_date")
	ErrDateInFuture        = ValidationErr("query params dates couldn't be in future")
)

type ValidationErr string

func (e ValidationErr) Error() string {
	return string(e)
}

type Request struct {
	StartDate *time.Time `form:"start_date" time_format:"2006-01-02" time_utc:"1"`
	EndDate   *time.Time `form:"end_date" time_format:"2006-01-02" time_utc:"1"`
}

func (r *Request) Validate() error {
	if r.StartDate == nil {
		return ErrStartDateIsRequired
	}

	if r.StartDate.After(time.Now()) {
		return ErrDateInFuture
	}

	if r.EndDate != nil {
		if r.EndDate.After(time.Now()) {
			return ErrDateInFuture
		}

		if r.StartDate.After(*r.EndDate) {
			return ErrEndDateToEearly
		}
	} else {
		dateToday := time.Now()
		r.EndDate = &dateToday
	}

	return nil
}

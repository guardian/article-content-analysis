package utils

import (
	"github.com/pkg/errors"
	"time"
)

func FormatDate(date string) (string, error) {
	var res string
	layout := "2006-01-02T15:04:05Z"
	t, err := time.Parse(layout, date)

	if err != nil {
		return res, errors.Wrap(err, "invalid date format")
	}

	return string(t.Format("2006-01-02")), nil
}

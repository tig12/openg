package tiglib

import (
	"fmt"
	"time"
)

// Returns a date YYYY-MM-DD
func DateIso(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
}

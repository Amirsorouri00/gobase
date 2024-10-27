package persian

import (
	"fmt"
	"time"
	_ "time/tzdata"

	"github.com/aslrousta/persian"
	ptime "github.com/yaa110/go-persian-calendar"
)

// Date is a formatted data-time.
type Date struct {
	DayOfWeek string `json:"day_of_week"`
	Date      string `json:"date"`
	Time      string `json:"time"`
}

// DateOf converts a time.Time to Date.
func DateOf(t time.Time) Date {
	zone, _ := time.LoadLocation("Asia/Tehran")
	local := ptime.New(t.In(zone))
	return Date{
		DayOfWeek: local.Format("E"),
		Date:      persian.ToPersianDigits(local.Format("d MMM yyyy")),
		Time:      persian.ToPersianDigits(local.Format("HH:mm")),
	}
}

// String concatenates the formatted date and time.
func (d Date) String() string {
	return fmt.Sprintf("%s %s", d.Date, d.Time)
}

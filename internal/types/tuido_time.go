package types

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
)

type TuidoTime time.Time

type DateRange struct {
	StartDate TuidoTime `json:"startDate"`
	EndDate   TuidoTime `json:"endDate"`
}

func (t *TuidoTime) UnmarshalJSON(data []byte) error {
	var timeStr string
	if err := json.Unmarshal(data, &timeStr); err != nil {
		return err
	}

	ts, err := time.Parse("2006-01-02T15:04:05-0700", timeStr)
	if err != nil {
		return fmt.Errorf("invalid time format: %s", timeStr)
	}

	*t = TuidoTime(ts)
	return nil
}

func (t TuidoTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format("2006-01-02T15:04:05-0700"))
}

func (t TuidoTime) ToMSFormat() string {
	return time.Time(t).Format("2006-01-02T15:04:05.000-0700")
}

func (t TuidoTime) String() string {
	return time.Time(t).Format("Monday 2006-01-02 15:04:05")
}

func (t TuidoTime) Humanize() string {
	return humanize.Time(time.Time(t))
}

func YYYYmmddToTuidoTime(s string) (TuidoTime, error) {
	temp, err := time.Parse("2006-01-02", s)

	if err != nil {
		return TuidoTime{}, err
	}

	return TuidoTime(temp), nil

}

package utils 

import (
	"time"
	"encoding/json"
)

const customTimeFormat = "2006-01-02 15:04:05" 

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	
	s := string(b)
	t, err := time.Parse(`"`+customTimeFormat+`"`, s)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	formatted := ct.Format(customTimeFormat)
	return json.Marshal(formatted)
}

func StringPtr(s string) *string {
	return &s
}
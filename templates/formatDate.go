package templates

import (
	"fmt"
	"strings"
	"time"
)

const dateFormat = "2006-01-02"

type Formatdate struct {
	time.Time
}

func (f Formatdate) MarshalJSON() ([]byte, error) {
	if f.Time.IsZero() {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", f.Format(dateFormat))), nil
}
func (f *Formatdate) UnmarshalJSON(v []byte) error {

	if f.Time.IsZero() {
		return nil
	}

	time, err := time.Parse(dateFormat, strings.ReplaceAll(string(v), "\"", ""))

	if err != nil {
		panic(err)
	}

	f.Time = time

	return nil
}

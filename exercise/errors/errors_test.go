package timeparse

import (
	"testing"
)

func TestErrorCheck(t *testing.T) {
	time1 := "12:43:12"
	parsed, err := Convert(time1)
	if err != nil {
		t.Errorf("error occured")
	}

	if parsed.Hour != 12 {
		t.Errorf("Not correct data type")
	}

}

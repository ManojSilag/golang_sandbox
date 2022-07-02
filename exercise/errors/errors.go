//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Parser struct {
	Hour, minute, second int
}

type Time string

func Convert(p string) (Parser, error) {
	timearray := strings.Split(p, ":")
	// fmt.Println(timearray)
	if len(timearray) < 3 {
		return Parser{}, errors.New(fmt.Sprintf("Time not in Proper formant"))
	}
	hour, err := strconv.Atoi(timearray[0])
	if err != nil {
		return Parser{}, errors.New(fmt.Sprintf("Time not in Proper formant"))
	}
	min, err1 := strconv.Atoi(timearray[1])
	if err1 != nil {
		return Parser{}, errors.New(fmt.Sprintf("Time not in Proper formant"))
	}
	sec, err2 := strconv.Atoi(timearray[2])
	if err2 != nil {
		return Parser{}, errors.New(fmt.Sprintf("Time not in Proper formant"))
	}
	P := Parser{Hour: hour, minute: min, second: sec}
	return P, nil
}

func timeparse() {
	time1 := "14:07:23"
	Parsed, err := Convert(time1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(Parsed)
	}
}

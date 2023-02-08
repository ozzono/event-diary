package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	TimeFormatToday = "2006-01-02"
	TimeFormatNow   = "15:04:05"
)

var (
	TimeStampFormat = fmt.Sprintf("%s %s", TimeFormatToday, TimeFormatNow)
)

func Today() string {
	return strings.Split(time.Now().Format(TimeStampFormat), " ")[0]
}

func Custom(input string) (string, error) {
	matched, err := regexp.Match(`\d{2}:\d{2}`, []byte(input))
	if err != nil {
		return "", errors.Wrap(err, "regexp.Match")
	}
	if !matched {
		return "", errors.New("invalid input; must match format hh:mm")
	}
	return fmt.Sprintf("%s %s:%s", Today(), input, "00"), nil
}

func Now() string {
	return time.Now().Format(TimeStampFormat)
}

func TimeParse(input string) (time.Time, error) {
	output, err := time.Parse(TimeStampFormat, input)
	if err != nil {
		return time.Time{}, errors.Wrap(err, "time.Parse")
	}
	return output, nil
}

func RInt(i int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(i)
}

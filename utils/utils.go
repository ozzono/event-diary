package utils

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	TimeFormatToday = "2006-01-02"
	TimeFormatNow   = "15:04:05"
)

func Today() string {
	return strings.Split(time.Now().Format(fmt.Sprintf("%s %s", TimeFormatToday, TimeFormatNow)), " ")[0]
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
	return time.Now().Format(fmt.Sprintf("%s %s", TimeFormatToday, TimeFormatNow))
}

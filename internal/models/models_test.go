package models

import (
	"event-diary/chart"
	"event-diary/utils"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	maxCount = 10
)

func TestReport(t *testing.T) {
	records := []*Record{}

	for i := 0; i < 24; i++ {
		records = append(records, newTestRecord(i)...)
	}
	t.Logf("len records %d", len(records))
	r, err := ReportData(records)
	assert.NoError(t, err)

	chart.Chart(r.ToOptData(), r.HourArr())
}

func newTestRecord(i int) []*Record {
	output := []*Record{}
	for j := 0; j < utils.RInt(maxCount); j++ {
		now := strings.Split(time.Now().Format(utils.TimeFormatNow), ":")
		now[0] = fmt.Sprint(i)
		output = append(output, &Record{
			EventTime: fmt.Sprintf("%s %s", utils.Today(), strings.Join(now, ":")),
		})
	}
	fmt.Printf("%d new records at %02dh\n", len(output), i)
	return output
}

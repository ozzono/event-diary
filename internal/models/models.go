package models

import (
	"event-diary/utils"
	"fmt"

	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/pkg/errors"
)

type Record struct {
	ID          uint   `gorm:"primaryKey"`
	Description string `json:"description"`
	EventTime   string `json:"event_time"`
	Reporter    string `json:"user"`
	RecordTime  string `json:"regtime"`
}

type Report struct {
	Hour  int
	Count int
}

type Reports []Report

func (r Reports) HourArr() []string {
	output := []string{}
	for i := range r {
		output = append(output, fmt.Sprintf("%02d", r[i].Hour))
	}
	return output
}

func (r Reports) ToOptData() []opts.BarData {
	output := []opts.BarData{}
	for i := range r {
		output = append(output, opts.BarData{Value: r[i].Count})
	}
	return output
}

func (reports Reports) SortReportsByHour() {
	for i := 1; i < len(reports); i++ {
		key := reports[i]
		j := i - 1
		for j >= 0 && reports[j].Hour > key.Hour {
			reports[j+1] = reports[j]
			j = j - 1
		}
		reports[j+1] = key
	}
}

var (
	EmptyEventTimeErr = errors.New("invalid event time")
	EmptyReporterErr  = errors.New("invalid reporter")
)

func (r Record) Valid() error {
	if r.EventTime == "" {
		return errors.Wrap(EmptyEventTimeErr, "cannot be empty")
	}

	if r.Reporter == "" {
		return errors.Wrap(EmptyReporterErr, "cannot be empty")
	}
	return nil
}

func ReportData(records []*Record) (Reports, error) {
	reportMap := map[int]int{}
	for i := 0; i < 24; i++ {
		reportMap[i] = 0
	}
	for i := range records {
		eventTime, err := utils.TimeParse(records[i].EventTime)
		if err != nil {
			return nil, errors.Wrap(err, "utils.TimeParse")
		}
		reportMap[eventTime.Hour()]++
	}

	r := map2arr(reportMap)
	var rs Reports = r
	rs.SortReportsByHour()
	return rs, nil
}

func map2arr(input map[int]int) []Report {
	output := []Report{}
	for key := range input {
		output = append(output, Report{
			Hour:  key,
			Count: input[key],
		})
	}
	return output
}

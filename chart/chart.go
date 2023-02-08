package chart

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func BarChart(data []opts.BarData, axis []string) {
	// create a new bar instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Resumo de convulsões",
		Subtitle: "Visualização por hora do dia",
	}))
	// Put data into instance
	bar.SetXAxis(axis).
		AddSeries("Hora do dia", data)
	// Where the magic happens
	f, _ := os.Create("bar.html")
	bar.Render(f)
}

func FunnelChart(data []opts.FunnelData, axis []string) {
	// create a new funnel instance
	funnel := charts.NewFunnel()
	// set some global options like Title/Legend/ToolTip or anything else
	funnel.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Resumo de convulsões",
		Subtitle: "Visualização por hora do dia",
	}))
	// Put data into instance
	funnel.AddSeries("Hora do dia", data)
	// Where the magic happens
	f, _ := os.Create("bar.html")
	funnel.Render(f)
}

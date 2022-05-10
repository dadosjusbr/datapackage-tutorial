package main

import (
	"log"
	"math"

	"github.com/ahmetb/go-linq/v3"
	"github.com/frictionlessdata/datapackage-go/datapackage"
	"github.com/frictionlessdata/tableschema-go/csv"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

const (
	pkgURL = "https://dadosjusbr.org/download/datapackage/tjal/tjal-2021-12.zip"
)

type remuneration struct {
	Value    float64 `tableheader:"valor"`
	Nature   string  `tableheader:"natureza"`
	Category string  `tableheader:"categoria"`
}

func main() {
	pkg, err := datapackage.Load(pkgURL)
	if err != nil {
		log.Fatalf("Error loading datapackage (%s):%v", pkgURL, err)
	}

	var remunerations []remuneration
	err = pkg.GetResource("remuneracao").Cast(&remunerations, csv.LoadHeaders())
	if err != nil {
		log.Fatalf("Error casting datapackage (%s):%v", pkgURL, err)
	}

	groups := linq.From(remunerations).WhereT(func(r remuneration) bool {
		return r.Nature != "D" && r.Value != 0 // removing discounts and remunerations with zero values.
	}).GroupByT(func(r remuneration) interface{} {
		return r.Category
	}, func(r remuneration) interface{} {
		return r.Value
	})

	xaxis := []string{}
	values := []plotter.Values{}
	for _, r := range groups.Results() {
		g := r.(linq.Group)
		xaxis = append(xaxis, g.Key.(string))
		sum := float64(0)
		for _, v := range g.Group {
			sum += v.(float64)
		}
		values = append(values, plotter.Values{sum})
	}

	p := plot.New()
	p.Title.Text = "Remunerações por categoria"
	p.X.Tick.Label.Rotation = math.Pi / 2

	p.NominalX(xaxis...)
	for i, v := range values {
		bar, err := plotter.NewBarChart(v, 15)
		if err != nil {
			panic(err)
		}
		bar.LineStyle.Width = vg.Length(0)
		bar.Color = plotutil.Color(i)
		bar.Offset = vg.Points(float64(i * 15))
		p.Add(bar)
	}

	if err := p.Save(3*vg.Inch, 3*vg.Inch, "bar.png"); err != nil {
		panic(err)
	}
}

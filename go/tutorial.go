package main

import (
	"fmt"
	"log"

	"github.com/frictionlessdata/datapackage-go/datapackage"
	"github.com/frictionlessdata/tableschema-go/csv"
	"gonum.org/v1/gonum/stat"
)

const (
	pkgURL = "https://dadosjusbr.org/download/datapackage/tjal/tjal-2021-12.zip"
)

func main() {
	pkg, err := datapackage.Load(pkgURL)
	if err != nil {
		log.Fatalf("Error loading datapackage (%s):%v", pkgURL, err)
	}

	var remunerations []float64
	err = pkg.GetResource("remuneracao").CastColumn("valor", &remunerations, csv.LoadHeaders())
	if err != nil {
		log.Fatalf("Error casting datapackage (%s):%v", pkgURL, err)
	}

	fmt.Println(stat.MeanStdDev(remunerations, nil))
}

package main

import (
	"fmt"
	"log"

	"github.com/frictionlessdata/datapackage-go/datapackage"
	"github.com/frictionlessdata/tableschema-go/csv"
	"gonum.org/v1/gonum/floats"
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

	fmt.Printf("O total de salários do TJAL em 12/2022 é %.2f\n", floats.Sum(remunerations))
}

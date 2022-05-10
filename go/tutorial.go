package main

import (
	"fmt"
	"log"

	"github.com/frictionlessdata/datapackage-go/datapackage"
	"github.com/frictionlessdata/tableschema-go/csv"
	"gonum.org/v1/gonum/floats"
	"github.com/ahmetb/go-linq/v3"
)

const (
	pkgURL = "https://dadosjusbr.org/download/datapackage/tjal/tjal-2021-12.zip"
)

type remuneration struct {
	Value float64 `tableheader:"valor"`
	Nature string `tableheader:"natureza"`
	Category string `tableheader:"categoria"`
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

	fmt.Printf(remunerations)
}

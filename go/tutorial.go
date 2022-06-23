package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/ahmetb/go-linq/v3"
	"github.com/frictionlessdata/datapackage-go/datapackage"
	"github.com/frictionlessdata/tableschema-go/csv"
)

const (
	pkgURL               = "https://dadosjusbr.org/download/datapackage/tjsc/tjsc-2021-11.zip"
	paycheckResName      = "contra_cheque"
	remunerationsResName = "remuneracao"
)

type remuneration struct {
	PaycheckID string  `tableheader:"id_contra_cheque"`
	Value      float64 `tableheader:"valor"`
	Nature     string  `tableheader:"natureza"`
	Category   string  `tableheader:"categoria"`
}

type paycheck struct {
	ID            string `tableheader:"id_contra_cheque"`
	Name          string `tableheader:"nome"`
	Remunerations []remuneration
}

func main() {
	fmt.Printf("Loading datapckage %s... ", pkgURL)
	pkg, err := datapackage.Load(pkgURL)
	if err != nil {
		log.Fatalf("Error loading datapackage (%s):%v", pkgURL, err)
	}
	fmt.Println("Done.")

	fmt.Printf("Loading paychecks table... ")
	var paycheckTable []paycheck
	if err = pkg.GetResource(paycheckResName).Cast(&paycheckTable, csv.LoadHeaders()); err != nil {
		log.Fatalf("Error loading paychecks (%s):%v", pkgURL, err)
	}
	fmt.Printf("%d paychecks loaded.\n", len(paycheckTable))

	fmt.Printf("Loading remunerations table... ")
	var remunerationsTable []remuneration
	if err = pkg.GetResource(remunerationsResName).Cast(&remunerationsTable, csv.LoadHeaders()); err != nil {
		log.Fatalf("Error loading remunerations (%s):%v", pkgURL, err)
	}
	fmt.Printf("%d remunerations loaded.\n", len(remunerationsTable))

	fmt.Printf("Taking debits out of the remunerations Table... ")
	var remunerations []remuneration
	linq.From(remunerationsTable).
		WhereT(func(r remuneration) bool {
			return r.Nature == "R"
		}).ToSlice(&remunerations)
	totalRem := linq.From(remunerations).
		SelectT(func(r remuneration) float64 {
			return r.Value
		}).
		SumFloats()
	fmt.Printf("%d remunerations selected. Totalling: R$ %.2f. Average: R$ %.2f\n", len(remunerations), totalRem, totalRem/float64(len(paycheckTable)))

	fmt.Printf("Joining paychecks and remunerations... ")
	var paychecks []paycheck
	linq.From(paycheckTable).
		GroupJoinT(linq.From(remunerations),
			func(p paycheck) string { return p.ID },
			func(r remuneration) string { return r.PaycheckID },
			func(p paycheck, r []remuneration) paycheck {
				p.Remunerations = append(p.Remunerations, r...)
				return p
			}).
		ToSlice(&paychecks)
	fmt.Println("Done ")

	fmt.Printf("Searching for Joana's paycheck... ")
	selectedPaycheck := linq.From(paychecks).FirstWithT(func(p paycheck) bool {
		return strings.HasPrefix(p.Name, "JOANA")
	}).(paycheck)
	selectedRem := linq.From(selectedPaycheck.Remunerations).
		SelectT(func(r remuneration) float64 {
			return r.Value
		}).SumFloats()
	fmt.Printf("Found Joana's paycheck. ID:%s. %d remunerations found, totalling: R$ %.2f. \n", selectedPaycheck.ID, len(selectedPaycheck.Remunerations), selectedRem)
}

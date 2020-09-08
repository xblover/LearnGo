package main

import (
	"encoding/csv"
	"log"
	"os"
)

func generateMatrixCSVFile(m [][]string) {

	f, err := os.Create("text.csv")
	if err != nil {
		log.Println(err.Error())
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(f)
	w.WriteAll(m)
	w.Flush()

}

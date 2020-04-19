package main

import (
	"encoding/csv"
	"os"
)

func main() {
	records := [][]string{
		{"First_name", "Last_name", "Number"},
		{"Ayanami", "Rei", "1"},
		{"Shikinami", "Asuka_Langley", "2"},
	}

	writer := csv.NewWriter(os.Stdout)
	for _, record := range records {
		writer.Write(record)
	}

	writer.Flush()
}

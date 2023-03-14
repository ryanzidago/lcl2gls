package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]

	// Open CSV file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	input_records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Print out the CSV data
	for _, input_record := range input_records {
		fmt.Println(input_record)
	}

	for _, input_record := range input_records {
		fmt.Println("Buchungstag", input_record[0],
			"Valutadatum", input_record[0],
			"Name Zahlungsbeteligter", "",
			"Verwendungszweck", input_record[4],
			"Betrag", input_record[1])
	}

	// Write CSV file
	output_file, err := os.Create("gls.csv")
	if err != nil {
		panic(err)
	}
	defer output_file.Close()

	writer := csv.NewWriter(output_file)

	// Reset the file pointer to the beginning of the file
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		// bezeichnung, iban auftrag, bic aufrag, bankname, buchungstag, valutadatum, name zahlungsbeteiligetr, iban zahlung, bic, buchungstext, verwendungszwekc, betrag, waehrung, saldo nach, bemerkung, kategorie, steuer, glaeubiger id, mandatsreferenz, kategorie
		data := []string{"", "", "", "", record[0], record[0], "", "", "", "", record[4], record[1]}

		err = writer.Write(data)
		if err != nil {
			panic(err)
		}
	}

	writer.Flush()
}

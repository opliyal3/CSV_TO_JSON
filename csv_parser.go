package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

type AccountInfo struct {
	ID          string
	FirstName   string
	LastName    string
	Email       string
	Description string
	Role        string
	Phone       string
}

func readCsv(filePath string) []AccountInfo {
	var table []AccountInfo
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	// skip first line
	if _, err := r.Read(); err != nil {
		log.Fatal(err)
	}

	csvData, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range csvData {
		var account AccountInfo
		account.ID = row[0]
		account.FirstName = row[1]
		account.LastName = row[2]
		account.Email = row[3]
		account.Description = row[4]
		account.Role = row[5]
		account.Phone = row[6]
		table = append(table, account)
	}

	return table
}

type Content []AccountInfo

func toJson(c Content) {
	jsonData, err := json.Marshal(c)
	//jsonData, err := json.MarshalIndent(records,"","  ") // convert to json

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}

// command line example
// go gun csvReader.go -csv <file_path>
func main() {
	path := flag.String("csv", "csv_file", "covert csv to json")
	flag.Parse()
	if *path == "" {
		log.Fatal("please give csv file")
	}

	records := readCsv(*path)
	toJson(records)
}

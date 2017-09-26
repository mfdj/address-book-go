package main

import (
	"./model"
	"./storage"
	"database/sql"
	"fmt"
	"encoding/csv"
	"log"
	"os"
)

func main() {
	db := storage.InitDb("goldie:gopher@/address_book_go")
	defer db.Close()

	addresses := fetchAddresses(db)

	// Make concurrent:
	// go func() {
	// }()

	fmt.Printf("Found %d addresses\n", len(addresses))
	for _, address := range addresses {
		fmt.Println(address.City)
	}

	// concurrent read then interleave results

	fetchCsv("./people.csv")
	fetchCsv("./addresses.csv")
}

func fetchCsv(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data := csv.NewReader(f)

	records, err := data.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i, row := range records {
		if i > 0 {
			fmt.Print(row)
		}
	}
}

func fetchAddresses(db *sql.DB) []*model.Address {
	rows, err := db.Query(`
		SELECT address.id, person_id, first, last, street, city, state, zip 
		FROM address
		INNER JOIN person
		ON address.person_id = person.id
	`)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	addresses := make([]*model.Address, 0)

	for rows.Next() {
		adrs := new(model.Address)
		err := rows.Scan(&adrs.Id, &adrs.PersonId, &adrs.First, &adrs.Last, &adrs.Street, &adrs.City, &adrs.State, &adrs.Zip)
		if err != nil {
			panic(err.Error())
		}
		addresses = append(addresses, adrs)
	}	

	return addresses
}

package main

import (
	"./model"
	"./storage"
	"database/sql"
	"fmt"
	"strconv"
)

func main() {
	db := storage.InitDb("goldie:gopher@/address_book_go")
	defer db.Close()

	addresses := fetchAddresses(db)
	length := len(addresses)

	fmt.Println("Addresses: " + strconv.Itoa(length))

	for i := 0; i < length; i++ {
		address := addresses[i]
		fmt.Println(address.City)
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

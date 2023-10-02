package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Address   []*Address
	Remark    string
}

func main() {
	// init struct data
	pa := &Address{"private", "ChangChun", "China"}
	wa := &Address{"work", "Boom", "China"}
	vc := &VCard{"Kol", "Mikkelson", []*Address{pa, wa}, "none"}
	// json format
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format:%s", js)
	// write data to jsonfile
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	// After opening file,the file should be closed
	defer file.Close()
	// encode file
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	// handle err
	if err != nil {
		log.Println("Error in encoding json")
	}

}

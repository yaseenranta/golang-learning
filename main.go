package main

import (
	"encoding/json"
	"fmt"
	"time"
)

/**
JSON Package Notes
	json:"-" remove field from json
	omitempty remove field when its empty
	 MarshalJSON implement formating field
**/

type Address struct {
	Addressline1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zipcode"`
	Country      string `json:"country,omitempty"`
}

type Person struct {
	Name       string     `json:"name"`
	Age        int        `json:"-"`
	Profession string     `json:"profession"`
	Address    Address    `json:"address"`
	CreatedAt  FormatDate `json:"created_at"`
	DeletedAt  FormatDate `json:"deleted_at,omitempty"`
}

const dateFormat = "2006-01-02"

type FormatDate struct {
	time.Time
}

func (f FormatDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", f.Format(dateFormat))), nil
}

func main() {
	t := time.Now()
	p := Person{
		Name:       "Muhammad Yaseen",
		Age:        29,
		Profession: "Developer",
		Address: Address{
			Addressline1: "house # 1111",
			AddressLine2: "area 311/b",
			City:         "karachi",
			State:        "sindh",
			ZipCode:      "33223",
			// Country:      "PK",
		},
		CreatedAt: FormatDate{t},
	}

	person, err := json.Marshal(p)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(person))
}

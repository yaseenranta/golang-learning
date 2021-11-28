package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
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
	if f.Time.IsZero() {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", f.Format(dateFormat))), nil
}
func (f *FormatDate) UnmarshalJSON(v []byte) error {

	if f.Time.IsZero() {
		return nil
	}

	time, err := time.Parse(dateFormat, strings.ReplaceAll(string(v), "\"", ""))

	if err != nil {
		panic(err)
	}

	f.Time = time

	return nil
}

func MarshalPractice(p *Person) {

	person, err := json.Marshal(p)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(person))
}

func UnmarshalPractice(p *Person) {

	jsonBytes, err := os.ReadFile("out.json")

	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(jsonBytes, p); err != nil {
		panic(err)
	}
	fmt.Println("Unmarshal func")
	fmt.Printf("%+v\n", p)

	// reqBytes := new(bytes.Buffer)

	// json.NewEncoder(reqBytes).Encode(p)
	// fmt.Println(reqBytes.String())
}

func DecoderPractice(p *Person) {
	f, err := os.Open("out.json")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	fmt.Println(p)

	err = json.NewDecoder(f).Decode(p)

	if err != nil {
		panic(err)
	}

	fmt.Println("Decode func")
	fmt.Printf("%+v\n", p)

}

func EncoderPractice(p *Person) {
	file, err := os.OpenFile("out2.json", os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	if err = json.NewEncoder(file).Encode(p); err != nil {
		panic(err)
	}

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

	MarshalPractice(&p)
	UnmarshalPractice(&p)
	DecoderPractice(&p)
	EncoderPractice(&p)
}

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/yaseenranta/golang-learning/templates"
)

func MarshalPractice(p *templates.Person) {

	person, err := json.Marshal(p)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(person))
}

func UnmarshalPractice(p *templates.Person) {

	jsonBytes, err := os.ReadFile("./json/out.json")

	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(jsonBytes, p); err != nil {
		panic(err)
	}
	fmt.Println("Unmarshal func")
	fmt.Printf("%+v\n", p)

}

func DecoderPractice(p *templates.Person) {
	f, err := os.Open("./json/out.json")

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

func EncoderPractice(p *templates.Person) {
	file, err := os.OpenFile("./json/out2.json", os.O_CREATE|os.O_WRONLY, 0600)

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
	p := templates.Person{
		Name:       "Muhammad Yaseen",
		Age:        29,
		Profession: "Developer",
		Address: templates.Address{
			Addressline1: "house # 1111",
			AddressLine2: "area 311/b",
			City:         "karachi",
			State:        "sindh",
			ZipCode:      "33223",
			// Country:      "PK",
		},
		CreatedAt: templates.Formatdate{t},
	}

	MarshalPractice(&p)
	UnmarshalPractice(&p)
	DecoderPractice(&p)
	EncoderPractice(&p)
}

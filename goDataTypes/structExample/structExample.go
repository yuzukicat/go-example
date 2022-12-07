package structExample

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID        int
	FirstName string `json:"name"`
	LastName  string
	Address   string `json:"address,omitempty"`
}

type Employee struct {
	// Struct embedding
	Information Person
	Manager     int
}

type Contractor struct {
	Person
	CompanyID int
}

func StructExample() {
	var john Person
	fmt.Println(john)

	person1001 := Person{1001, "John", "Doe", "Doe's street"}
	fmt.Println(person1001)

	var persion1002 Person
	persion1002.ID = 1002
	fmt.Println(persion1002.ID)

	// you can use the & operator to yield a pointer to the struct
	persion1002Copy := &person1001
	persion1002Copy.FirstName = "David"
	fmt.Println(person1001)

	var employee2001 Employee
	employee2001.Information.FirstName = "Blob"

	contractor3001 := Contractor{
		Person: Person{
			ID:        3001,
			FirstName: "Bill",
		},
	}
	contractor3001.LastName = "Gate"
	fmt.Println(contractor3001)

	// Encode and decode structs with JSON
	employeeSlice := []Employee{
		{
			Information: Person{
				ID: 1003,
			},
		},
		{
			Information: Person{
				ID: 1004,
			},
		},
	}

	emplployeeJson, _ := json.Marshal(employeeSlice)
	fmt.Printf("%s\n", emplployeeJson)

	var emplployeeJsonDeco []Employee
	json.Unmarshal(emplployeeJson, &emplployeeJsonDeco)
	fmt.Printf("%v", emplployeeJsonDeco)
}

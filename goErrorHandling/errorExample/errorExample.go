package main

import (
	"errors"
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func GetInformation(id int) (*Employee, error) {
	employee, err := ApiCallEmployee(1000)
	if err != nil {
		return nil, fmt.Errorf("get information error: %v", err)
	}
	return employee, nil
}

func GetInformationWithRetry(id int) (*Employee, error) {
	for retryTimes := 0; retryTimes < 3; retryTimes++ {
		employee, err := ApiCallEmployee(1000)
		if err == nil {
			return employee, nil
		}
		fmt.Println("server is not responding, retrying...")
		time.Sleep(time.Second * 2)
	}
	//To-do: log error instead of print it to console
	return nil, fmt.Errorf("server failed to response")
}

func GetInformationWithReusableError(id int) (*Employee, error) {
	employee, err := ApiCallEmployee(1000)
	if errors.Is(err, ErrorNotFound) {
		return nil, fmt.Errorf("get information error:ErrorNotFound: %v", err)
	}
	return employee, nil
}

// Create reusable errors
var ErrorNotFound = errors.New("Employee is not found")

func ApiCallEmployee1001(id int) (*Employee, error) {
	if id != 1001 {
		return nil, ErrorNotFound
	}
	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}

func ApiCallEmployee(id int) (*Employee, error) {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}

func main() {
	employee, err := GetInformation(1001)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(employee)
	}
}

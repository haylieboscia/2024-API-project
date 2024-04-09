package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Departments []Departments `json:"departments`
}

type Departments struct {
	ID   int    `json:"departmentId"`
	Name string `json:"displayname"`
}

func main() {

	response, err := http.Get("https://collectionapi.metmuseum.org/public/collection/v1/departments")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(responseData))

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	for _, department := range responseObject.Departments {
		fmt.Println("Department ID: ", department.ID)
		fmt.Println("Department Name: ", department.Name)

	}
}

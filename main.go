package main

import (
	"bytes"
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

// *notes: varname datatype

func PrettyPrint(str []byte) (string, error) {
	var prettyJson bytes.Buffer
	if err := json.Indent(&prettyJson, str, "", "    "); err != nil {
		return "", err
	}
	return prettyJson.String(), nil

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

	// Printing the return json
	//fmt.Println(string(responseData))

	// this is creating a new object called responseObject of type Response( struct above)
	var responseObject Response

	// by unmarshelling the JSON puts all the info into one block (ex: {dataWanted1 datawanted2})
	json.Unmarshal(responseData, &responseObject)

	// pretty prints out data
	res, err := PrettyPrint(responseData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	// data in orginal format
	fmt.Println(responseObject)

	//
	/*for _, department := range responseObject.Departments {
		fmt.Println("Department ID: ", department.ID)
		fmt.Println("Department Name: ", department.Name)

	}*/
}

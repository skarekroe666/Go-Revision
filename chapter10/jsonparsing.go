package chapter10

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

type Users struct {
	Users []User
}

func JsonParsing() {
	jsonfile, err := os.Open("chapter10/users.json")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println("successfuly opened the json file")
	defer jsonfile.Close()

	byteVal, _ := io.ReadAll(jsonfile)

	var u Users
	listUsers(byteVal, u)
}

func listUsers(b []byte, u Users) {
	json.Unmarshal(b, &u)

	for _, i := range u.Users {
		fmt.Println("User Name: " + i.Name)
		fmt.Println("User Type: " + i.Type)
		fmt.Println("User Age: " + strconv.Itoa(i.Age))
		fmt.Println("Social handle: " + i.Social.Facebook)
		fmt.Println()
	}
}

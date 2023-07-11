// save the file as main.go
// and run the "go run main.go" command in the terminal and make sure that you are in a respective directory where the main file is present.
//The program will display a menu with options. Enter the number corresponding to the option you want to choose, and press Enter.
//The program will perform the selected action based on your input.

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const (
	baseURL = "https://reqres.in/api" // API base URL
)

type CreateRequest struct {
	Email string `json:"email"`
	Job   string `json:"job"`
}

type UpdateRequest struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UpdateResponse struct {
	Name     string `json:"name"`
	Job      string `json:"job"`
	UpdateAt string `json:"updatedAt"`
}

type CreateResponse struct {
	Email     string `json:"email"`
	Job       string `json:"job"`
	Id        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

// method to create a user.
func createUser(token string) {
	createData := CreateRequest{
		Email: "sai@gmail.com",
		Job:   "Developer",
	}

	createDataJSON, err := json.Marshal(createData)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", baseURL+"/users", bytes.NewBuffer(createDataJSON))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var createResponse CreateResponse
	if err := json.NewDecoder(resp.Body).Decode(&createResponse); err != nil {
		log.Fatal(err)
	}

	fmt.Println(createResponse)
	fmt.Println(resp.StatusCode)
}

// login method
func login() string {
	loginData := url.Values{
		"email":    {"eve.holt@reqres.in"},
		"password": {"cityslicka"},
	}

	resp, err := http.PostForm(baseURL+"/login", loginData)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Login failed. Status code: %d, Response: %s", resp.StatusCode, string(body))
	}

	var loginResponse LoginResponse
	if err := json.Unmarshal(body, &loginResponse); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Login successful.")
	fmt.Println("Token:", loginResponse.Token)
	fmt.Println(resp.StatusCode)

	return loginResponse.Token
}

//method to get all users to the respective page.

func getAllUsers(token string) {
	req, err := http.NewRequest("GET", baseURL+"/users?page=1", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res)
	fmt.Println("Listed users successfully")
	fmt.Println(resp.StatusCode)
}

//method to update user according to the input given.

func updateUser(token string) {
	updateRequest := UpdateRequest{
		Name: "sai seeramreddy",
		Job:  "developer",
	}

	updateDataJSON, err := json.Marshal(updateRequest)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("PUT", baseURL+"/users/2", bytes.NewBuffer(updateDataJSON))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var updateResponse UpdateResponse
	if err := json.NewDecoder(resp.Body).Decode(&updateResponse); err != nil {
		log.Fatal(err)
	}

	fmt.Println(updateResponse)
	if resp.StatusCode == 200 {
		fmt.Println("Updated Successfully")
	}
	fmt.Println(resp.StatusCode)
}

// method to delete record according to the id given.
func deleteUser(token string) {
	req, err := http.NewRequest("DELETE", baseURL+"/users/9", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp)
	if resp.StatusCode == 204 {
		fmt.Println("Deleted Successfully")
	}
	fmt.Println(resp.StatusCode)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Select an option:")
	fmt.Println("1. Login")
	fmt.Println("2. Create user")
	fmt.Println("3. Get all users")
	fmt.Println("4. Update user")
	fmt.Println("5. Delete user")

	optionStr, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	optionStr = strings.TrimSpace(optionStr)
	option, err := strconv.Atoi(strings.TrimRight(optionStr, "\r\n"))
	if err != nil {
		log.Fatal(err)
	}

	var token string

	switch option {
	case 1:
		token = login()
	case 2:
		token = login()
		if token != "" {
			createUser(token)
		}
	case 3:
		token = login()
		if token != "" {
			getAllUsers(token)
		}
	case 4:
		token = login()
		if token != "" {
			updateUser(token)
		}
	case 5:
		token = login()
		if token != "" {
			deleteUser(token)
		}
	default:
		fmt.Println("Invalid option")
	}
}

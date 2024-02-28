package req

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type PostData struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}

func Post() {
	// Sample data
	postData := PostData{
		Title:  "foo!",
		Body:   "bar",
		UserId: 1,
	}

	// Encode the data as JSON
	jsonData, err := json.Marshal(postData)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Define the URL of the API server
	apiUrl := "https://jsonplaceholder.typicode.com/posts"

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new request object
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	// req.Header.Set("app-id", "615987155bd63ef467c354ac") // Add your dummy API app ID here

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
		fmt.Println("Request successful:", resp.StatusCode)
	} else {
		fmt.Println("Error: Unexpected status code:", resp.StatusCode)
	}

	// Read the response body
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	// Print the response
	fmt.Println("Response:", response)
}

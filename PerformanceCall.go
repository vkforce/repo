package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

func main() {
	for i := 1; i <= 6; i++ {

		filePath := filepath.Join("C:/Users/vedhas.kharche/Documents/odbc(T)/query-3-fields", fmt.Sprintf("query_%v.json", i))
		jsonData, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		var queryData any
		err = json.Unmarshal(jsonData, &queryData)
		if err != nil {
			panic(err)
		}

		postQuery(queryData, fmt.Sprintf("query_%v", i))

	}
	return
}

func postQuery(request any, queryId string) {
	// URL to send the POST request
	url := "https://perf-cluster-us-east-2.insights.qa.forcepointone.com/api/odbcquery"
	client := &http.Client{}

	bearerToken := "fp_oidc_at_Y5wdkyQkNfXAB9uchnEsjprPn8RjcDC7azmAkQ3NVvo.JnO5cx8U1swaEpeXgY8Mjd7PhU6PqOzwwiY8DbrIrPw"

	requestString, err := json.Marshal(request)
	if err != nil {
		log.Println(err)
		//continue
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestString))
	if err != nil {
		log.Println(err)
		//continue
	}

	// Set the authorization header
	req.Header.Set("Authorization", "Bearer "+bearerToken)

	requestTime := time.Now()

	fmt.Println(queryId, ": Request Time:-", requestTime)

	// Send the request using the client
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		//continue
	}
	defer resp.Body.Close()

	// Read the response body
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		//continue
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		log.Printf("POST request failed with status code: %d, Response: %s\n", resp.StatusCode, string(respData))
		//continue
	}

	fmt.Println("Response Time:-", time.Now().Sub(requestTime).Seconds())
	// fmt.Println(string(respData))

}

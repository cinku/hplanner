package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var USERS_COUNT = 60
var API_URL = "https://firestore.googleapis.com/v1/projects/ittalkvue/databases/(default)/documents/requests"
var SLEEP_TRESHOLD = 7

func main() {
	cities := []string{
		"Amman",
		"Athens",
		"Barcelona",
		"Bari",
		"Berlin",
		"Cagliari",
		"Copenhagen",
		"Lisbon",
		"London",
		"Marseille",
		"Milan",
		"Palma Mallorca",
		"Prague",
		"Rome",
		"Stockholm",
		"Zadar",
	}
	months := []string{
		"May",
		"June",
		"July",
	}

	var requests []string

	fmt.Println("SIMULATING USERS...")

	for i := 0; i < USERS_COUNT; i++ {
		rand.Seed(time.Now().UnixNano())

		city := cities[rand.Intn(len(cities))]
		month := months[rand.Intn(len(months))]

		docID := createRequest(city, month)

		decideIfYouNeedToSleep()

		requests = append(requests, docID)
	}

	fmt.Println("SIMULATING DOWNTIME FOR THE ROBOT TO PROCESS STUFF")

	time.Sleep(time.Duration(10) * time.Second)

	fmt.Println("SIMULATING ROBOT RESPONSES...")

	for _, element := range requests {
		updateRequest(element)
		time.Sleep(time.Duration(1) * time.Second)
	}

}

type FirebaseResponse struct {
	Name string
}

func decideIfYouNeedToSleep() {
	min := 1
	max := 10
	number := rand.Intn(max-min) + min
	if number > SLEEP_TRESHOLD {
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func createRequest(city string, month string) string {
	client := &http.Client{}
	bodyString := []byte(fmt.Sprintf(`{ "fields": { "city": { "stringValue": "%s" }, "month": { "stringValue": "%s" },
	"name": { "stringValue": "testUser" }, "email": { "stringValue": "" }, "status": { "stringValue": "processing" }, "canceled": { "booleanValue": false } } }`, city, month))
	req, err := http.NewRequest("POST", API_URL, bytes.NewBuffer(bodyString))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	target := new(FirebaseResponse)

	json.NewDecoder(resp.Body).Decode(&target)

	tokens := strings.Split(target.Name, "/")
	docID := tokens[len(tokens)-1]
	return docID
}

func updateRequest(id string) {
	client := &http.Client{}
	updateURL := API_URL + "/" + id + "?updateMask.fieldPaths=status"
	bodyString := []byte(`{ "fields": { "status": { "stringValue": "done" } } }`)
	putReq, err := http.NewRequest("PATCH", updateURL, bytes.NewBuffer(bodyString))
	putReq.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(putReq)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

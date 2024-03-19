package helpers

import (
	"encoding/xml"
	"internal/database"
	"io/ioutil"
	"log"
	"net/http"
)

func FetchDataFromFeed(url string) {
	// Fetch data from feed
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// parse xml
	defer resp.Body.Close()
	xmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// parse xml
	var feed database.Feed
	err = xml.Unmarshal(xmlData, &feed)
	if err != nil {
		log.Fatal(err)
	}
	
}
package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"net/http"
	"io/ioutil"
)
	
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to loading .env file")
	}
	applicationId := os.Getenv("APPLICATION_ID")

	url := "https://app.rakuten.co.jp/services/api/Travel/SimpleHotelSearch/20170426"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	params := request.URL.Query()
	params.Set("applicationId", applicationId)
	params.Set("largeClassCode", "japan")
	params.Set("middleClassCode", "tokyo")
	params.Set("smallClassCode", "tokyo")
	params.Set("detailClassCode", "A")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(body)
}
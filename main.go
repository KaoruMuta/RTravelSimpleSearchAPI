package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type SearchResults struct {
	PagingInfo PagingInfo `json:"pagingInfo"`
	Hotels [][]Hotel `json:"hotels"`
}

type PagingInfo struct {
	RecordCount int `json:"recordCount"`
	PageCount int `json:"pageCount"`
	Page int `json:"page"`
	First int `json:"first"`
	Last int `json:"last"`
}

type Hotel struct {
	HotelBasicInfo HotelBasicInfo `json:"hotelBasicInfo"`
	HotelRatingInfo HotelRatingInfo `json:"hotelRatingInfo"`
	HotelDetailInfo HotelDetailInfo `json:"hotelDetailInfo"`
	HotelFacilitiesInfo HotelFacilitiesInfo `json:"hotelFacilitiesInfo"`
	HotelPolicyInfo HotelPolicyInfo `json:"hotelPolicyInfo"`
	HotelOtherInfo HotelOtherInfo `json:"hotelOtherInfo"`
}

type HotelBasicInfo struct {
	HotelNo int `json:"hotelNo"`
}

type HotelRatingInfo struct {
	ServiceAverage float32 `json:"serviceAverage"`
}

type HotelDetailInfo struct {
	ReserveTelephoneNo string `json:"reserveTelephoneNo"`
}

type HotelFacilitiesInfo struct {
	HotelRoomNum int `json:"hotelRoomNum"`
}

type HotelPolicyInfo struct {
	Note string `json:"note"`
}

type HotelOtherInfo struct {
	Privilege string `json:"privilege"`
}

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
	params.Set("formatVersion", "2")

	request.URL.RawQuery = params.Encode()
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

	var searchResults SearchResults
	if err := json.Unmarshal(body, &searchResults); err != nil {
		log.Fatal(err)
	}

	fmt.Println(searchResults.Hotels[0][0].HotelBasicInfo.HotelNo)
}
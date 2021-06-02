package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiResult struct {
	Date        string
	Day         string
	Icon        string
	Description string
	Status      string
	Degree      string
	Min         string
	Max         string
	Night       string
	Humidity    string
}

type ApiResponse struct {
	Success bool        `[]byte:"success" json:"success,omitempty"`
	City    string      `[]byte:"city"    json:"city,omitempty"`
	Result  []ApiResult `[]byte:"result"  json:"result,omitempty"`
}

func ApiRequest(city string) (ApiResponse, error) {
	url := "https://api.collectapi.com/weather/getWeather?data.lang=tr&data.city="+city

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "apikey 4XImnunyJMzNV0OnbezWLG:5LAgt0mZrDublTe3ZhSkzX")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var newBody ApiResponse
	json.Unmarshal(body, &newBody)
	return newBody, nil
}

func main() {
	fmt.Printf("Şehir ismi giriniz: ")
	var city string
	fmt.Scanln(&city)

	request, _ := ApiRequest(city)

	fmt.Println(city + " İçin Hava Durumu Tahminleri")
	for _, element := range request.Result{
		fmt.Println("Gün       : " + element.Day)
		fmt.Println("Tarih     : " + element.Date)
		fmt.Println("Açıklama  : " + element.Description)
		fmt.Println("Derece    : " + element.Degree)
		fmt.Println("")
	}
}
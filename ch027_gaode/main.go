package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// 高德地图的key
const Key = "key"

func main() {
	//ch1()

	ch2()
}

// 地理编码是将地址描述转换为经纬度坐标的过程。
func ch1() {
	address := "河北大学东门"
	city := "保定"

	result, err := geoCode(address, city, Key)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}

func geoCode(address, city, key string) (string, error) {
	baseURL := "https://restapi.amap.com/v3/geocode/geo"
	params := url.Values{}
	params.Add("key", key)
	params.Add("city", city)
	params.Add("address", address)
	params.Add("citylimit", "true")

	fullURL := baseURL + "?" + params.Encode()
	resp, err := http.Get(fullURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// 路线规划
func ch2() {
	origin := "116.407526,39.90403"
	destination := "116.384548,39.914927"

	result, err := routePlanning(origin, destination, Key)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}

func routePlanning(origin, destination, key string) (string, error) {
	baseURL := "https://restapi.amap.com/v3/direction/driving"
	params := url.Values{}
	params.Add("key", key)
	params.Add("origin", origin)
	params.Add("destination", destination)

	fullURL := baseURL + "?" + params.Encode()
	resp, err := http.Get(fullURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

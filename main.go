package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type Status struct {
	Status Weather `json:"status"`
}

type Weather struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	for {
		rand.Seed(time.Now().UTC().UnixNano())
		a := randInt(1, 100)
		b := randInt(1, 100)
		goJson(a, b)
		decJson()
		time.Sleep(15 * time.Second)
	}

}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func goJson(a int, b int) {

	var status Status

	status.Status.Water = a
	status.Status.Wind = b

	file, _ := os.Create("weather.json")
	defer file.Close()

	value, _ := json.Marshal(status)
	file.Write(value)
}

func decJson() {

	var statusWater, statusWind string

	jsonFile, err := os.Open("weather.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var status Status

	json.Unmarshal(byteValue, &status)

	waterValue := status.Status.Water
	windValue := status.Status.Wind

	if waterValue < 5 {
		statusWater = "Aman"
	} else if waterValue >= 6 && waterValue <= 8 {
		statusWater = "Siaga"
	} else if waterValue > 8 {
		statusWater = "Bahaya"
	}

	if windValue < 6 {
		statusWind = "Aman"
	} else if windValue >= 7 && windValue <= 15 {
		statusWind = "Siaga"
	} else if windValue > 15 {
		statusWind = "Bahaya"
	}

	fmt.Printf("Water: %d, Status: %s \n", waterValue, statusWater)
	fmt.Printf("Wind: %d, Status: %s \n", windValue, statusWind)
}

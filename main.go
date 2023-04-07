package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Element struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	// Run the app
	RunApp()

	// stop with interupt
	// Press ctrl + c to stop program
	StopWithInterupt()
}

func ResultRequest() {
	// struct data
	data := Element{
		Water: MyNum(),
		Wind:  MyNum(),
	}

	reqJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(reqJson))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))

	var element Element

	err = json.Unmarshal(body, &element)
	if err != nil {
		log.Fatalln(err)
	}

	var statusWater, statusWind string
	// water
	if data.Water > 8 {
		statusWater = "bahaya"
	} else if data.Water > 5 {
		statusWater = "siaga"
	} else {
		statusWater = "aman"
	}
	// wind
	if data.Wind > 8 {
		statusWind = "bahaya"
	} else if data.Wind > 5 {
		statusWind = "siaga"
	} else {
		statusWind = "aman"
	}

	fmt.Printf("status water: %s\n", statusWater)
	fmt.Printf("status wind: %s\n", statusWind)
}

func MyNum() int {
	rand.Seed(time.Now().Unix())

	rangeLower := 1
	rangeUpper := 100

	randomNum := rangeLower + rand.Intn(rangeUpper-rangeLower+1)

	return randomNum
}

func RunApp() {
	ticker := time.NewTicker(15 * time.Second)
	func() {
		for {
			ResultRequest()
			fmt.Println("")
			fmt.Println("Wait for 15 seconds...")
			fmt.Println("")
			<-ticker.C
		}
	}()
}

func StopWithInterupt() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for {
		menu()
	}
}

// < --- Menu functions --- >
func menu() {
	showMenu()
	option := scanCommand()
	switch option { // Here we don't have to provide the break statement
	case 1:
		site := getConfig("../cfg/monitor.json")
		startMonitoring(site.Site)
	case 2:
		showLogs()
	case 0:
		os.Exit(0)
	default:
		return
	}
}
func clear() {
	fmt.Print("\033[H\033[2J") // // ANSI clear code
}
func scanCommand() int {
	var option int
	fmt.Scan(&option) // Looks like we got some pointers here :)))
	clear()           // Let's clear this shit
	return option
}
func showMenu() {
	fmt.Println("\nWelcome to the site monitoring application")
	fmt.Println("\tWhat do you wanna do?")
	fmt.Println("[1] - Start monitoring")
	fmt.Println("[2] - Show logs")
	fmt.Println("[0] - Exit")
	fmt.Print("\n» ")
}

// < --- Monitor functions --- >
func startMonitoring(site []string) {
	// I'm watching you
	fmt.Println("Started monitoring...")
	sucess, failed := 0, 0
	for index, site := range site {
		fmt.Printf("[%d] -> ", index+1)
		resp, err := http.Get(site)
		// Validate response and errors
		if validateGet(resp, err) {
			sucess++
		} else {
			failed++
		}
	}
	fmt.Printf("Number of sucess requests: %d, number of failed requests: %d. Total: %d \n", sucess, failed, sucess+failed)
}
func validateGet(resp *http.Response, err error) bool {
	// Check for any error
	if err != nil {
		fmt.Println("An error occurred:", err)
		return false
	}
	// Check for the get status code
	if resp.StatusCode != 200 {
		fmt.Printf("Couldn't connect to %s, status code: %d \n", resp.Request.URL, resp.StatusCode)
		return false
	}

	fmt.Printf("Successfully connected to %s, with status code %d \n", resp.Request.URL, resp.StatusCode)
	return true
}

// < --- Log Functions --- >
func showLogs() {
	// No logs, no crimes.
	fmt.Println("Here we have some logs...")
}

// < --- Config --- >
type Config struct {
	Site []string `json:"sites"`
}

func getConfig(filePath string) Config {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config Config
	json.Unmarshal(byteValue, &config)
	return config
}

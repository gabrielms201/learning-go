package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var (
	config = getConfig("../cfg/monitor.json")
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
		startMonitoring()
	case 2:
		showLogs(config.LogPath)
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
func startMonitoring() {
	// I'm watching you
	fmt.Println("Started monitoring...")
	for i := 0; i < config.CheckNumber; i++ {
		fmt.Printf("Check Number: (%d):\n", i)
		sucess, failed := 0, 0
		for index, site := range config.Site {
			fmt.Printf("(%d)[%d] -> ", i, index+1)
			resp, err := http.Get(site)
			// Validate response and errors
			if validateGet(resp, err) {
				writeLog(site, true)
				sucess++
			} else {
				writeLog(site, false)
				failed++
			}
		}
		fmt.Printf("Number of sucess requests: %d, number of failed requests: %d. Total: %d \n", sucess, failed, sucess+failed)
		fmt.Println("Waiting...")
		time.Sleep(time.Duration(config.DelayAfterCheck) * time.Second)
	}
}
func validateGet(resp *http.Response, err error) bool {
	// Check for any error
	if err != nil {
		str := fmt.Sprintf("An error occurred: %s \n", err)
		fmt.Println(str)
		verboseLog(str)
		return false
	}
	// Check for the get status code
	if resp.StatusCode != 200 {
		str := fmt.Sprintf("Couldn't connect to %s, status code: %d \n", resp.Request.URL, resp.StatusCode)
		fmt.Printf(str)
		verboseLog(str)
		return false
	}

	str := fmt.Sprintf("Successfully connected to %s, with status code %d \n", resp.Request.URL, resp.StatusCode)
	fmt.Printf(str)
	verboseLog(str)
	return true
}

// < --- Log Functions --- >
func showLogs(logPath string) {
	// No logs, no crimes.
	fmt.Println("Here we have some logs...")
	file, err := ioutil.ReadFile(logPath)
	if err != nil {
		fmt.Println("Error reading log file:", err)
	}
	fmt.Println(string(file))
}

// < --- Config --- >
type Config struct {
	Site            []string `json:"sites"`
	CheckNumber     int      `json:"check_number"`
	DelayAfterCheck int      `json:"delay_after_check"`
	LogPath         string   `json:"log_path"`
	FullLogPath     string   `json:"full_log_path"`
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

func writeLog(site string, status bool) {
	file, err := os.OpenFile(config.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening log file:", err)
	}
	time := time.Now().Format("[02/01/2006 15:04:05] ")
	if status {
		file.WriteString(time + "SITE:\t" + site + ", ONLINE \n")
	} else {
		file.WriteString(time + "SITE:\t" + site + ", OFFLINE \n")
	}

	file.Close()
}
func verboseLog(text string) {
	file, err := os.OpenFile(config.FullLogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening log file:", err)
	}
	time := time.Now().Format("[02/01/2006 15:04:05] ")
	file.WriteString(time + text)
	file.Close()
}

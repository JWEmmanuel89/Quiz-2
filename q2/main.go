// Project: Go channels (Buffered Channel)
package main

import "fmt"

func main() {
	//Create channels
	location := make(chan string, 3)
	duration := make(chan string, 3)
	size := make(chan int, 3)
	vesselName := make(chan string, 3)

	//Input data into channels
	location <- "North"
	location <- "Central"
	location <- "South"
	duration <- "6 weeks"
	duration <- "4 weeks"
	duration <- "8 weeks"
	size <- 12
	size <- 8
	size <- 10
	vesselName <- "Vessel 1"
	vesselName <- "Vessel 5"
	vesselName <- "Vessel 10"

	//Print data
	for i := 0; i < 3; i++ {
		fmt.Println("AOR:", <-location)
		fmt.Println("Duration:", <-duration)
		fmt.Println("Team Size:", <-size)
		fmt.Println("Vessel Name:", <-vesselName)
		fmt.Println()
	}
}

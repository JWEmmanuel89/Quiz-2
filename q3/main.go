//Project: GoRountines

package main

import (
	"fmt"
	"time"
)

// Create functions that will be used with goroutines
func north() {
	fmt.Println("AOR: North")
	fmt.Println("Duration: 6 weeks")
	fmt.Println("Team Size:", 12)
	fmt.Println("Vessel Name: Vessel 1")
	fmt.Println()
}

func central() {
	fmt.Println("AOR: Central")
	fmt.Println("Duration: 4 weeks")
	fmt.Println("Team Size:", 8)
	fmt.Println("Vessel Name: Vessel 5")
	fmt.Println()
}

func south() {
	fmt.Println("AOR: South")
	fmt.Println("Duration: 8 weeks")
	fmt.Println("Team Size:", 10)
	fmt.Println("Vessel Name: Vessel 10")
	fmt.Println()
}

func main() {
	//Call functions using go
	go north()
	go central()
	go south()
	//Use of sleep delay
	time.Sleep(5 * time.Second)
}

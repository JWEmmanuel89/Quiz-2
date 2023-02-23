// Project: WaitGroup
package main

import (
	"fmt"
	"sync"
)

// Functions using wait group
func north(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("AOR: North")
	fmt.Println("Duration: 6 weeks")
	fmt.Println("Team Size:", 12)
	fmt.Println("Vessel Name: Vessel 1")
	fmt.Println()
}

func central(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("AOR: Central")
	fmt.Println("Duration: 4 weeks")
	fmt.Println("Team Size:", 8)
	fmt.Println("Vessel Name: Vessel 5")
	fmt.Println()
}

func south(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("AOR: South")
	fmt.Println("Duration: 8 weeks")
	fmt.Println("Team Size:", 10)
	fmt.Println("Vessel Name: Vessel 10")
	fmt.Println()
}

func main() {
	//Create wait group variable
	var wg sync.WaitGroup

	//Add number of functions to wait group
	wg.Add(3)

	//Call go functions
	go north(&wg)
	go central(&wg)
	go south(&wg)

	//Delay main
	wg.Wait()
}

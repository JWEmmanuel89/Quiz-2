// Project: Go Interface
package main

import "fmt"

// Create Interface
type Deploy interface {
	//Create method signatures
	location() string
	duration() string
	size() int
	vessel() string
}

// Struct to implement interface
type Info struct {
	area       string
	time       string
	teamSize   int
	vesselName string
}

// Implement location()
func (i Info) location() string {
	return i.area
}

// Implement duration()
func (i Info) duration() string {
	return i.time
}

// Implement size()
func (i Info) size() int {
	return i.teamSize
}

// Implement vessel()
func (i Info) vessel() string {
	return i.vesselName
}

// Access methods of interface
func final(d Deploy) {
	fmt.Println("AOR:", d.location())
	fmt.Println("Duration:", d.duration())
	fmt.Println("Team Size:", d.size())
	fmt.Println("Vessl Name:", d.vessel())
	fmt.Println()
}

// Main function
func main() {
	//Create struct members and assign values
	f1 := Info{"North", "6 weeks", 12, "Vessel 1"}
	f2 := Info{"Central", "4 weeks", 8, "Vessel 5"}
	f3 := Info{"South", "8 weeks", 10, "Vessel 10"}

	//Call final() with variable f1, f2 and f3
	final(f1)
	final(f2)
	final(f3)

}

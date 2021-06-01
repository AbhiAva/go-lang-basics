/////////////////////////////////////////////////////

// ---FIVE - Maps --- //
package main

import "fmt"

// Imports are done this way

var sampleMap map[string]int

func main() {
	sampleMap = map[string]int{
		"Chandler": 23,
		"Joey":     24,
	}

	currency := map[string]string{
		"USD": "US Dollars",
		"INR": "Indian Rupee",
		"AUD": "Australian Dollars",
	}

	// a. Adding to the map
	currency["EU"] = "Euros"
	fmt.Println("Currency with Euros added:\n", currency)

	// b. Removing from the map
	delete(currency, "AUD")
	fmt.Println("Currency with Australian dollars removed:\n", currency)

	// c. Replacing one entry with another
	currency["EU"] = "Pounds"
	fmt.Println("Currency exchanged (No pun intended lol):\n", currency)

	// Ranging through the map
	for key, value := range currency {
		fmt.Printf("%v might be equal to %v\n", key, value)
	}

}

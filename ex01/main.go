package main

import "fmt"

func main() {

	var firstName string = "Samuel"
	var lastName string = "Ming"
	var dayOfBirth int = 9
	var monthOfBirth int = 9
	var yearOfBirth int = 2009
	var numberOfSiblings int = 1
	var heightInMeters float64 = 1.75
	var zodiacSign rune = 'G'

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}

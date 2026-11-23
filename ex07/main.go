package main

import "fmt"

func main() {
	modules := make(map[int]string)

	modules[346] = "Cloud Lösungen konzipieren und realisieren"
	modules[117] = "Informatik- und Netzinfrastruktur für ein kleines Unternehmen realisieren"
	modules[104] = "Datenmodell implementieren"

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 117)

	// TODO: add one
	modules[154] = "Applikationen für die Produktion vorbereiten"

	// TODO: replace one
	modules[117] = "Horst"

	fmt.Println(modules)
}

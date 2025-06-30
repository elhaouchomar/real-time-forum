package main

import (
	"fmt"
	"os"
)

func main() {
	// Vérifier si le nombre d'arguments est égal à 2
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}

	// Initialiser les chaînes et une carte pour éviter les doublons
	s1, s2 := os.Args[1], os.Args[2]
	seen := make(map[rune]bool)

	// Afficher l'union des caractères
	for _, char := range s1 + s2 {
		if !seen[char] {
			seen[char] = true
			fmt.Print(string(char))
		}
	}

	// Afficher une nouvelle ligne à la fin
	fmt.Println()
}

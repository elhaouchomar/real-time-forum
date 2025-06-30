package main

import (
	"fmt"
	"os"
)

func inter(s1, s2 string) {
	// Créer un tableau de booléens pour suivre les caractères déjà affichés
	seen := make(map[rune]bool)

	// Parcourir chaque caractère de s1
	for _, char := range s1 {
		// Vérifier si le caractère est présent dans s2 et s'il n'a pas déjà été affiché
		if !seen[char] && contains(s2, char) {
			// Marquer ce caractère comme affiché
			seen[char] = true
			// Afficher le caractère
			fmt.Print(string(char))
		}
	}
	fmt.Println() // Nouvelle ligne après l'affichage
}

// Fonction pour vérifier si un caractère est dans une chaîne
func contains(s string, c rune) bool {
	for _, char := range s {
		if char == c {
			return true
		}
	}
	return false
}

func main() {
	// Vérifier qu'il y a exactement 2 arguments
	if len(os.Args) != 3 {
		return
	}

	// Récupérer les arguments
	s1 := os.Args[1]
	s2 := os.Args[2]

	// Appeler la fonction pour afficher les caractères en commun sans doublons
	inter(s1, s2)
}

package main

import (
	"fmt"
	"strings"
)

func FifthAndSkip(str string) string {
	// Vérifier si la chaîne est vide
	if str == "" {
		return "\n"
	}

	// Si la chaîne a moins de 5 caractères, afficher "Invalid Input"
	if len(str) < 5 {
		return "Invalid Input\n"
	}

	// Supprimer les espaces de la chaîne
	str = strings.ReplaceAll(str, " ", "")

	// Initialiser une variable pour accumuler les résultats
	var result string

	// Parcourir la chaîne de caractères par groupes de 5
	for i := 0; i < len(str); i++ {
		// Ajouter le caractère à la chaîne résultat tous les 5 caractères
		if (i+1)%5 == 0 {
			result += string(str[i]) + " "
		} else if (i+1)%5 != 0 {
			result += string(str[i])
		}
	}

	// Retourner le résultat sans le dernier espace
	return strings.TrimSpace(result)
}

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz")) // abcde ghijk mnopq stuwx z
	fmt.Print(FifthAndSkip("This is a short sentence"))  // Thisi ashor sente ce
	fmt.Print(FifthAndSkip("1234"))                      // Invalid Input
}

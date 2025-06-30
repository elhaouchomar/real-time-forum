package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func reverseStrCap(args []string) {
	// Si aucun argument n'est fourni, ne rien afficher
	if len(args) == 0 {
		return
	}

	// Parcourir chaque argument
	for _, arg := range args {
		// Diviser l'argument en mots
		words := strings.Fields(arg)

		// Parcourir chaque mot
		for i, word := range words {
			// Si le mot a plus d'un caractère, on traite le mot
			if len(word) > 1 {
				// La première partie du mot (tout sauf la dernière lettre) en minuscule
				// La dernière lettre en majuscule
				words[i] = strings.ToLower(word[:len(word)-1]) + strings.ToUpper(string(word[len(word)-1]))
			} else {
				// Si le mot a un seul caractère, on le met en majuscule
				words[i] = strings.ToUpper(word)
			}
		}

		// Afficher les mots modifiés pour cet argument, séparés par un espace
		fmt.Print(strings.Join(words, " ") + " ")
	}

	// Afficher un retour à la ligne
	fmt.Println()
}

func main() {
	// Vérifier s'il y a des arguments fournis
	if len(os.Args) < 2 {
		return
	}

	// Appeler la fonction pour traiter les arguments
	reverseStrCap(os.Args[1:])
}

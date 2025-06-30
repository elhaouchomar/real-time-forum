package main

import (
	"fmt"
	"os"
)

func hiddenp(s1, s2 string) {
	// Si s1 est vide, il est considéré comme caché dans n'importe quelle chaîne
	if s1 == "" {
		fmt.Println(1)
		return
	}

	// Index pour parcourir s1
	index := 0

	// Parcourir chaque caractère de s2
	for i := 0; i < len(s2); i++ {
		if s2[i] == s1[index] {
			index++
			// Si tous les caractères de s1 sont trouvés, on affiche 1 et on termine
			if index == len(s1) {
				fmt.Println(1)
				return
			}
		}
	}

	// Si on a parcouru toute s2 sans trouver tous les caractères de s1, on affiche 0
	fmt.Println(0)
}

func main() {
	// Vérifier qu'il y a exactement 2 arguments
	if len(os.Args) != 3 {
		return
	}

	// Récupérer les arguments
	s1 := os.Args[1]
	s2 := os.Args[2]

	// Appeler la fonction pour vérifier si s1 est caché dans s2
	hiddenp(s1, s2)
}

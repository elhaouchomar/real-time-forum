package main

import (
	"fmt"
	"os"
)

func main() {
	// Vérifier si le nombre d'arguments est égal à 2
	if len(os.Args) != 3 {
		return
	}

	s1, s2 := os.Args[1], os.Args[2]

	// Pointeur pour parcourir la chaîne s2
	j := 0

	// Essayer de trouver chaque caractère de s1 dans s2 dans l'ordre
	for i := 0; i < len(s1); i++ {
		// Chercher le caractère s1[i] dans s2 à partir de la position j
		for j < len(s2) && s2[j] != s1[i] {
			j++
		}

		// Si on ne trouve pas le caractère s1[i] dans s2, on quitte
		if j == len(s2) {
			return
		}
		j++
	}

	// Si tous les caractères de s1 ont été trouvés dans s2, afficher s1
	fmt.Println(s1)
}

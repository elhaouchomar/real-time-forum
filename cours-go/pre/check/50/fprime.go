package main

import (
	"fmt"
	"os"
	"strconv"
)

// Fonction qui retourne les facteurs premiers d'un nombre
func primeFactors(n int) {
	if n < 2 {
		return // Si le nombre est inférieur à 2, il n'a pas de facteur premier
	}

	// Trouver les facteurs premiers en divisant par 2 jusqu'à ce qu'il ne soit plus divisible
	for n%2 == 0 {
		fmt.Print("2")
		if n > 2 {
			fmt.Print("*")
		}
		n /= 2
	}

	// Tester les facteurs impairs à partir de 3
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			fmt.Print(i)
			if n > i {
				fmt.Print("*")
			}
			n /= i
		}
	}

	// Si n est un nombre premier plus grand que 2, on l'affiche
	if n > 2 {
		fmt.Print(n)
	}
	fmt.Println() // Nouvelle ligne après l'affichage
}

func main() {
	// Vérifier que le nombre d'arguments est correct
	if len(os.Args) != 2 {
		return
	}

	// Convertir l'argument en entier
	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 0 {
		return // Si l'argument n'est pas valide ou inférieur ou égal à 0, on ne fait rien
	}

	// Appeler la fonction pour afficher les facteurs premiers
	primeFactors(num)
}

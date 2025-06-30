package main

import (
	"fmt"
)

func SaveAndMiss(arg string, num int) string {
	// Si num est 0 ou négatif, on retourne la chaîne d'origine
	if num <= 0 {
		return arg
	}

	// Initialiser la chaîne pour stocker les caractères sauvegardés
	var result string

	// Parcourir la chaîne en sets de taille 'num'
	for i := 0; i < len(arg); i += num * 2 {
		// Ajouter les caractères du set "sauvé"
		end := i + num
		if end > len(arg) {
			end = len(arg)
		}
		result += arg[i:end]
	}

	return result
}

func main() {
	fmt.Println(SaveAndMiss("123456789", 3))  // "123789"
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))  // "abcghimnostuz"
	fmt.Println(SaveAndMiss("", 3))  // ""
	fmt.Println(SaveAndMiss("hello you all ! ", 0))  // "hello you all ! "
	fmt.Println(SaveAndMiss("what is your name?", 0))  // "what is your name?"
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))  // "go Exercise Save and Miss"
}

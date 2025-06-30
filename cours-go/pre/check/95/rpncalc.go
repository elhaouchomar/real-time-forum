package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fonction pour évaluer l'expression RPN
func evaluateRPN(expression string) (int, bool) {
	tokens := strings.Fields(expression) // Séparer les éléments par espace
	stack := []int{}

	for _, token := range tokens {
		// Essayer de convertir le token en nombre
		num, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, num) // Ajouter le nombre à la pile
			continue
		}

		// S'assurer qu'il y a au moins deux opérandes avant d'appliquer un opérateur
		if len(stack) < 2 {
			return 0, false
		}

		// Récupérer les deux derniers nombres
		b := stack[len(stack)-1]
		a := stack[len(stack)-2]
		stack = stack[:len(stack)-2] // Enlever les deux derniers éléments

		// Appliquer l'opération
		switch token {
		case "+":
			stack = append(stack, a+b)
		case "-":
			stack = append(stack, a-b)
		case "*":
			stack = append(stack, a*b)
		case "/":
			if b == 0 {
				return 0, false // Division par zéro invalide
			}
			stack = append(stack, a/b)
		case "%":
			if b == 0 {
				return 0, false // Modulo par zéro invalide
			}
			stack = append(stack, a%b)
		default:
			return 0, false // Opérateur invalide
		}
	}

	// Il doit rester exactement un élément dans la pile
	if len(stack) != 1 {
		return 0, false
	}
	return stack[0], true
}

func main() {
	// Vérifier qu'il y a exactement un argument
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}

	result, valid := evaluateRPN(os.Args[1])
	if !valid {
		fmt.Println("Error")
		return
	}

	fmt.Println(result)
}

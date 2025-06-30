package main

import "fmt"

// Chunk découpe un tableau en plusieurs sous-tableaux de taille `size`
func Chunk(slice []int, size int) {
	if size == 0 {
		fmt.Println() // Si la taille est 0, on imprime un retour à la ligne
		return
	}

	// Découper le slice en chunks de taille `size`
	var result [][]int
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}

	// Afficher les chunks
	fmt.Println(result)
}

func main() {
	// Cas de test 1: Slice vide, taille 10
	Chunk([]int{}, 10)

	// Cas de test 2: Taille 0, devrait juste imprimer une nouvelle ligne
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)

	// Cas de test 3: Taille de chunk 3
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)

	// Cas de test 4: Taille de chunk 5
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)

	// Cas de test 5: Taille de chunk 4
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

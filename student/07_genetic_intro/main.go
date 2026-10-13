package main

import "fmt"

func fitness(candidate string) int {
	score := 0
	for i := 0; i < len(candidate); i++ {
		if candidate[i] == '1' {
			score++
		}
	}
	return score
}

func crossover(parent1, parent2 string, cutPoint int) string {
	return parent1[:cutPoint] + parent2[cutPoint:]
}

func mutate(candidate string, pos int) string {
	mutated := []byte(candidate)
	if mutated[pos] == '0' {
		mutated[pos] = '1'
	} else {
		mutated[pos] = '0'
	}
	return string(mutated)
}

func main() {
	parent1 := "1000"
	parent2 := "1010"
	cutPoint := 2

	child := crossover(parent1, parent2, cutPoint)
	fmt.Printf("Child %s has fitness %d\n", child, fitness(child))

	mutatedChild := mutate(child, 1)
	fmt.Printf("Mutated %s has fitness %d\n", mutatedChild, fitness(mutatedChild))
}

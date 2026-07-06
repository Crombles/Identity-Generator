package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	startOfTheProgram := 0
	if startOfTheProgram == 0 {
		fmt.Print("\nHi! Let's generate a random person. Choose gender: Random - 1, Male - 2, Female - 3, 0 - Exit (Type a number): ")
	}
	rand.Seed(time.Now().UnixNano())

	scanner := bufio.NewScanner(os.Stdin)

	for {
		nameMan := []string{"John", "James", "Robert", "Michael", "Alejandro", "Carlos", "Juan", "Manuel", "David", "William", "Diego", "Mateo"}[rand.Intn(12)]
		nameWom := []string{"Mary", "Patricia", "Jennifer", "Linda", "Sofia", "Isabella", "Camila", "Valentina", "Elizabeth", "Barbara", "Valeria", "Lucia"}[rand.Intn(12)]
		nameWomsur := []string{"Smith", "Johnson", "Williams", "Brown", "Garcia", "Rodriguez", "Gonzalez", "Fernandez", "Jones", "Miller", "Lopez", "Martinez"}[rand.Intn(12)]
		nameMansur := []string{"Smith", "Johnson", "Williams", "Brown", "Garcia", "Rodriguez", "Gonzalez", "Fernandez", "Jones", "Miller", "Lopez", "Martinez"}[rand.Intn(12)]
		gender := []string{"Female", "Male"}[rand.Intn(2)]

		if startOfTheProgram == 1 {
			fmt.Print("\nChoose gender: Random - 1, Male - 2, Female - 3, Exit - 0, (Type a number): ")
		}

		scanner.Scan()
		generationSelection := strings.TrimSpace(scanner.Text())

		if generationSelection == "3" {
			fmt.Println("\nFirst Name:", nameWom, "Last Name:", nameWomsur)
		}
		if generationSelection == "2" {
			fmt.Println("\nFirst Name:", nameMan, "Last Name:", nameMansur)
		}
		if generationSelection == "1" {
			fmt.Println("\nGender: ", gender)
			if gender == "Male" {
				fmt.Println("\nFirst Name:", nameMan, "Last Name:", nameMansur)
			}
			if gender == "Female" {
				fmt.Println("\nFirst Name:", nameWom, "Last Name:", nameWomsur)
			}
		}

		startOfTheProgram = 1

		if generationSelection == "0" {
			fmt.Println("\nShutting down...")
			return
		}
	}
}

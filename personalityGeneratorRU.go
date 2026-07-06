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
		fmt.Print("\nПривет! Сгенерируем случайную личность. Выбери пол: Рандом - 1, Мужской - 2, Женский - 3, 0 - Выход (Напиши цифру): ")
	}
	rand.Seed(time.Now().UnixNano())
	
	scanner := bufio.NewScanner(os.Stdin)

	for {
		patronymicWom := []string{"Алексеевна", "Никитична", "Ильинична", "Александровна", "Михайловна", "Рафаэлевна", "Германовна", "Владимировна", "Вадимовна", "Дмитриевна", "Евгеньевна", "Сергеевна"}[rand.Intn(12)]
		patronymicMan := []string{"Алексеевич", "Никитич", "Ильич", "Александрович", "Михайлович", "Рафаэлевич", "Германович", "Владимирович", "Вадимович", "Дмитриевич", "Сергеевич", "Евгеньевич"}[rand.Intn(12)]
		nameMan := []string{"Алексей", "Никита", "Илья", "Александр", "Михаил", "Рафаэль", "Герман", "Владимир", "Дмитрий", "Евгений", "Сергей", "Вадим"}[rand.Intn(12)]
		nameWom := []string{"Елена", "Юлия", "Ирина", "Снежанна", "Татьяна", "Лиза", "Светлана", "Анастасия", "Виктория", "Дарья", "Анна", "Мария"}[rand.Intn(12)]
		nameWomsur := []string{"Орлова", "Ведмидова", "Лиходедова", "Кузнецова", "Соколова", "Лебедева", "Разумовская", "Воронцова", "Шмидт", "Козлова", "Морозова", "Новикова"}[rand.Intn(12)]
		nameMansur := []string{"Орлов", "Ведмидов", "Лиходедова", "Кузнецов", "Соколов", "Лебедев", "Разумовский", "Воронцов", "Шмидт", "Козлова", "Морозов", "Новиков"}[rand.Intn(12)]
		gender := []string{"Женский", "Мужской"}[rand.Intn(2)]

		if startOfTheProgram == 1 {
			fmt.Print("\nВыбери пол: Рандом - 1, Мужской - 2, Женский - 3, Выход - 0, (Напиши цифру): ")
		}

		scanner.Scan()
		generationSelection := strings.TrimSpace(scanner.Text())

		if generationSelection == "3" {
			fmt.Println("\nИмя:", nameWom, "Фамилия:", nameWomsur, "Отчество:", patronymicWom)
		}
		if generationSelection == "2" {
			fmt.Println("\nИмя:", nameMan, "Фамилия:", nameMansur, "Отчество:", patronymicMan)
		}
		if generationSelection == "1" {
			fmt.Println("\nПол: ", gender)
			if gender == "Мужской" {
				fmt.Println("\nИмя:", nameMan, "Фамилия:", nameMansur, "Отчество:", patronymicMan)
			}
			if gender == "Женский" {
				fmt.Println("\nИмя:", nameWom, "Фамилия:", nameWomsur, "Отчество:", patronymicWom)
			}
		}

		startOfTheProgram = 1

		if generationSelection == "0" {
			fmt.Println("\nСворачиваемся...")
			return
		}
	}
}

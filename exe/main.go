package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// Vérifier si un fichier est passé en argument
	if len(os.Args) < 2 {
		log.Fatal("Veuillez fournir un fichier en paramètre.")
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		fmt.Println("d")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	wordsArr := []string{}
	for scanner.Scan() {
		wordsArr = append(wordsArr, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	hiddenWord := wordsArr[rand.Intn(len(wordsArr))]
	amoutOfLetters := (len(hiddenWord) / 10) + 1
	display := make([]rune, len(hiddenWord))
	for i := range display {
		display[i] = '_'
	}

	indicesChoisis := []int{}
	for len(indicesChoisis) < amoutOfLetters {
		ind := rand.Intn(len(hiddenWord))
		dejaChoisi := false
		for _, i := range indicesChoisis {
			if i == ind {
				dejaChoisi = true
				break
			}
		}
		if !dejaChoisi {
			indicesChoisis = append(indicesChoisis, ind)
			display[ind] = rune(hiddenWord[ind])
		}
	}

	isRunning := true
	for isRunning {
		fmt.Println(string(display))
		var option string
		fmt.Println("Choisissez une lettre")
		fmt.Scan(&option)
		for k, _ := range hiddenWord {
			if option == string(hiddenWord[k]) {
				display[k] = rune(hiddenWord[k])
			}
			if display[k] != 95 {
			}
		}
		clear()
	}

}

func clear() {
	fmt.Printf("\033[H\033[2J")
}

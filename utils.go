package hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func ReadWordsFromFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
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
	return wordsArr
}

func SelectRandomWord(wordsArr []string) string {
	return wordsArr[rand.Intn(len(wordsArr))]
}
func GetDisplayAmount(str string) int {
	return (len(str) / 10) + 1
}

func InitializeDisplay(hiddenWord string) []rune {
	amountofL := GetDisplayAmount(hiddenWord)
	display := make([]rune, len(hiddenWord))
	for i := range display {
		display[i] = '_'
	}

	indicesChoisis := []int{}
	for len(indicesChoisis) < amountofL {
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
	return display
}

func Clear() {
	fmt.Printf("\033[H\033[2J")
}

func CheckComp(dis []rune) bool {
	isRunning := false
	for _, k := range dis {
		if k == 95 {
			isRunning = true
		}
	}
	return isRunning
}

func printHangman(c int) {
	switch c {
	case 6:
		fmt.Println(`
  +---+
  |   |
      |
      |
      |
      |
=========`)
	case 5:
		fmt.Println(`
  +---+
  |   |
  O   |
      |
      |
      |
=========`)
	case 4:
		fmt.Println(`
+---+
|   |
O   |
|   |
    |
    |
=========`)

	case 3:
		fmt.Println(`
+---+
|   |
O   |
/|  |
    |
    |
=========`)

	case 2:
		fmt.Println(`
 +---+
 |   |
 O   |
/|\\ |
     |
     |
=========`)

	case 1:
		fmt.Println(`
 +---+
 |   |
 O   |
/|\\ |
/    |
     |
=========`)

	case 0:
		fmt.Println(`
 +---+
 |   |
 O   |
/|\\ |
/ \\ |
     |
=========`)
	}
}

package hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
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

func GetFiles() {
	rand.Seed(time.Now().UnixNano())
	dir := "./mots"

	// Lire le contenu du dossier
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du dossier :", err)
		return
	}

	var files []string
	// Parcourir les entrées
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".txt" {
			color := getColor(entry)
			files = append(files, entry.Name()[0:len(entry.Name())-4])
		}
	}
	// Vérifier si un fichier est passé en argument
	fmt.Println("Choisissez un thème de mots parmis la lise :")

	for _, k := range files {

		fmt.Print(" -")
		fmt.Println(k)
	}
}

func SelectFile() string {
	var chosenFile string
	fmt.Scan(&chosenFile)
	fileName := ".\\mots\\" + chosenFile + ".txt"
	return fileName
}

func getColor(fileName string) {

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

package main

import (
	hangman "Hangman" // replace with your actual module name
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func main() {
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
			files = append(files, entry.Name()[0:len(entry.Name())-4])
		}
	}
	// Vérifier si un fichier est passé en argument
	fmt.Println("quel fichier??")
	for _, k := range files {
		fmt.Println(k)
	}
	var chosenFile string
	fmt.Scan(&chosenFile)
	fileName := ".\\mots\\" + chosenFile + ".txt"
	wordsArr := hangman.ReadWordsFromFile(fileName)

	hiddenWord := hangman.SelectRandomWord(wordsArr)
	display := hangman.InitializeDisplay(hiddenWord)

	// Logic for choosing letters
	hangman.PlayGame(hiddenWord, display)
}

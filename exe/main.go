package main

//si tu veux ajouter un nouveau fichier avec des mots, créé un fichier .txt, ouvre le et met a la première ligne la difficultée de ton fichier de mots
//ensuite copie colle tes mots, chaque mot est différentié de l'autre pas un saut a la ligne, pas de majuscules, d'accents ou autres caractères spéciaux --- espaces authorisés

import (
	hangman "Hangman" // replace with your actual module name
)

func main() {
	hangman.GetFiles()

	fileName := hangman.SelectFile()
	wordsArr := hangman.ReadWordsFromFile(fileName)

	hiddenWord := hangman.SelectRandomWord(wordsArr)
	display := hangman.InitializeDisplay(hiddenWord)

	// Logic for choosing letters
	hangman.PlayGame(hiddenWord, display)
}

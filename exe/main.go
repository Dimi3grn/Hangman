package main

import (
	hangman "Hangman" // replace with your actual module name
)

func main() {
	hangman.Clear()
	hangman.GetFiles()

	fileName := hangman.SelectFile()
	wordsArr := hangman.ReadWordsFromFile(fileName)

	hiddenWord := hangman.SelectRandomWord(wordsArr)
	display := hangman.InitializeDisplay(hiddenWord)

	// Logic for choosing letters
	hangman.PlayGame(hiddenWord, display)
}

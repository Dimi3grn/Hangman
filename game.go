package hangman

import (
	"fmt"
	"strings"
)

func PlayGame(hiddenWord string, display []rune) {
	Clear()
	isRunning := true
	tries := 6                     // Nombre limité d'essais
	attemptedLetters := []string{} // Liste des lettres déjà tentées
	attemptedWords := []string{}

	fmt.Println("Le mot à deviner a", len(hiddenWord), "lettres.")
	fmt.Println("Affichage actuel :", string(display)) // Afficher l'affichage initial

	for isRunning && tries > 0 {
		printHangman(tries)
		fmt.Println("Choisissez une lettre ou proposez un mot entier (tentatives restantes :", tries, ")")

		var option string
		fmt.Scanln(&option)                // Utilisez Scanln pour permettre des mots avec des espaces
		option = strings.TrimSpace(option) // Remove any leading/trailing spaces

		// Vérifier si l'utilisateur a proposé un mot entier
		if len(option) > 1 {
			if contains(attemptedWords, option) {
				Clear()
				fmt.Println("Vous avez déjà proposé ce mot, essayez-en un autre.")
				fmt.Println("Le mot à deviner a", len(hiddenWord), "lettres.")
				fmt.Println("Affichage actuel :", string(display))
				continue
			}

			// Comparer le mot entier proposé avec le mot caché
			if strings.ReplaceAll(option, " ", "") == strings.ReplaceAll(hiddenWord, " ", "") {
				// Si le mot est correct, on met à jour l'affichage et on termine le jeu
				for i, char := range hiddenWord {
					display[i] = char
				}
				isRunning = false
				fmt.Println("Vous avez deviné le mot, bien joué à vous !")
				break
			} else {
				// Si le mot est incorrect, retire deux essais
				Clear()
				tries -= 2
				attemptedWords = append(attemptedWords, option)
				if tries <= 0 {
					fmt.Println("Dommage ! Vous avez épuisé vos tentatives. Le mot était :", hiddenWord)
					break
				}
				fmt.Println("Mauvaise proposition ! Deux tentatives en moins.")
				fmt.Println("Le mot à deviner a", len(hiddenWord), "lettres.")
				fmt.Println("Affichage actuel :", string(display))
				continue
			}
		}

		// Vérifier si la lettre a déjà été tentée
		if len(option) == 1 {
			if contains(attemptedLetters, option) {
				Clear()
				fmt.Println("Vous avez déjà choisi cette lettre, essayez-en une autre.")
				fmt.Println("Le mot à deviner a", len(hiddenWord), "lettres.")
				fmt.Println("Affichage actuel :", string(display)) // Afficher l'affichage initial
				continue                                           // Ne pas réduire le nombre d'essais
			}

			// Ajouter la lettre à la liste des lettres tentées
			attemptedLetters = append(attemptedLetters, option)

			// Mettre à jour l'affichage
			if !UpdateDisplay(hiddenWord, display, option) {
				Clear()
				tries-- // Décrémente le nombre d'essais si la lettre n'est pas trouvée
				fmt.Println("Mauvaise lettre !")
			}
		}

		Clear()
		fmt.Println("Affichage actuel :", string(display)) // Afficher l'affichage mis à jour
		// Vérifier les conditions de victoire uniquement ici
		if CheckComp(display) == false {
			isRunning = false
			fmt.Println("Vous avez deviné le mot, bien joué à vous !")
		} else if tries == 0 {
			fmt.Println("Dommage ! Vous avez épuisé vos tentatives. Le mot était :", hiddenWord)
		}
	}
}

func contains(attemptedLetters []string, letter string) bool {
	for _, attempted := range attemptedLetters {
		if attempted == letter {
			return true
		}
	}
	return false
}

func UpdateDisplay(hiddenWord string, display []rune, option string) bool {
	correctGuess := false
	for k, char := range hiddenWord {
		// Handle spaces separately so that they're not hidden
		if option == string(char) && char != ' ' {
			display[k] = char
			correctGuess = true
		}
	}
	return correctGuess // Renvoie true si la lettre est correcte
}

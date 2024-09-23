package hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func Choseword(k string) string {
	file, err := os.Open(k)
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
	return hiddenWord
}

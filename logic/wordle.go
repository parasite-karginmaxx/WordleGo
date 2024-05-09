package logic

import (
	"fmt"
	"strings"
)

const WORD_LENGTH = 5

type Game interface {
	Start()
}

type WordleGame struct {
	secretWord string
	guesses    int
	maxGuesses int
}

func NewWordleGame(secretWord string, maxGuesses int) *WordleGame {
	return &WordleGame{
		secretWord: secretWord,
		guesses:    0,
		maxGuesses: maxGuesses,
	}
}

func (wg *WordleGame) Start() {
	fmt.Println("Добро пожаловать в игру Wordle!")
	fmt.Printf("Угадайте слово из 5 букв.\n")

	for {
		if wg.guesses >= wg.maxGuesses {
			wg.displayGameOver()
			return
		}

		if !wg.takeUserInput() {
			continue
		}

		if wg.checkGuess(guess) {
			return
		}
	}
}

func (wg *WordleGame) displayGameOver() {
	fmt.Println("У вас закончились попытки. Загаданное слово было:")
	colorVector := getFilledSymbols(wg.secretWord, wg.secretWord)
	displayWord(wg.secretWord, colorVector)
}

func (wg *WordleGame) takeUserInput() bool {
	fmt.Printf("Попытка #%d. Введите ваше предположение: ", wg.guesses+1)
	var guess string
	fmt.Scanln(&guess)
	guess = strings.ToLower(guess)

	if len(guess) != WORD_LENGTH {
		fmt.Printf("Пожалуйста, введите слово из 5 букв.\n")
		return false
	}

	wg.guesses++
	return true
}

func (wg *WordleGame) checkGuess(guess string) bool {
	if guess == wg.secretWord {
		fmt.Println("Поздравляем! Вы угадали слово!")
		colorVector := getFilledSymbols(wg.secretWord, guess)
		displayWord(wg.secretWord, colorVector)
		return true
	} else {
		colorVector := getFilledSymbols(wg.secretWord, guess)
		displayWord(guess, colorVector)
		return false
	}
}

func getFilledSymbols(secretWord string, guess string) [WORD_LENGTH]string {
	colorVector := [WORD_LENGTH]string{}
	for i := range colorVector {
		colorVector[i] = "Gray"
	}
	yellowLock := [WORD_LENGTH]bool{}

	for j, guessLetter := range guess {
		for k, letter := range secretWord {
			if guessLetter == letter && j == k {
				colorVector[j] = "Green"
				yellowLock[k] = true
				break
			}
		}
	}
	for j, guessLetter := range guess {
		for k, letter := range secretWord {
			if guessLetter == letter && colorVector[j] != "Green" && !yellowLock[k] {
				colorVector[j] = "Yellow"
				yellowLock[k] = true
			}
		}
	}
	return colorVector
}

func displayWord(word string, colorVector [WORD_LENGTH]string) {
	for i, c := range word {
		switch colorVector[i] {
		case "Green":
			fmt.Print("\033[42m\033[1;30m")
		case "Yellow":
			fmt.Print("\033[43m\033[1;30m")
		case "Gray":
			fmt.Print("\033[40m\033[1;37m")
		}
		fmt.Printf(" %c ", c)
		fmt.Print("\033[m\033[m")
	}
	fmt.Println()
}

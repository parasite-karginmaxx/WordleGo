package logic

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const WORD_LENGTH = 10

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

func GetFilledSymbols(secretWord string, guess string) [WORD_LENGTH]string {
	color_vector := [WORD_LENGTH]string{}
	for i := range color_vector {
		color_vector[i] = "Gray"
	}
	// stores whether an index is allowed to cause another index to be yellow
	yellow_lock := [WORD_LENGTH]bool{}

	for j, guess_letter := range guess {
		for k, letter := range secretWord {
			if guess_letter == letter && j == k {
				color_vector[j] = "Green"
				// now the kth index can no longer cause another index to be yellow
				yellow_lock[k] = true
				break
			}
		}
	}
	for j, guess_letter := range guess {
		for k, letter := range secretWord {
			if guess_letter == letter && color_vector[j] != "Green" && !yellow_lock[k] {
				color_vector[j] = "Yellow"
				yellow_lock[k] = true
			}
		}
	}
	return color_vector
}

func (wg *WordleGame) Start() {
	fmt.Println("Добро пожаловать в игру Wordle!")
	fmt.Printf("Угадайте слово из 5 букв.\n")
	for {
		if wg.guesses >= wg.maxGuesses {
			fmt.Println("У вас закончились попытки. Загаданное слово было:")
			color_vector := GetFilledSymbols(wg.secretWord, wg.secretWord)
			DisplayWord(wg.secretWord, color_vector)
			return
		}

		fmt.Printf("Попытка #%d. Введите ваше предположение: ", wg.guesses+1)
		var guess string
		fmt.Scanln(&guess)
		guess = strings.ToLower(guess)

		if len(guess) != WORD_LENGTH {
			fmt.Printf("Пожалуйста, введите слово из 5 букв.\n")
			continue
		}

		wg.guesses++
		if guess == wg.secretWord {
			fmt.Println("Поздравляем! Вы угадали слово!")
			color_vector := GetFilledSymbols(wg.secretWord, guess)
			DisplayWord(wg.secretWord, color_vector)
			return
		} else {
			color_vector := GetFilledSymbols(wg.secretWord, guess)
			DisplayWord(guess, color_vector)
		}
	}
}

func ReadWordsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, scanner.Err()
}

func RandomWord(words []string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Uint64()
	return words[rand.Intn(len(words))]
}

func DisplayWord(word string, color_vector [WORD_LENGTH]string) {
	for i, c := range word {
		switch color_vector[i] {
		case "Green":
			fmt.Print("\033[42m\033[1;30m")
		case "Yellow":
			fmt.Print("\033[43m\033[1;30m")
		case "Grey":
			fmt.Print("\033[40m\033[1;37m")
		}
		fmt.Printf(" %c ", c)
		fmt.Print("\033[m\033[m")
	}
	fmt.Println()
}

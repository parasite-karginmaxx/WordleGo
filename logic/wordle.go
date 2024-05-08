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

// func (wg *WordleGame) hint(guess string) string {
// 	var builder strings.Builder
// 	for i, ch := range wg.secretWord {
// 		if strings.ContainsRune(guess, ch) {
// 			builder.WriteRune(ch)
// 		} else {
// 			builder.WriteRune('_')
// 		}
// 		if i < len(wg.secretWord)-1 {
// 			builder.WriteRune(' ')
// 		}
// 	}
// 	return builder.String()
// }

func (wg *WordleGame) Start() {
	fmt.Println("Добро пожаловать в игру Wordle!")
	// fmt.Println(wg.secretWord)
	fmt.Printf("Угадайте слово из %d букв.\n", len(wg.secretWord)/2)
	for {
		if wg.guesses >= wg.maxGuesses {
			fmt.Println("У вас закончились попытки. Загаданное слово было:")
			color_vector := get_filled_color_vector("Green")
			display_word(wg.secretWord, color_vector)
			return
		}

		fmt.Printf("Попытка #%d. Введите ваше предположение: ", wg.guesses+1)
		var guess string
		fmt.Scanln(&guess)
		guess = strings.ToLower(guess)

		if len(guess) != len(wg.secretWord) {
			fmt.Printf("Пожалуйста, введите слово из %d букв.\n", len(wg.secretWord)/2)
			continue
		}

		wg.guesses++
		if guess == wg.secretWord {
			fmt.Println("Поздравляем! Вы угадали слово!")
			color_vector := get_filled_color_vector("Green")
			display_word(wg.secretWord, color_vector)
			return
		} else {
			//fmt.Printf("Ваше предположение не верно. Подсказка: %s\n", wg.hint(guess))
			color_vector := get_filled_color_vector("Grey")

			// stores whether an index is allowed to cause another index to be yellow
			yellow_lock := [WORD_LENGTH]bool{}

			for j, guess_letter := range guess {
				for k, letter := range wg.secretWord {
					if guess_letter == letter && j == k {
						color_vector[j] = "Green"
						// now the kth index can no longer cause another index to be yellow
						yellow_lock[k] = true
						break

					}
				}
			}
			for j, guess_letter := range guess {
				for k, letter := range wg.secretWord {
					if guess_letter == letter && color_vector[j] != "Green" && !yellow_lock[k] {
						color_vector[j] = "Yellow"
						yellow_lock[k] = true
					}
				}
			}
			display_word(guess, color_vector)
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

func get_filled_color_vector(color string) [WORD_LENGTH]string {
	color_vector := [WORD_LENGTH]string{}
	for i := range color_vector {
		color_vector[i] = color
	}
	return color_vector
}

func display_word(word string, color_vector [WORD_LENGTH]string) {
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

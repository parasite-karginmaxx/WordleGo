package console

import (
	"fmt"
	"main/logic"
)

func Console_main() {
	words, err := logic.ReadWordsFromFile("logic/library.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	// Выбираем случайное слово из списка
	secretWord := logic.RandomWord(words)
	// Создаем новую игру Wordle
	game := logic.NewWordleGame(secretWord, 5)

	// Начинаем игру
	game.Start()
}

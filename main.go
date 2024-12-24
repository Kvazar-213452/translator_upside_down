package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func convertToUkrainian(text string) string {
	englishToUkrainian := map[rune]rune{
		'q': 'й', 'w': 'ц', 'e': 'у', 'r': 'к', 't': 'е', 'y': 'н', 'u': 'г', 'i': 'ш', 'o': 'щ', 'p': 'з',
		'[': 'х', ']': 'ї', '\\': 'є', 'a': 'ф', 's': 'і', 'd': 'в', 'f': 'а', 'g': 'п', 'h': 'р', 'j': 'о',
		'k': 'л', 'l': 'д', ';': 'ж', '\'': 'є', 'z': 'я', 'x': 'ч', 'c': 'с', 'v': 'м', 'b': 'и', 'n': 'т',
		'm': 'ь', ',': 'б', '.': 'ю', '/': 'ї', 'A': 'Ф', 'S': 'І', 'D': 'В', 'F': 'А', 'G': 'П', 'H': 'Р',
		'J': 'О', 'K': 'Л', 'L': 'Д', ':': 'Ж', '"': 'Є', 'Z': 'Я', 'X': 'Ч', 'C': 'С', 'V': 'М', 'B': 'И',
		'N': 'Т', 'M': 'Ь', '<': 'Б', '>': 'Ю', '?': 'Ї',
	}

	var result strings.Builder
	for _, char := range text {
		if ukChar, found := englishToUkrainian[char]; found {
			result.WriteRune(ukChar)
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input:")
	inputText, _ := reader.ReadString('\n')
	inputText = strings.TrimSpace(inputText)

	outputText := convertToUkrainian(inputText)
	fmt.Println("Результат українською мовою:")
	fmt.Println(outputText)
}

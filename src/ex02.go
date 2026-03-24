package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func TopKWords(text string, k int) []string {
	// Проверка на пустую строку
	if len(strings.TrimSpace(text)) == 0 {
		return []string{}
	}

	// Проверка на некорректное k
	if k <= 0 {
		return []string{}
	}

	words := strings.Fields(text)
	fmap := make(map[string]int)
	for _, word := range words {
		fmap[word]++
	}

	type WordFreq struct {
		word string
		freq int
	}

	WordFreqs := make([]WordFreq, 0, len(fmap))
	for word, freq := range fmap {
		WordFreqs = append(WordFreqs, WordFreq{word, freq})
	}

	sort.Slice(WordFreqs, func(i, j int) bool {
		if WordFreqs[i].freq == WordFreqs[j].freq {
			return WordFreqs[i].word < WordFreqs[j].word
		}
		return WordFreqs[i].freq > WordFreqs[j].freq
	})

	// Определяем размер результата
	resultSize := k
	if k > len(WordFreqs) {
		resultSize = len(WordFreqs)
	}

	result := make([]string, resultSize)
	for i := 0; i < resultSize; i++ {
		result[i] = WordFreqs[i].word
	}
	return result
}

func main2() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введи символы:")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	fmt.Println("Введи число k:")
	frek, _ := reader.ReadString('\n')
	frek = strings.TrimSpace(frek)

	k, err := strconv.Atoi(frek)
	if err != nil {
		fmt.Println("Ошибка: k должно быть числом")
		return
	}

	result := TopKWords(text, k)
	fmt.Println(strings.Join(result, " "))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main3() {
	reader := bufio.NewReader(os.Stdin)

	//fmt.Println("1 Список символов:")
	text1, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	text1 = strings.TrimSpace(text1)

	//fmt.Println("2 Список символов:")
	text2, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	text2 = strings.TrimSpace(text2)

	if text1 == "" || text2 == "" {
		fmt.Println("Invalid input")
		return
	}

	list1, err := parseInList(text1)
	if err != nil {
		fmt.Println("Invalid Input")
		return
	}

	list2, err := parseInList(text2)
	if err != nil {
		fmt.Println("Invalid Input")
		return
	}

	secondSet := make(map[int]bool)
	for _, num := range list2 {
		secondSet[num] = true
	}

	result := make([]int, 0)
	seen := make(map[int]bool)
	for _, num := range list1 {
		if secondSet[num] && !seen[num] {
			result = append(result, num)
			seen[num] = true
		}
	}

	if len(result) == 0 {
		fmt.Println("Empty intersection")
	} else {
		for i, num := range result {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(num)
		}
	}
	fmt.Println()
}

func parseInList(line string) ([]int, error) {
	parts := strings.Fields(line)
	result := make([]int, 0, len(parts))
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}
	return result, nil
}

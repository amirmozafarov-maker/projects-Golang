package main

import (
	"reflect"
	"testing"
)

func TestTopKWords(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		k        int
		expected []string
	}{
		{
			name:     "обычное поведение",
			text:     "aa bb cc aa cc cc cc aa ab ac bb",
			k:        3,
			expected: []string{"cc", "aa", "bb"},
		},
		{
			name:     "пустая строка",
			text:     "",
			k:        5,
			expected: []string{},
		},
		{
			name:     "строка с пробелами (пустая после TrimSpace)",
			text:     "   ",
			k:        5,
			expected: []string{},
		},
		{
			name:     "K больше уникальных слов",
			text:     "apple banana apple",
			k:        5,
			expected: []string{"apple", "banana"},
		},
		{
			name:     "одинаковая частота, лексикографический порядок",
			text:     "cat dog bird cat dog bird",
			k:        3,
			expected: []string{"bird", "cat", "dog"},
		},
		{
			name:     "одно слово",
			text:     "hello hello hello",
			k:        1,
			expected: []string{"hello"},
		},
		{
			name:     "слова с разной частотой",
			text:     "a b b c c c d d d d",
			k:        4,
			expected: []string{"d", "c", "b", "a"},
		},
		{
			name:     "K равно 0",
			text:     "some words here",
			k:        0,
			expected: []string{},
		},
		{
			name:     "K отрицательное",
			text:     "hello world",
			k:        -1,
			expected: []string{},
		},
		{
			name:     "слова с пробелами в начале и конце",
			text:     "  test test  word  ",
			k:        2,
			expected: []string{"test", "word"},
		},
		{
			name:     "большой текст",
			text:     "the quick brown fox jumps over the lazy dog the quick brown fox",
			k:        3,
			expected: []string{"the", "brown", "fox"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TopKWords(tt.text, tt.k)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TopKWords(%q, %d) = %v, want %v", tt.text, tt.k, result, tt.expected)
			}
		})
	}
}

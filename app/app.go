package app

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"kazdream/models"
	"sort"
	"unicode"
)

func App() {
	data, err := ioutil.ReadFile(models.FileName)
	if err != nil {
		return
	}
	words := bytes.FieldsFunc(data, func(c rune) bool {
		return !unicode.IsLetter(c)
	})
	toLower(words)
	sortByteSlices(words)
	result := countWords(words)
	sort.Slice(result, func(i, j int) bool {
		return result[i].Amount > result[j].Amount
	})
	printWords(result)
}

func printWords(words []*models.Word) {
	for i, word := range words {
		if i == models.AmountWordsForPrint {
			break
		}
		fmt.Printf("% 8d %s\n", word.Amount, word.Value)
	}
}

func sortByteSlices(words [][]byte) {
	sort.Slice(words, func(i, j int) bool {
		return bytes.Compare(words[i], words[j]) < 0
	})
}

func toLower(words [][]byte) {
	for i := range words {
		words[i] = bytes.ToLower(words[i])
	}
}

func countWords(words [][]byte) []*models.Word {
	wrodsWithAmount := []*models.Word{}
	count := 1
	for i := range words {
		if i == len(words)-1 {
			break
		}
		if bytes.Compare(words[i], words[i+1]) == 0 {
			// if bytes.Equal(words[i], words[i+1]) {
			count++
		} else {
			w := &models.Word{
				Value:  words[i],
				Amount: count,
			}
			wrodsWithAmount = append(wrodsWithAmount, w)
			count = 1
		}
	}
	return wrodsWithAmount
}

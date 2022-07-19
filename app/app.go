package app

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"kazdream/models"
	"sort"
	"unicode"
)

const (
	fileName            = "mobydick.txt"
	AmountWordsForPrint = 20
)

func App() {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	words := bytes.FieldsFunc(data, func(c rune) bool {
		return !unicode.IsLetter(c)
	})
	toLower(words)
	SortByteSlices(words)
	result := CountWords(words)
	sort.Slice(result, func(i, j int) bool {
		return result[i].Amount > result[j].Amount
	})
	PrintWords(result)
}

func PrintWords(words []*models.Word) {
	for i, word := range words {
		if i == AmountWordsForPrint {
			break
		}
		sp := "   "
		if word.Amount < 1000 {
			sp = "    "
		}
		fmt.Printf("%s%d %s\n", sp, word.Amount, word.Value)
	}
}

func SortByteSlices(words [][]byte) {
	sort.Slice(words, func(i, j int) bool { return bytes.Compare(words[i], words[j]) < 0 })
}

func toLower(words [][]byte) {
	for i := range words {
		words[i] = bytes.ToLower(words[i])
	}
}

func CountWords(words [][]byte) []*models.Word {
	var wrodsWithAmount []*models.Word
	count := 1
	for i := range words {
		if i != len(words)-1 {
			// if bytes.Compare(words[i], words[i+1]) == 0 {
			if bytes.Equal(words[i], words[i+1]) {
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
	}
	return wrodsWithAmount
}

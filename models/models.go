package models

const (
	FileName            = "mobydick.txt"
	AmountWordsForPrint = 20
)

type Word struct {
	Value  []byte
	Amount int
}

// type Words []*Word

// func (v Words) Len() int           { return len(v) }
// func (v Words) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
// func (v Words) Less(i, j int) bool { return v[i].Amount > v[j].Amount }

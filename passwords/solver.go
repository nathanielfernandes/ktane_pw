package passwords

var /* const */ WORDS = []string{
	"about",
	"after",
	"again",
	"below",
	"could",
	"every",
	"first",
	"found",
	"great",
	"house",
	"large",
	"learn",
	"never",
	"other",
	"place",
	"plant",
	"point",
	"right",
	"small",
	"sound",
	"spell",
	"still",
	"study",
	"their",
	"there",
	"these",
	"thing",
	"think",
	"three",
	"water",
	"where",
	"which",
	"world",
	"would",
	"write",
}

func filterWords(words []string, letters [6]rune, col int) []string {
	if letters == [6]rune{' ', ' ', ' ', ' ', ' ', ' '} {
		return words
	}

	filtered := []string{}
	for _, word := range words {
		for _, c := range letters {
			if c == rune(word[col]) {
				filtered = append(filtered, word)
			}
		}
	}

	return filtered
}

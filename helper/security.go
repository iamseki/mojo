package helper

/*
'A' || 'M' returns 1 which means January, 'B' || 'N' returns 2 which means Fabruary an so on...

The codesheet to decode option ticker is available at: https://www.investopedia.com/ask/answers/05/052505.asp
*/
func MonthByLetterInOptionTickr(letter rune) int {
	monthByLetter := map[rune]int{
		'A': 1,
		'M': 1,
		'B': 2,
		'N': 2,
		'C': 3,
		'O': 3,
		'D': 4,
		'P': 4,
		'E': 5,
		'Q': 5,
		'F': 6,
		'R': 6,
		'G': 7,
		'S': 7,
		'H': 8,
		'T': 8,
		'I': 9,
		'U': 9,
		'J': 10,
		'V': 10,
		'K': 11,
		'W': 11,
		'L': 12,
		'X': 12,
	}

	return monthByLetter[letter]
}

package mmi

import (
	"math/rand"
	"time"
)

var funnyRandomSentences = map[int]string{
	0:  "NASA actually has an Office of Planetary Protection in case life is found on another planet.",
	1:  "“NASA” stands for National Aeronautics and Space Administration.",
	2:  "Never have I ever done a duck face selfie.",
	3:  "Before software can be reusable it first has to be usable.",
	4:  "It’s not a bug – it’s an undocumented feature.",
	5:  "Programming is like sex. One mistake and you have to support it for the rest of your life.",
	6:  "Algorithm: Word used by programmers when they don’t want to explain what they did.",
	7:  "Walking on water and developing software from a specification are easy if both are frozen.",
	8:  "Programming made the impossible possible. You can have a null object and a constant variable.",
	9:  "In C we had to code our own bugs. In C++ we can inherit them.",
	10: "Why do Java programmers have to wear glasses? Because they don’t C#.",
	11: "A SQL query goes into a bar, walks up to two tables, and asks, ‘Can I join you?",
	12: "Computers are good at following instructions but not at reading your mind.",
	13: "When all else fails … reboot.",
}

func GetRandomSentence() string {
	rand.Seed(time.Now().UnixNano())

	min := 0
	max := 13
	num := rand.Intn(max-min+1) + min

	return funnyRandomSentences[num]
}

func GetFunnySentence(n int) string {
	return funnyRandomSentences[n]
}

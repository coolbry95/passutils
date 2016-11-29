// Package rules contains all of the functions for hashcat rules
package rules

//TODO add documentation
// https://godoc.org/github.com/coolbry95/passutils/ruleprocessor/rules

import (
	"unicode"
)

// reason for var temp []rune and copy(temp, word)
// is slices are passed by value but reference their underlying array
// if a slice is passed and is manipulated like word[i] = 'j' then
// the word variable passed gets changed also

// may change to not using copy
// then changes would be applied to underlying slice

// global memorized word
var Saved []rune

// ToNum converts A-Z to ints
func ToNum(char rune) int {
	if char <= 57 && char >= 48 {
		return int(char) - 48
	}
	return int(char) - 65 + 10
}

func ToNumByte(num uint8) int {
	return ToNum(rune(num))
}

func ToAlpha(num int) rune {
	if num < 10 {
		return rune(48 + num)
	}
	return rune(65 + num - 10)
}

// :
func Nothing(word []rune) []rune {
	return word
}

// l
func Lowercase(word []rune) []rune {
	if len(word) == 0 {
		return word
	}
	// keep strings.Tolower() and then convert back to rune?
	// only operates on single rune
	for i, v := range word {
		word[i] = unicode.ToLower(v)
	}
	return word
}

// u
func Uppercase(word []rune) []rune {
	if len(word) == 0 {
		return word
	}
	for i, v := range word {
		word[i] = unicode.ToUpper(v)
	}
	return word
}

// c
func Capitalize(word []rune) []rune {
	if len(word) == 0 {
		return word
	}
	for i, v := range word {
		if i == 0 {
			word[i] = unicode.ToUpper(v)
			continue
		}
		word[i] = unicode.ToLower(v)
	}
	return word
}

// C
func InvertCapitalize(word []rune) []rune {
	if len(word) == 0 {
		return word
	}
	for i, v := range word {
		if i == 0 {
			word[i] = unicode.ToLower(v)
			continue
		}
		word[i] = unicode.ToUpper(v)
	}
	return word
}

// t
func ToggleCase(word []rune) []rune {
	if len(word) == 0 {
		return word
	}
	for i, v := range word {
		switch {
		case unicode.IsUpper(v):
			word[i] = unicode.ToLower(v)
		case unicode.IsLower(v):
			word[i] = unicode.ToUpper(v)
		default:
			word[i] = v
		}
	}
	return word
}

// TN
func ToggleAt(word []rune, n int) []rune {
	if len(word) == 0 {
		return word
	}
	if len(word)-1 < n {
		return word
	}
	switch {
	case unicode.IsUpper(word[n]):
		word[n] = unicode.ToLower(word[n])
	case unicode.IsLower(word[n]):
		word[n] = unicode.ToUpper(word[n])
	default:
		return word
	}
	return word
}

// https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
// r
func Reverse(word []rune) []rune {
	if len(word) == 0 {
		return word
	}
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		word[i], word[j] = word[j], word[i]
	}
	return word
}

// d
func Duplicate(word []rune) []rune {
	if len(word) == 0 {
		return word
	}
	temp := make([]rune, len(word), len(word)*2)
	copy(temp, word)
	temp = append(temp, word[:]...)
	return temp
}

// pN
func DuplicateN(word []rune, n int) []rune {
	if len(word) == 0 {
		return word
	}
	temp := make([]rune, 0, len(word)*n)
	for i := 0; i <= n; i++ {
		temp = append(temp, word[:]...)
	}
	return temp
}

// f didnt change yet
func Reflect(word []rune) []rune {
	if len(word) == 0 {
		return word
	}
	firstHalf := make([]rune, len(word))
	copy(firstHalf, word)
	temp := make([]rune, len(word), len(word))
	for i, j := len(word)-1, 0; i >= 0; i, j = i-1, j+1 {
		temp[j] = firstHalf[i]
	}
	firstHalf = append(firstHalf, temp[:]...)
	return firstHalf
}

// {
func RotateLeft(word []rune) []rune {
	if len(word) == 0 {
		return word
	}
	if len(word)-1 < 0 {
		return word
	}

	temp := make([]rune, 0, len(word))
	temp = append(temp, word[1:]...)
	temp = append(temp, word[0])

	return temp
}

// }
func RotateRight(word []rune) []rune {
	if len(word) == 0 {
		return word
	}
	if len(word)-1 == 0 {
		return word
	}
	temp := make([]rune, 0, len(word))
	temp = append(temp, word[len(word)-1])
	temp = append(temp, word[:len(word)-1]...)
	return temp
}

// $X
func AppendCharacter(word []rune, char rune) []rune {
	if len(word) == 0 {
		return word
	}
	word = append(word, char)
	return word
}

// ^X
func PrependCharacter(word []rune, char rune) []rune {
	if len(word) == 0 {
		return word
	}
	temp := make([]rune, 1, len(word)+1)

	temp[0] = char

	temp = append(temp, word[:]...)
	return temp
}

// [
func TruncateLeft(word []rune) []rune {
	if len(word) == 0 {
		return word
	}
	return word[1:]
}

// ]
func TruncateRight(word []rune) []rune {
	if len(word) == 0 {
		return word
	}
	return word[:len(word)-1]
}

// DN
func DeleteN(word []rune, n int) []rune {
	if len(word) == 0 {
		return word
	}
	if len(word)-1 < n {
		return word
	}
	/*
		temp := make([]rune, 0, len(word)-n)
		temp = append(temp, word[:n]...)
		temp = append(temp, word[n+1:]...)
	*/
	word = append(word[:n], word[n+1:]...)
	return word
}

// xNM
func ExtractRange(word []rune, n, m int) []rune {
	if len(word) == 0 {
		return word
	}

	if n > len(word) || m > len(word) {
		return nil
	}

	word = word[n : m+n]

	return word
}

// ONM
func OmitRange(word []rune, n, m int) []rune {
	if len(word) == 0 {
		return word
	}
	if n > len(word)-1 || m > len(word)-1 || n == m {
		return word
	}
	// no idea about this now lol
	// FIXME
	// m delete the entire word
	/*
		if m == len(word) -1 {
			return ""
		}
	*/
	/*
		temp := make([]rune, 0, n+m)
		temp = append(temp, word[:n]...)
		temp = append(temp, word[m+n:]...)
	*/
	word = append(word[:n], word[m+n:]...)
	return word
}

// iNX
func InsertAtN(word []rune, n int, char rune) []rune {
	if len(word) == 0 {
		return word
	}
	// TODO breaks this function!?! if it is len(word)-1
	if len(word) < n {
		return word
	}

	temp := make([]rune, 0, len(word)+1)
	temp = append(temp, word[:n]...)
	temp = append(temp, char)
	temp = append(temp, word[n:]...)
	//word = append(word[:n], char, word[n:]...)
	//word = append(word[:n], append([]rune{char}, word[n:]...)...)
	return temp
}

// oNX
func OverwriteAtN(word []rune, n int, char rune) []rune {
	if len(word) == 0 {
		return word
	}
	if len(word)-1 < n {
		return word
	}
	/*
		temp := make([]rune, 0, len(word)+1)
		temp = append(temp, word[:n]...)
		temp = append(temp, char)
		temp = append(temp, word[n+1:]...)
	*/
	word[n] = char
	return word
}

// 'N
func TruncateAtN(word []rune, n int) []rune {
	if len(word) == 0 {
		return word
	}
	if len(word)-1 < n {
		return word
	}
	// dont know here either lol
	// FIXME
	if n == 0 {
		return word
	}
	// or here
	// FIXME
	/*
		temp := make([]rune, len(word))
		copy(temp, word)
		return temp[:n]
	*/
	word = word[:n]
	return word
}

// sXY
func Replace(word []rune, x, y rune) []rune {
	if len(word) == 0 {
		return word
	}

	for i := range word {
		if word[i] == x {
			word[i] = y
		}
	}
	return word
}

// @X
func Purge(word []rune, x rune) []rune {
	if len(word) == 0 {
		return word
	}

	// use cap len(word) because worse case is nothing is purged
	temp := make([]rune, 0, len(word))
	for i := 0; i < len(word); i++ {
		if word[i] == x {
			continue
		}
		temp = append(temp, word[i])
	}
	return temp
}

// zN
func DuplicateFirstN(word []rune, n int) []rune {
	if len(word) == 0 {
		return word
	}
	temp := make([]rune, 0, len(word)+n)

	for i := 0; i < n; i++ {
		temp = append(temp, word[0])
	}

	temp = append(temp, word[:]...)
	return temp
}

// ZN
func DuplicateLastN(word []rune, n int) []rune {
	if len(word) == 0 {
		return word
	}

	temp := make([]rune, len(word), len(word)+n)

	copy(temp, word)

	for i := 0; i < n; i++ {
		temp = append(temp, word[len(word)-1])
	}

	return temp
}

// q
func DuplicateAll(word []rune) []rune {
	if len(word) == 0 {
		return word
	}
	temp := make([]rune, 0, len(word)*2)
	for i := 0; i < len(word); i++ {
		temp = append(temp, word[i])
		temp = append(temp, word[i])
	}
	return temp
}

// XNMI
func ExtractMemory(word []rune, n, m, i int) []rune {
	if len(word) == 0 {
		return word
	}
	// minus one because they will be indexing the saved word
	if n > len(Saved) || m > len(Saved) || i > len(word) {
		return nil
	}

	temp := make([]rune, len(word[:i]), len(word)+(n+m))
	//temp = append(temp, word[:i]...)
	copy(temp, word[:i])
	temp = append(temp, Saved[n:m+n]...)
	temp = append(temp, word[i:]...)

	return temp
}

// 4
func AppendMemory(word []rune) []rune {
	if len(word) == 0 {
		return word
	}
	//temp := make([]rune, 0, len(word)+len(Saved))
	// use copy instead??
	// append to word instead of making new variable

	/*
		temp = append(temp, word[:]...)
		temp = append(temp, Saved[:]...)
	*/
	word = append(word, Saved[:]...)
	return word
}

// 6
func PrependMemory(word []rune) []rune {
	if len(word) == 0 {
		return word
	}

	temp := make([]rune, 0, len(word)+len(Saved))
	// use copy instead??
	temp = append(temp, Saved[:]...)
	temp = append(temp, word[:]...)

	return temp
}

// M
func Memorize(word []rune) {
	// dont know if you need
	// Saved = copy(Saved, word[:])
	// is this possible Saved = word
	//copy(Saved, word)
	//Saved = word
	Saved = make([]rune, len(word), len(word))
	copy(Saved, word)
}

// reject greater and less use the 0-9 a-z
// need to change becuase hashcat does not use this
// <N
func RejectLess(word []rune, n int) bool {
	if len(word) > n {
		return true
	}

	return false
}

// >N
func RejectGreater(word []rune, n int) bool {
	if len(word) < n {
		return true
	}

	return false
}

// use index instead?
// !X
func RejectContain(word []rune, char rune) bool {
	for _, v := range word {
		if v == char {
			return true
		}
	}
	return false
}

// /X
func RejectNotContain(word []rune, char rune) bool {
	for _, v := range word {
		if v == char {
			return false
		}
	}
	return true
}

// (X
func RejectEqualFirst(word []rune, char rune) bool {
	if word[0] != char {
		return true
	}

	return false
}

// )X
func RejectEqualLast(word []rune, char rune) bool {
	if word[len(word)-1] != char {
		return true
	}

	return false
}

// =NX
func RejectEqualAt(word []rune, char rune, n int) bool {
	if word[n] != char {
		return true
	}

	return false
}

// %NX
func RejectContains(word []rune, char rune, n int) bool {
	count := 0
	for _, v := range word {
		if v == char {
			count++
		}
	}
	if count < n {
		return true
	}

	return false
}

// Q
func RejectMemory(word []rune) bool {
	if word == nil && Saved == nil {
		return false
	}
	if word == nil || Saved == nil {
		return false
	}
	if len(word) < len(Saved) {
		return false
	}

	for i := range word {
		if word[i] != Saved[i] {
			return false
		}
	}

	return true
}

// k
func SwapFront(word []rune) []rune {
	if len(word) == 1 || len(word) == 0 {
		return word
	}
	word[0], word[1] = word[1], word[0]
	return word
}

// K
func SwapBack(word []rune) []rune {
	if len(word) == 1 || len(word) == 0 {
		return word
	}

	word[len(word)-1], word[len(word)-2] = word[len(word)-2], word[len(word)-1]

	return word
}

// *XY
func SwapAtN(word []rune, x, y int) []rune {
	if len(word) == 0 {
		return word
	}
	if x == 0 || y == 0 {
		return word
	}
	if len(word)-1 <= x || len(word)-1 < y {
		return word
	}

	word[x], word[y] = word[y], word[x]
	return word
}

// LN
func BitwiseShiftLeft(word []rune, n int) []rune {
	if len(word) == 0 {
		return word
	}

	word[n] = word[n] << 1

	return word
}

// RN
func BitwiseShiftRight(word []rune, n int) []rune {
	if len(word) == 0 {
		return word
	}
	word[n] = word[n] >> 1
	return word
}

// +N
func ASCIIIncrementPlus(word []rune, n int) []rune {
	if len(word) == 0 {
		return word
	}
	word[n] = word[n] + 1
	return word
}

// -N
func ASCIIIncrementMinus(word []rune, n int) []rune {
	if len(word) == 0 {
		return word
	}
	word[n] = word[n] - 1
	return word
}

// .N
func ReplaceNPlus(word []rune, n int) []rune {
	if len(word) == 0 {
		return word
	}
	if n >= len(word)-1 {
		return word
	}
	/*
		temp := make([]rune, 0, len(word))
		temp = append(temp, word[:n]...)
		temp = append(temp, word[n+1])
		temp = append(temp, word[n+1:]...)
	*/
	word[n] = word[n+1]
	return word
}

// ,N
func ReplaceNMinus(word []rune, n int) []rune {
	if len(word) == 0 {
		return word
	}
	if n > len(word)-1 {
		return word
	}
	// FIXME
	// index error on 0
	if n == 0 {
		return word
	}
	temp := make([]rune, 0, len(word))
	temp = append(temp, word[:n]...)
	temp = append(temp, word[n-1])
	temp = append(temp, word[n+1:]...)
	return temp
}

// yN
func DuplicateBlockFront(word []rune, n int) []rune {
	if len(word) == 0 {
		return word
	}

	temp := make([]rune, len(word[:n]), len(word)+n)
	copy(temp, word[:n])
	temp = append(temp, word[:]...)

	return temp
}

// YN
func DuplicateBlockBack(word []rune, n int) []rune {
	if len(word) == 0 {
		return word
	}

	temp := make([]rune, len(word), len(word)+n)
	copy(temp, word)
	temp = append(temp, word[len(word)-n:]...)

	return temp
}

// E
func Title(word []rune) []rune {
	if len(word) == 0 {
		return word
	}
	word[0] = unicode.ToUpper(word[0])
	for i := 1; i <= len(word)-1; i++ {
		if word[i] == ' ' {
			/*
				if i == len(word)-1 {
					break
				}
			*/
			word[i+1] = unicode.ToUpper(word[i+1])
			i++
		} else {
			word[i] = unicode.ToLower(word[i])
		}
	}
	return word
	/*
		temp := make([]rune, len(word))
		copy(temp, word)
		temp[0] = unicode.ToUpper(temp[0])
		for i := 1; i <= len(temp)-1; i++ {
			if temp[i] == ' ' {
				//
					if i == len(word)-1 {
						break
					}
				//
				temp[i+1] = unicode.ToUpper(temp[i+1])
				i++
			} else {
				temp[i] = unicode.ToLower(temp[i])
			}
		}
		return temp
	*/
}

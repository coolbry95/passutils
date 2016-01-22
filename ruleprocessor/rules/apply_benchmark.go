package rules

import (
	"testing"
)

//Saved = []rune{'a', 'a', 'v', 'w', 'l', 'o'. 't', 'r', 'w'}

var result int

//func ToNum(char uint8) int {
func BenchmarkToNum(b *testing.B) {
	var r int
	// read comment about single word
	// same applies here
	var i uint8 = 12
	for n:=0; n < b.N; n++ {
		r = ToNumByte(i)
	}

	result = r
}

// dont really know how else i should do this
// just going to use this one word for benchmarking
// maybe have a slice of slices to do more realistic benchmarking?
var mangled []rune
var plain = []rune{'g', 'o', 'l', 'a','n','g'}

//func Nothing(word []rune) []rune {
func BenchmarkNothing(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = Nothing(plain)
	}

	mangled = r
}

//func Lowercase(word []rune) []rune {
func BenchmarkLowercase(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = Lowercase(plain)
	}

	mangled = r
}

//func Uppercase(word []rune) []rune {
func BenchmarkUppercase(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = Uppercase(plain)
	}

	mangled = r
}

//func Capitalize(word []rune) []rune {
func BenchmarkCapitalize(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = Capitalize(plain)
	}

	mangled = r
}

//func InvertCapitalize(word []rune) []rune {
func BenchmarkInvertCapitalize(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = InvertCapitalize(plain)
	}

	mangled = r
}

//func ToggleCase(word []rune) []rune {
func BenchmarkToggleCase(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = ToggleCase(plain)
	}

	mangled = r
}

//func ToggleAt(word []rune, n int) []rune {
func BenchmarkToggleAt(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = ToggleAt(plain, 2)
	}

	mangled = r
}

//func Reverse(word []rune) []rune {
func BenchmarkReverse(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = Reverse(plain)
	}

	mangled = r
}

//func Duplicate(word []rune) []rune {
func BenchmarkDuplicate(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = Duplicate(plain)
	}

	mangled = r
}

//func DuplicateN(word []rune, n int) []rune {
func BenchmarkDuplicateN(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = DuplicateN(plain, 2)
	}

	mangled = r
}

//func Reflect(word []rune) []rune {
func BenchmarkReflect(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = Reflect(plain)
	}

	mangled = r
}

//func RotateLeft(word []rune) []rune {
func BenchmarkRotateLeft(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = RotateLeft(plain)
	}

	mangled = r
}

//func RotateRight(word []rune) []rune {
func BenchmarkRotateRight(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = RotateRight(plain)
	}

	mangled = r
}

//func AppendCharacter(word []rune, char rune) []rune {
func BenchmarkAppendCharacter(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = AppendCharacter(plain, 'a')
	}

	mangled = r
}

//func PrependCharacter(word []rune, char rune) []rune {
func BenchmarkPrependCharacter(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = PrependCharacter(plain, 'a')
	}

	mangled = r
}

//func TruncateLeft(word []rune) []rune {
func BenchmarkTruncateLeft(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = TruncateLeft(plain)
	}

	mangled = r
}

//func TruncateRight(word []rune) []rune {
func BenchmarkTruncateRight(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = TruncateRight(plain)
	}

	mangled = r
}

//func DeleteN(word []rune, n int) []rune {
func BenchmarkDeleteN(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = DeleteN(plain, 2)
	}

	mangled = r
}
//func ExtractRange(word []rune, n, m int) []rune {
func BenchmarkExtractRange(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = ExtractRange(plain, 1, 3)
	}

	mangled = r
}

//func OmitRange(word []rune, n, m int) []rune {
func BenchmarkOmitRange(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = OmitRange(plain, 1, 3)
	}

	mangled = r
}

//func InsertAtN(word []rune, n int, char rune) []rune {
func BenchmarkInsertAtN(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = InsertAtN(plain, 2, 'a')
	}

	mangled = r
}

//func OverwriteAtN(word []rune, n int, char rune) []rune {
func BenchmarkOverwriteAtN(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = OverwriteAtN(plain, 2, 'a')
	}

	mangled = r
}

//func TruncateAtN(word []rune, n int) []rune {
func BenchmarkTruncateAtN(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = TruncateAtN(plain, 2)
	}

	mangled = r
}

//func Replace(word []rune, x, y rune) []rune {
func BenchmarkReplace(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = Replace(plain, 'o', 'a')
	}

	mangled = r
}

//func Purge(word []rune, x rune) []rune {
func BenchmarkPurge(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = Purge(plain, 'o')
	}

	mangled = r
}

//func DuplicateFirstN(word []rune, n int) []rune {
func BenchmarkDuplicateFirstN(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = DuplicateFirstN(plain, 2)
	}

	mangled = r
}

//func DuplicateLastN(word []rune, n int) []rune {
func BenchmarkDuplicateLastN(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = DuplicateLastN(plain, 2)
	}

	mangled = r
}

//func DuplicateAll(word []rune) []rune {
func BenchmarkDuplicateAll(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = DuplicateAll(plain)
	}

	mangled = r
}

//func ExtractMemory(word []rune, n, m, i int) []rune {
func BenchmarkExtractMemory(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = ExtractMemory(plain, 1, 2, 3)
	}

	mangled = r
}

//func AppendMemory(word []rune) []rune {
func BenchmarkAppendMemory(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = AppendMemory(plain)
	}

	mangled = r
}

//func PrependMemory(word []rune) []rune {
func BenchmarkPrependMemory(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = PrependMemory(plain)
	}

	mangled = r
}

//func Memorize(word []rune) {
func BenchmarkMemorize(b *testing.B) {
	for n:=0; n < b.N; n++ {
		Memorize(plain)
	}
}

var truth bool

//func RejectLess(word []rune, n int) bool {
func BenchmarkRejectLess(b *testing.B) {
	var r bool

	for n:=0; n < b.N; n++ {
		r = RejectLess(plain, 2)
	}

	truth = r
}

//func RejectGreater(word []rune, n int) bool {
func BenchmarkRejectGreater(b *testing.B) {
	var r bool

	for n:=0; n < b.N; n++ {
		r = RejectGreater(plain, 2)
	}

	truth = r
}

//func RejectContain(word []rune, char rune) bool {
func BenchmarkRejectContain(b *testing.B) {
	var r bool

	for n:=0; n < b.N; n++ {
		r = RejectContain(plain, 'a')
	}

	truth = r
}

//func RejectNotContain(word []rune, char rune) bool {
func BenchmarkRejectNotContain(b *testing.B) {
	var r bool

	for n:=0; n < b.N; n++ {
		r = RejectNotContain(plain, 'a')
	}

	truth = r
}

//func RejectEqualFirst(word []rune, char rune) bool {
func BenchmarkRejectEqualFirst(b *testing.B) {
	var r bool

	for n:=0; n < b.N; n++ {
		r = RejectEqualFirst(plain, 'g')
	}

	truth = r
}

//func RejectEqualLast(word []rune, char rune) bool {
func BenchmarkRejectEqualLast(b *testing.B) {
	var r bool

	for n:=0; n < b.N; n++ {
		r = RejectEqualLast(plain, 'g')
	}

	truth = r
}

//func RejectEqualAt(word []rune, char rune, n int) bool {
func BenchmarkRejectEqualAt(b *testing.B) {
	var r bool

	for n:=0; n < b.N; n++ {
		r = RejectEqualAt(plain, 'g', 0)
	}

	truth = r
}

//func RejectContains(word []rune, char rune, n int) bool {
func BenchmarkRejectContains(b *testing.B) {
	var r bool

	for n:=0; n < b.N; n++ {
		r = RejectContains(plain, 'g', 1)
	}

	truth = r
}

//func RejectMemory(word []rune) bool {
func BenchmarkRejectMemory(b *testing.B) {
	var r bool

	for n:=0; n < b.N; n++ {
		r = RejectMemory(plain)
	}

	truth = r
}

//func SwapFront(word []rune) []rune {
func BenchmarkSwapFront(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = SwapFront(plain)
	}

	mangled = r
}

//func SwapBack(word []rune) []rune {
func BenchmarkSwapBack(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = SwapBack(plain)
	}

	mangled = r
}

//func SwapAtN(word []rune, x, y int) []rune {
func BenchmarkSwapAtN(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = SwapAtN(plain, 2, 0)
	}

	mangled = r
}

//func BitwiseShiftLeft(word []rune, n int) []rune {
func BenchmarkBitwiseShiftLeft(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = BitwiseShiftLeft(plain, 2)
	}

	mangled = r
}

//func BitwiseShiftRight(word []rune, n int) []rune {
func BenchmarkBitwiseShiftRight(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = BitwiseShiftRight(plain, 2)
	}

	mangled = r
}

//func ASCIIIncrementPlus(word []rune, n int) []rune {
func BenchmarkASCIIIncrementPlus(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = ASCIIIncrementPlus(plain, 2)
	}

	mangled = r
}

//func ASCIIIncrementMinus(word []rune, n int) []rune {
func BenchmarkASCIIIncrementMinus(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = ASCIIIncrementMinus(plain, 2)
	}

	mangled = r
}

//func ReplaceNPlus(word []rune, n int) []rune {
func BenchmarkReplaceNPlus(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = ReplaceNPlus(plain, 2)
	}

	mangled = r
}

//func ReplaceNMinus(word []rune, n int) []rune {
func BenchmarkReplaceNMinus(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = ReplaceNMinus(plain, 2)
	}

	mangled = r
}

//func DuplicateBlockFront(word []rune, n int) []rune {
func BenchmarkDuplicateBlockFront(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = DuplicateBlockFront(plain, 2)
	}

	mangled = r
}

//func DuplicateBlockBack(word []rune, n int) []rune {
func BenchmarkDuplicateBlockBack(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = DuplicateBlockBack(plain, 2)
	}

	mangled = r
}

//func Title(word []rune) []rune {
func BenchmarkTitle(b *testing.B) {
	var r []rune

	for n:=0; n < b.N; n++ {
		r = Title(plain)
	}

	mangled = r
}


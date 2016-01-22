package rules

import (
	"testing"
)

// should change tests to not be so static
// add in rule to struct
func compare(a, b []rune) bool {
	if a == nil && b == nil {
		return false
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return true
	}
	for i := range a {
		if a[i] != b[i] {
			return true
		}
	}

	return false
}

// only rule without data
// is special anyway
//func TestExtractMemory(t *testing.T) {

// converts A-Z to ints
func TestToNum(t *testing.T) {
	alpha := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	number := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	first := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	second := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36}

	for index, num := range number {
		v := ToNum(num)
		if v != first[index] {
			t.Errorf("expected %d, got %d", first[index], v)
		}
	}
	for index, num := range alpha {
		v := ToNum(num)
		if v != second[index] {
			t.Errorf("expected %d, got %d", second[index], v)
		}
	}

}

// :
func TestNothing(t *testing.T) {
	//:
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', '3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'm', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 's', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 'i', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 'i', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'u', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := Nothing(tt.in)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// l
func TestLowercase(t *testing.T) {
	//l
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', '3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'm', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 's', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'d', 'o', 'r', 'i', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 'i', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'u', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := Lowercase(tt.in)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// u
func TestUppercase(t *testing.T) {
	//u
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'O', 'W', 'N'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'S', 'P', '3', '3', 'D'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'R', 'O', 'E', 'A', 'R', 'O', 'E', 'A'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'Y', 'R', 'A', 'L', 'L', 'I', 'H'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'G', 'E', 'M', 'M', 'A', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'L', 'O', 'V', 'E', '4', 'E', 'V', 'E', 'R', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'T', 'E', 'L', 'E', 'P', 'H', 'O', 'N', 'E', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'V', 'N', 'V', 'N', 'V', 'N'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'F', 'D', 'S', '2', 'F', 'D'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'A', 'M', 'U', 'R'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'S', 'O', 'C', 'C', 'E', 'R', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'P', 'A', 'R', 'O', 'A', 'C', 'H'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'O', 'R', 'I', 'T', 'O', 'S'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'S', 'K', 'I', 'P', 'E', 'R'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'T', 'H', 'X', '1', 'M', 'A', 'N'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'C', 'H', 'U', 'L', 'A', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 'T', 'R', 'E', 'B', 'L', 'A'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'N', 'A', 'N', 'C', 'Y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := Uppercase(tt.in)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// c
func TestCapitalize(t *testing.T) {
	//c
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'O', 'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'S', 'p', '3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'R', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'Y', 'r', 'a', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'G', 'e', 'm', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'L', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'T', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'V', 'n', 'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'F', 'd', 's', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'S', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'P', 'a', 'r', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 'i', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'S', 'k', 'i', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'T', 'h', 'x', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'C', 'h', 'u', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'N', 'a', 'n', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := Capitalize(tt.in)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// C
func TestInvertCapitalize(t *testing.T) {
	//C
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'W', 'N'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'P', '3', '3', 'D'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'O', 'E', 'A', 'R', 'O', 'E', 'A'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'R', 'A', 'L', 'L', 'I', 'H'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'E', 'M', 'M', 'A', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'O', 'V', 'E', '4', 'E', 'V', 'E', 'R', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'E', 'L', 'E', 'P', 'H', 'O', 'N', 'E', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'N', 'V', 'N', 'V', 'N'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'D', 'S', '2', 'F', 'D'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'A', 'M', 'U', 'R'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'O', 'C', 'C', 'E', 'R', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'A', 'R', 'O', 'A', 'C', 'H'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'d', 'O', 'R', 'I', 'T', 'O', 'S'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'K', 'I', 'P', 'E', 'R'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'H', 'X', '1', 'M', 'A', 'N'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'H', 'U', 'L', 'A', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 'T', 'R', 'E', 'B', 'L', 'A'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'A', 'N', 'C', 'Y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := InvertCapitalize(tt.in)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// t
func TestToggleCase(t *testing.T) {
	//t
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'O', 'W', 'N'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'S', 'P', '3', '3', 'D'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'R', 'O', 'E', 'A', 'R', 'O', 'E', 'A'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'Y', 'R', 'A', 'L', 'L', 'I', 'H'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'G', 'E', 'M', 'M', 'A', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'L', 'O', 'V', 'E', '4', 'E', 'V', 'E', 'R', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'T', 'E', 'L', 'E', 'P', 'H', 'O', 'N', 'E', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'V', 'N', 'V', 'N', 'V', 'N'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'F', 'D', 'S', '2', 'F', 'D'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'A', 'M', 'U', 'R'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'S', 'O', 'C', 'C', 'E', 'R', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'P', 'A', 'R', 'O', 'A', 'C', 'H'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'d', 'O', 'R', 'I', 'T', 'O', 'S'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'S', 'K', 'I', 'P', 'E', 'R'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'T', 'H', 'X', '1', 'M', 'A', 'N'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'C', 'H', 'U', 'L', 'A', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 'T', 'R', 'E', 'B', 'L', 'A'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'N', 'A', 'N', 'C', 'Y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := ToggleCase(tt.in)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// TN
func TestToggleAt(t *testing.T) {
	//T3
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', '3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', 'A', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', 'L', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'm', 'M', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', 'E', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', 'E', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'N', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 's', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'U', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'c', 'C', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', 'O', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 'I', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 'i', 'P', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'u', 'L', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'E', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'C', 'y', '2', '1'}},

		// made by hand
		{[]rune{'T', 'H', 'R', 'E', 'A', 'S', 'D'}, []rune{'T', 'H', 'R', 'e', 'A', 'S', 'D'}},
	}

	for _, tt := range ruletest {
		actual := ToggleAt(tt.in, 3)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// r
func TestReverse(t *testing.T) {
	//r
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'n', 'w', 'o'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'d', '3', '3', 'p', 's'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'a', 'e', 'o', 'r', 'a', 'e', 'o', 'r'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'2', '6', '6', '6', '5', '5', '5'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'h', 'i', 'l', 'l', 'a', 'r', 'y'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'8', 'a', 'm', 'm', 'e', 'g'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'1', '2', 'r', 'e', 'v', 'e', '4', 'e', 'v', 'o', 'l'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'2', '1', 'e', 'n', 'o', 'h', 'p', 'e', 'l', 'e', 't'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'n', 'v', 'n', 'v', 'n', 'v'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'d', 'f', '2', 's', 'd', 'f'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'r', 'u', 'm', 'a', '`'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'3', '2', '7', '0', 'r', 'e', 'c', 'c', 'o', 's'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'h', 'c', 'a', 'o', 'r', 'a', 'p'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'s', 'o', 't', 'i', 'r', 'o', 'D'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'r', 'e', 'p', 'i', 'k', 's'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'n', 'a', 'm', '1', 'x', 'h', 't'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'4', 'a', 'l', 'u', 'h', 'c'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'a', 'l', 'b', 'e', 'r', 't', '1'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'1', '2', 'y', 'c', 'n', 'a', 'n'}},
	}

	for _, tt := range ruletest {
		actual := Reverse(tt.in)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// d
func TestDuplicate(t *testing.T) {
	//d
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', 'n', 'o', 'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', '3', '3', 'd', 's', 'p', '3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2', '5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', 'l', 'l', 'i', 'h', 'y', 'r', 'a', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'm', 'm', 'a', '8', 'g', 'e', 'm', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1', 'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2', 't', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5', '5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'n', 'v', 'n', 'v', 'n', 'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 's', '2', 'f', 'd', 'f', 'd', 's', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'u', 'r', '`', 'a', 'm', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3', 's', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', 'o', 'a', 'c', 'h', 'p', 'a', 'r', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 'i', 't', 'o', 's', 'D', 'o', 'r', 'i', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 'i', 'p', 'e', 'r', 's', 'k', 'i', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', '1', 'm', 'a', 'n', 't', 'h', 'x', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'u', 'l', 'a', '4', 'c', 'h', 'u', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'e', 'b', 'l', 'a', '1', 't', 'r', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'c', 'y', '2', '1', 'n', 'a', 'n', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := Duplicate(tt.in)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// pN
func TestDuplicateN(t *testing.T) {
	//p2
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', 'n', 'o', 'w', 'n', 'o', 'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', '3', '3', 'd', 's', 'p', '3', '3', 'd', 's', 'p', '3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2', '5', '5', '5', '6', '6', '6', '2', '5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', 'l', 'l', 'i', 'h', 'y', 'r', 'a', 'l', 'l', 'i', 'h', 'y', 'r', 'a', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'm', 'm', 'a', '8', 'g', 'e', 'm', 'm', 'a', '8', 'g', 'e', 'm', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1', 'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1', 'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2', 't', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2', 't', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5', '5', '1', '5', '1', '5', '5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'n', 'v', 'n', 'v', 'n', 'v', 'n', 'v', 'n', 'v', 'n', 'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 's', '2', 'f', 'd', 'f', 'd', 's', '2', 'f', 'd', 'f', 'd', 's', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'u', 'r', '`', 'a', 'm', 'u', 'r', '`', 'a', 'm', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3', 's', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3', 's', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', 'o', 'a', 'c', 'h', 'p', 'a', 'r', 'o', 'a', 'c', 'h', 'p', 'a', 'r', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 'i', 't', 'o', 's', 'D', 'o', 'r', 'i', 't', 'o', 's', 'D', 'o', 'r', 'i', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 'i', 'p', 'e', 'r', 's', 'k', 'i', 'p', 'e', 'r', 's', 'k', 'i', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', '1', 'm', 'a', 'n', 't', 'h', 'x', '1', 'm', 'a', 'n', 't', 'h', 'x', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'u', 'l', 'a', '4', 'c', 'h', 'u', 'l', 'a', '4', 'c', 'h', 'u', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'e', 'b', 'l', 'a', '1', 't', 'r', 'e', 'b', 'l', 'a', '1', 't', 'r', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'c', 'y', '2', '1', 'n', 'a', 'n', 'c', 'y', '2', '1', 'n', 'a', 'n', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := DuplicateN(tt.in, 2)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// f
func TestReflect(t *testing.T) {
	//f
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', 'n', 'n', 'w', 'o'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', '3', '3', 'd', 'd', '3', '3', 'p', 's'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a', 'a', 'e', 'o', 'r', 'a', 'e', 'o', 'r'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2', '2', '6', '6', '6', '5', '5', '5'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', 'l', 'l', 'i', 'h', 'h', 'i', 'l', 'l', 'a', 'r', 'y'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'm', 'm', 'a', '8', '8', 'a', 'm', 'm', 'e', 'g'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1', '1', '2', 'r', 'e', 'v', 'e', '4', 'e', 'v', 'o', 'l'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2', '2', '1', 'e', 'n', 'o', 'h', 'p', 'e', 'l', 'e', 't'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5', '5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'n', 'v', 'n', 'n', 'v', 'n', 'v', 'n', 'v'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 's', '2', 'f', 'd', 'd', 'f', '2', 's', 'd', 'f'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'u', 'r', 'r', 'u', 'm', 'a', '`'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3', '3', '2', '7', '0', 'r', 'e', 'c', 'c', 'o', 's'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', 'o', 'a', 'c', 'h', 'h', 'c', 'a', 'o', 'r', 'a', 'p'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 'i', 't', 'o', 's', 's', 'o', 't', 'i', 'r', 'o', 'D'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 'i', 'p', 'e', 'r', 'r', 'e', 'p', 'i', 'k', 's'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', '1', 'm', 'a', 'n', 'n', 'a', 'm', '1', 'x', 'h', 't'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'u', 'l', 'a', '4', '4', 'a', 'l', 'u', 'h', 'c'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'e', 'b', 'l', 'a', 'a', 'l', 'b', 'e', 'r', 't', '1'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'c', 'y', '2', '1', '1', '2', 'y', 'c', 'n', 'a', 'n'}},
	}

	for _, tt := range ruletest {
		actual := Reflect(tt.in)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// {
func TestRotateLeft(t *testing.T) {
	//{
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'w', 'n', 'o'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'p', '3', '3', 'd', 's'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'o', 'e', 'a', 'r', 'o', 'e', 'a', 'r'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '6', '6', '6', '2', '5'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'r', 'a', 'l', 'l', 'i', 'h', 'y'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'e', 'm', 'm', 'a', '8', 'g'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1', 'l'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2', 't'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'1', '5', '1', '5', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'n', 'v', 'n', 'v', 'n', 'v'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'d', 's', '2', 'f', 'd', 'f'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'a', 'm', 'u', 'r', '`'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3', 's'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'a', 'r', 'o', 'a', 'c', 'h', 'p'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'o', 'r', 'i', 't', 'o', 's', 'D'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'k', 'i', 'p', 'e', 'r', 's'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'h', 'x', '1', 'm', 'a', 'n', 't'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'h', 'u', 'l', 'a', '4', 'c'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'t', 'r', 'e', 'b', 'l', 'a', '1'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'a', 'n', 'c', 'y', '2', '1', 'n'}},
	}

	for _, tt := range ruletest {
		actual := RotateLeft(tt.in)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// }
func TestRotateRight(t *testing.T) {
	//}
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'n', 'o', 'w'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'d', 's', 'p', '3', '3'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'a', 'r', 'o', 'e', 'a', 'r', 'o', 'e'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'2', '5', '5', '5', '6', '6', '6'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'h', 'y', 'r', 'a', 'l', 'l', 'i'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'8', 'g', 'e', 'm', 'm', 'a'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'1', 'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'2', 't', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '5', '1', '5', '1'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'n', 'v', 'n', 'v', 'n', 'v'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'d', 'f', 'd', 's', '2', 'f'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'r', '`', 'a', 'm', 'u'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'3', 's', 'o', 'c', 'c', 'e', 'r', '0', '7', '2'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'h', 'p', 'a', 'r', 'o', 'a', 'c'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'s', 'D', 'o', 'r', 'i', 't', 'o'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'r', 's', 'k', 'i', 'p', 'e'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'n', 't', 'h', 'x', '1', 'm', 'a'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'4', 'c', 'h', 'u', 'l', 'a'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'a', '1', 't', 'r', 'e', 'b', 'l'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'1', 'n', 'a', 'n', 'c', 'y', '2'}},
	}

	for _, tt := range ruletest {
		actual := RotateRight(tt.in)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// $X
func TestAppendCharacter(t *testing.T) {
	//$1
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', 'n', '1'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', '3', '3', 'd', '1'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a', '1'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2', '1'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', 'l', 'l', 'i', 'h', '1'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'm', 'm', 'a', '8', '1'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2', '1'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5', '1'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'n', 'v', 'n', '1'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 's', '2', 'f', 'd', '1'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'u', 'r', '1'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3', '1'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', 'o', 'a', 'c', 'h', '1'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 'i', 't', 'o', 's', '1'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 'i', 'p', 'e', 'r', '1'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', '1', 'm', 'a', 'n', '1'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'u', 'l', 'a', '4', '1'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'e', 'b', 'l', 'a', '1'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'c', 'y', '2', '1', '1'}},
	}

	for _, tt := range ruletest {
		actual := AppendCharacter(tt.in, '1')
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// ^X
func TestPrependCharacter(t *testing.T) {
	//^1
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'1', 'o', 'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'1', 's', 'p', '3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'1', 'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'1', '5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'1', 'y', 'r', 'a', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'1', 'g', 'e', 'm', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'1', 'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'1', 't', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'1', '5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'1', 'v', 'n', 'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'1', 'f', 'd', 's', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'1', '`', 'a', 'm', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'1', 's', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'1', 'p', 'a', 'r', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'1', 'D', 'o', 'r', 'i', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'1', 's', 'k', 'i', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'1', 't', 'h', 'x', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'1', 'c', 'h', 'u', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', '1', 't', 'r', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'1', 'n', 'a', 'n', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := PrependCharacter(tt.in, '1')
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// [
func TestTruncateLeft(t *testing.T) {
	//[
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'p', '3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'o', 'e', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'r', 'a', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'e', 'm', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'n', 'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'d', 's', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'a', 'm', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'a', 'r', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'o', 'r', 'i', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'k', 'i', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'h', 'x', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'h', 'u', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'t', 'r', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'a', 'n', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := TruncateLeft(tt.in)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// ]
func TestTruncateRight(t *testing.T) {
	//]
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', '3', '3'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', 'a', 'r', 'o', 'e'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', 'l', 'l', 'i'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'm', 'm', 'a'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'n', 'v'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 's', '2', 'f'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'u'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', 'o', 'a', 'c'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 'i', 't', 'o'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 'i', 'p', 'e'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', '1', 'm', 'a'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'u', 'l', 'a'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'e', 'b', 'l'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'c', 'y', '2'}},
	}

	for _, tt := range ruletest {
		actual := TruncateRight(tt.in)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// DN
func TestDeleteN(t *testing.T) {
	//D3
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 's', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 'i', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'u', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := DeleteN(tt.in, 3)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// xNM
func TestExtractRange(t *testing.T) {
	//x02
	var ruletest = []struct {
		in  []rune
		out []rune
	}{

		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a'}},
	}

	for _, tt := range ruletest {
		actual := ExtractRange(tt.in, 0, 2)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}
}

// ONM
func TestOmitRange(t *testing.T) {
	//O02
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'e', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'a', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'m', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'s', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'m', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'c', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'r', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'r', 'i', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'i', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'x', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'u', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'r', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := OmitRange(tt.in, 0, 2)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// iNX
func TestInsertAtN(t *testing.T) {
	//i4!
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', '3', '3', '!', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', 'a', '!', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '!', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', 'l', '!', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'm', 'm', '!', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', 'e', '!', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', 'e', '!', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '!', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'n', '!', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 's', '2', '!', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'u', '!', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'c', 'c', '!', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', 'o', '!', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 'i', '!', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 'i', 'p', '!', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', '1', '!', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'u', 'l', '!', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'e', '!', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'c', '!', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := InsertAtN(tt.in, 4, '!')
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// oNX
func TestOverwriteAtN(t *testing.T) {
	//o3$
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', '3', '$', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', '$', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '$', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', '$', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'm', '$', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', '$', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', '$', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '$', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', '$', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 's', '$', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', '$', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'c', '$', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', '$', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', '$', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 'i', '$', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', '$', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'u', '$', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', '$', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', '$', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := OverwriteAtN(tt.in, 3, '$')
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// 'N
func TestTruncateAtN(t *testing.T) {
	//'6
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', '3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', 'a', 'r', 'o'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', 'l', 'l', 'i'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'm', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', 'e', '4', 'e'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', 'e', 'p', 'h'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 's', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'c', 'c', 'e', 'r'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', 'o', 'a', 'c'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 'i', 't', 'o'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 'i', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', '1', 'm', 'a'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'u', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'e', 'b', 'l'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'c', 'y', '2'}},
	}

	for _, tt := range ruletest {
		actual := TruncateAtN(tt.in, 6)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// sXY
func TestReplace(t *testing.T) {
	//ss$
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'$', 'p', '3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'm', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', '$', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'$', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 'i', 't', 'o', '$'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'$', 'k', 'i', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'u', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := Replace(tt.in, 's', '$')
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// @X
func TestPurge(t *testing.T) {
	//@s
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'p', '3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'm', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 'i', 't', 'o'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'k', 'i', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'u', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := Purge(tt.in, 's')
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// zN
func TestDuplicateFirstN(t *testing.T) {
	//z2
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'o', 'o', 'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 's', 's', 'p', '3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'r', 'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'y', 'y', 'r', 'a', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'g', 'g', 'e', 'm', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'l', 'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 't', 't', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '5', '5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'v', 'v', 'n', 'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'f', 'f', 'd', 's', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', '`', '`', 'a', 'm', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 's', 's', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'p', 'p', 'a', 'r', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'D', 'D', 'o', 'r', 'i', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 's', 's', 'k', 'i', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 't', 't', 'h', 'x', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'c', 'c', 'h', 'u', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', '1', '1', 't', 'r', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'n', 'n', 'a', 'n', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := DuplicateFirstN(tt.in, 2)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// ZN
func TestDuplicateLastN(t *testing.T) {
	//Z2
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', 'n', 'n', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', '3', '3', 'd', 'd', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a', 'a', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2', '2', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', 'l', 'l', 'i', 'h', 'h', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'm', 'm', 'a', '8', '8', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1', '1', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2', '2', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5', '5', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'n', 'v', 'n', 'n', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 's', '2', 'f', 'd', 'd', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'u', 'r', 'r', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3', '3', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', 'o', 'a', 'c', 'h', 'h', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 'i', 't', 'o', 's', 's', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 'i', 'p', 'e', 'r', 'r', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', '1', 'm', 'a', 'n', 'n', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'u', 'l', 'a', '4', '4', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'e', 'b', 'l', 'a', 'a', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'c', 'y', '2', '1', '1', '1'}},
	}

	for _, tt := range ruletest {
		actual := DuplicateLastN(tt.in, 2)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// q
func TestDuplicateAll(t *testing.T) {
	//q
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'o', 'w', 'w', 'n', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 's', 'p', 'p', '3', '3', '3', '3', 'd', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'r', 'o', 'o', 'e', 'e', 'a', 'a', 'r', 'r', 'o', 'o', 'e', 'e', 'a', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '5', '5', '5', '6', '6', '6', '6', '6', '6', '2', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'y', 'r', 'r', 'a', 'a', 'l', 'l', 'l', 'l', 'i', 'i', 'h', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'g', 'e', 'e', 'm', 'm', 'm', 'm', 'a', 'a', '8', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'l', 'o', 'o', 'v', 'v', 'e', 'e', '4', '4', 'e', 'e', 'v', 'v', 'e', 'e', 'r', 'r', '2', '2', '1', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 't', 'e', 'e', 'l', 'l', 'e', 'e', 'p', 'p', 'h', 'h', 'o', 'o', 'n', 'n', 'e', 'e', '1', '1', '2', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '5', '1', '1', '5', '5', '1', '1', '5', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'v', 'n', 'n', 'v', 'v', 'n', 'n', 'v', 'v', 'n', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'f', 'd', 'd', 's', 's', '2', '2', 'f', 'f', 'd', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', '`', 'a', 'a', 'm', 'm', 'u', 'u', 'r', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 's', 'o', 'o', 'c', 'c', 'c', 'c', 'e', 'e', 'r', 'r', '0', '0', '7', '7', '2', '2', '3', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'p', 'a', 'a', 'r', 'r', 'o', 'o', 'a', 'a', 'c', 'c', 'h', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'D', 'o', 'o', 'r', 'r', 'i', 'i', 't', 't', 'o', 'o', 's', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 's', 'k', 'k', 'i', 'i', 'p', 'p', 'e', 'e', 'r', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 't', 'h', 'h', 'x', 'x', '1', '1', 'm', 'm', 'a', 'a', 'n', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'c', 'h', 'h', 'u', 'u', 'l', 'l', 'a', 'a', '4', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', '1', 't', 't', 'r', 'r', 'e', 'e', 'b', 'b', 'l', 'l', 'a', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'n', 'a', 'a', 'n', 'n', 'c', 'c', 'y', 'y', '2', '2', '1', '1'}},
	}

	for _, tt := range ruletest {
		actual := DuplicateAll(tt.in)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// XNMI
func TestExtractMemory(t *testing.T) {
	//X427
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'r', 'o', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2', '6', '6'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', 'l', 'l', 'i', 'h', 'l', 'i'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', 'e', '4', 'e', 'v', '4', 'e', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'p', 'h', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'c', 'c', 'e', 'r', '0', 'e', 'r', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', 'o', 'a', 'c', 'h', 'a', 'c'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 'i', 't', 'o', 's', 't', 'o'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', '1', 'm', 'a', 'n', 'm', 'a'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'e', 'b', 'l', 'a', 'b', 'l'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'c', 'y', '2', '1', 'y', '2'}},
	}

	for _, tt := range ruletest {
		Memorize(tt.in)
		actual := ExtractMemory(tt.in, 4, 2, 7)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}
}

// 4
func TestAppendMemory(t *testing.T) {
	//M4
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', 'n', 'o', 'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', '3', '3', 'd', 's', 'p', '3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2', '5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', 'l', 'l', 'i', 'h', 'y', 'r', 'a', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'm', 'm', 'a', '8', 'g', 'e', 'm', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1', 'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2', 't', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5', '5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'n', 'v', 'n', 'v', 'n', 'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 's', '2', 'f', 'd', 'f', 'd', 's', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'u', 'r', '`', 'a', 'm', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3', 's', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', 'o', 'a', 'c', 'h', 'p', 'a', 'r', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 'i', 't', 'o', 's', 'D', 'o', 'r', 'i', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 'i', 'p', 'e', 'r', 's', 'k', 'i', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', '1', 'm', 'a', 'n', 't', 'h', 'x', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'u', 'l', 'a', '4', 'c', 'h', 'u', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'e', 'b', 'l', 'a', '1', 't', 'r', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'c', 'y', '2', '1', 'n', 'a', 'n', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		Memorize(tt.in)
		actual := AppendMemory(tt.in)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// 6
func TestPrependMemory(t *testing.T) {
	//M6
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', 'n', 'o', 'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', '3', '3', 'd', 's', 'p', '3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2', '5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', 'l', 'l', 'i', 'h', 'y', 'r', 'a', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'm', 'm', 'a', '8', 'g', 'e', 'm', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1', 'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2', 't', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5', '5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'n', 'v', 'n', 'v', 'n', 'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 's', '2', 'f', 'd', 'f', 'd', 's', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'u', 'r', '`', 'a', 'm', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3', 's', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', 'o', 'a', 'c', 'h', 'p', 'a', 'r', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 'i', 't', 'o', 's', 'D', 'o', 'r', 'i', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 'i', 'p', 'e', 'r', 's', 'k', 'i', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', '1', 'm', 'a', 'n', 't', 'h', 'x', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'u', 'l', 'a', '4', 'c', 'h', 'u', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'e', 'b', 'l', 'a', '1', 't', 'r', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'c', 'y', '2', '1', 'n', 'a', 'n', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		Memorize(tt.in)
		actual := PrependMemory(tt.in)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// M
func TestMemorize(t *testing.T) {
	//M
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', '3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'm', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 's', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 'i', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 'i', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'u', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		Memorize(tt.in)
		// dont know if that will work there
		if compare(Saved, tt.out) {
			t.Errorf("expected %s, got %s", tt.out, Saved)
		}
	}

}

// <M
func TestRejectLess(t *testing.T) {
	//<7
	var ruletest = []struct {
		in  []rune
		out bool
	}{
		{[]rune{'o', 'w', 'n'}, false},
		{[]rune{'s', 'p', '3', '3', 'd'}, false},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, true},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, false},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, false},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, false},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, true},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, true},
		{[]rune{'5', '1', '5', '1', '5'}, false},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, false},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, false},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, false},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, true},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, false},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, false},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, false},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, false},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, false},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, false},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, false},
	}

	for _, tt := range ruletest {
		actual := RejectLess(tt.in, 7)
		if actual != tt.out {
			t.Errorf("expected %t, got %t, %s", tt.out, actual, tt.in)
		}
	}
}

// >N
func TestRejectGreater(t *testing.T) {
	//>7
	var ruletest = []struct {
		in  []rune
		out bool
	}{
		{[]rune{'o', 'w', 'n'}, true},
		{[]rune{'s', 'p', '3', '3', 'd'}, true},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, false},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, false},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, false},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, true},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, false},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, false},
		{[]rune{'5', '1', '5', '1', '5'}, true},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, true},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, true},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, true},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, false},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, false},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, false},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, true},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, false},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, true},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, false},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, false},
	}

	for _, tt := range ruletest {
		actual := RejectGreater(tt.in, 7)
		if actual != tt.out {
			t.Errorf("expected %t, got %t, %s", tt.out, actual, tt.in)
		}
	}

}

// !X
func TestRejectContain(t *testing.T) {
	//!e
	var ruletest = []struct {
		in  []rune
		out bool
	}{
		{[]rune{'o', 'w', 'n'}, false},
		{[]rune{'s', 'p', '3', '3', 'd'}, false},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, true},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, false},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, false},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, true},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, true},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, true},
		{[]rune{'5', '1', '5', '1', '5'}, false},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, false},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, false},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, false},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, true},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, false},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, false},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, true},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, false},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, false},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, true},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, false},
	}

	for _, tt := range ruletest {
		actual := RejectContain(tt.in, 'e')
		if actual != tt.out {
			t.Errorf("expected %t, got %t", tt.out, actual)
		}
	}

}

// /X
func TestRejectNotContain(t *testing.T) {
	///e
	var ruletest = []struct {
		in  []rune
		out bool
	}{
		{[]rune{'o', 'w', 'n'}, true},
		{[]rune{'s', 'p', '3', '3', 'd'}, true},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, false},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, true},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, true},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, false},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, false},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, false},
		{[]rune{'5', '1', '5', '1', '5'}, true},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, true},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, true},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, true},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, false},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, true},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, true},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, false},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, true},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, true},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, false},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, true},
	}

	for _, tt := range ruletest {
		actual := RejectNotContain(tt.in, 'e')
		if actual != tt.out {
			t.Errorf("expected %t, got %t", tt.out, actual)
		}
	}

}

// (X
func TestRejectEqualFirst(t *testing.T) {
	//(s
	var ruletest = []struct {
		in  []rune
		out bool
	}{
		{[]rune{'o', 'w', 'n'}, true},
		{[]rune{'s', 'p', '3', '3', 'd'}, false},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, true},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, true},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, true},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, true},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, true},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, true},
		{[]rune{'5', '1', '5', '1', '5'}, true},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, true},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, true},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, true},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, false},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, true},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, true},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, false},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, true},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, true},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, true},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, true},
	}

	for _, tt := range ruletest {
		actual := RejectEqualFirst(tt.in, 's')
		if actual != tt.out {
			t.Errorf("expected %t, got %t", tt.out, actual)
		}
	}

}

// )X
func TestRejectEqualLast(t *testing.T) {
	//)e
	var ruletest = []struct {
		in  []rune
		out bool
	}{
		{[]rune{'o', 'w', 'n'}, false},
		{[]rune{'s', 'p', '3', '3', 'd'}, true},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, true},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, true},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, true},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, true},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, true},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, true},
		{[]rune{'5', '1', '5', '1', '5'}, true},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, false},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, true},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, true},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, true},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, true},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, true},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, true},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, false},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, true},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, true},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, true},
	}

	for _, tt := range ruletest {
		actual := RejectEqualLast(tt.in, 'n')
		if actual != tt.out {
			t.Errorf("expected %t, got %t, %s", tt.out, actual, tt.in)
		}
	}
}

// =NX
func TestRejectEqualAt(t *testing.T) {
	//=1o
	var ruletest = []struct {
		in  []rune
		out bool
	}{
		{[]rune{'o', 'w', 'n'}, true},
		{[]rune{'s', 'p', '3', '3', 'd'}, true},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, false},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, true},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, true},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, true},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, false},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, true},
		{[]rune{'5', '1', '5', '1', '5'}, true},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, true},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, true},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, true},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, false},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, true},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, false},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, true},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, true},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, true},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, true},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, true},
	}

	for _, tt := range ruletest {
		actual := RejectEqualAt(tt.in, 'o', 1)
		if actual != tt.out {
			t.Errorf("expected %t, got %t", tt.out, actual)
		}
	}

}

// %NX
func TestRejectContains(t *testing.T) {
	//%1a
	var ruletest = []struct {
		in  []rune
		out bool
	}{
		{[]rune{'o', 'w', 'n'}, true},
		{[]rune{'s', 'p', '3', '3', 'd'}, true},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, false},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, true},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, false},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, false},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, true},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, true},
		{[]rune{'5', '1', '5', '1', '5'}, true},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, true},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, true},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, false},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, true},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, false},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, true},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, true},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, false},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, false},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, false},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, false},
	}

	for _, tt := range ruletest {
		actual := RejectContains(tt.in, 'a', 1)
		if actual != tt.out {
			t.Errorf("expected %t, got %t", tt.out, actual)
		}
	}

}

// Q
func TestRejectMemory(t *testing.T) {
	//rMrQ
	var ruletest = []struct {
		in  []rune
		out bool
	}{
		{[]rune{'o', 'w', 'n'}, false},
		{[]rune{'s', 'p', '3', '3', 'd'}, false},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, false},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, false},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, false},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, true},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, false},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, false},
		{[]rune{'5', '1', '5', '1', '5'}, false},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, false},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, false},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, false},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, false},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, false},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, false},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, false},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, false},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, false},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, false},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, false},
	}

	Memorize([]rune{'g', 'e', 'm', 'm', 'a', '8'})

	for _, tt := range ruletest {
		actual := RejectMemory(tt.in)
		if actual != tt.out {
			t.Errorf("expected %t, got %t", tt.out, actual)
		}
	}

}

// k
func TestSwapFront(t *testing.T) {
	//k
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'w', 'o', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'p', 's', '3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'o', 'r', 'e', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'r', 'y', 'a', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'e', 'g', 'm', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'o', 'l', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'e', 't', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'1', '5', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'n', 'v', 'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'d', 'f', 's', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'a', '`', 'm', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'o', 's', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'a', 'p', 'r', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'o', 'D', 'r', 'i', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'k', 's', 'i', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'h', 't', 'x', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'h', 'c', 'u', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'t', '1', 'r', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'a', 'n', 'n', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := SwapFront(tt.in)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// K
func TestSwapBack(t *testing.T) {
	//K
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'n', 'w'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', '3', 'd', '3'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', 'a', 'r', 'o', 'a', 'e'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '2', '6'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', 'l', 'l', 'h', 'i'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'm', 'm', '8', 'a'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '1', '2'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '2', '1'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '5', '1'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'n', 'n', 'v'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 's', '2', 'd', 'f'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'r', 'u'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '3', '2'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', 'o', 'a', 'h', 'c'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 'i', 't', 's', 'o'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 'i', 'p', 'r', 'e'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', '1', 'm', 'n', 'a'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'u', 'l', '4', 'a'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'e', 'b', 'a', 'l'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'c', 'y', '1', '2'}},
	}

	for _, tt := range ruletest {
		actual := SwapBack(tt.in)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// *XY
func TestSwapAtN(t *testing.T) {
	//*34
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', '3', 'd', '3'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', 'r', 'a', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'm', 'a', 'm', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', '4', 'e', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', 'p', 'e', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '5', '1'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'v', 'n', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 's', 'f', '2', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'r', 'u'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'c', 'e', 'c', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', 'a', 'o', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 't', 'i', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 'i', 'e', 'p', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', 'm', '1', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'u', 'a', 'l', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'b', 'e', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'y', 'c', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := SwapAtN(tt.in, 3, 4)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// LN
func TestBitwiseShiftLeft(t *testing.T) {
	//L2
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', 'f', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'Ê', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', 'j', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'Â', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'Ú', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'ì', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'Ø', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', 'j', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'ì', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 'æ', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'Ú', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'Æ', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'ä', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'ä', 'i', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 'Ò', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'ð', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'ê', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'ä', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'Ü', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := BitwiseShiftLeft(tt.in, 2)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// RN
func TestBitwiseShiftRight(t *testing.T) {
	//R2
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', '7'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', '', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', '2', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', '0', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', '6', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', ';', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', '6', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', ';', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', '9', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', '6', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', '1', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', '9', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', '9', 'i', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', '4', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', '<', '1', 'm', 'a', 'n'}},
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', '7'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', ':', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', '9', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', '7', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := BitwiseShiftRight(tt.in, 2)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// +N
func TestASCIIIncrementPlus(t *testing.T) {
	//+2
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', 'o'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', '4', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'f', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '6', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'b', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'n', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'w', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'm', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '6', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'w', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 't', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'n', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'd', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 's', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 's', 'i', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 'j', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'y', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'v', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 's', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'o', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := ASCIIIncrementPlus(tt.in, 2)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// -N
func TestASCIIIncrementMinus(t *testing.T) {
	//-1
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'v', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'o', '3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'n', 'e', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '4', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'q', 'a', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'd', 'm', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'n', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'd', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '0', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'm', 'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'c', 's', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', '`', 'm', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'n', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', '`', 'r', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'n', 'r', 'i', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'j', 'i', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'g', 'x', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'g', 'u', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 's', 'r', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', '`', 'n', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := ASCIIIncrementMinus(tt.in, 1)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// .N
func TestReplaceNPlus(t *testing.T) {
	//.1
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'n', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', '3', '3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'e', 'e', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'a', 'a', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'm', 'm', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'v', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'l', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '5', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'v', 'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 's', 's', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'm', 'm', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'c', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'r', 'r', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'r', 'r', 'i', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'i', 'i', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'x', 'x', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'u', 'u', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 'r', 'r', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'n', 'n', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := ReplaceNPlus(tt.in, 1)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// ,N
func TestReplaceNMinus(t *testing.T) {
	//,1
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'o', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 's', '3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'r', 'e', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'y', 'a', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'g', 'm', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'l', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 't', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '5', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'v', 'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'f', 's', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', '`', 'm', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 's', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'p', 'r', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'D', 'r', 'i', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 's', 'i', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 't', 'x', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'c', 'u', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', '1', 'r', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'n', 'n', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := ReplaceNMinus(tt.in, 1)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// yN
func TestDuplicateBlockFront(t *testing.T) {
	//y2
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', 'o', 'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', 's', 'p', '3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'y', 'r', 'a', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'g', 'e', 'm', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 't', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'n', 'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 'f', 'd', 's', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', '`', 'a', 'm', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 's', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'p', 'a', 'r', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'D', 'o', 'r', 'i', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 's', 'k', 'i', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 't', 'h', 'x', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'c', 'h', 'u', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', '1', 't', 'r', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'a', 'n', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := DuplicateBlockFront(tt.in, 2)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// YN
func TestDuplicateBlockBack(t *testing.T) {
	//Y2
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'o', 'w', 'n', 'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'s', 'p', '3', '3', 'd', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'y', 'r', 'a', 'l', 'l', 'i', 'h', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'g', 'e', 'm', 'm', 'a', '8', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'v', 'n', 'v', 'n', 'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'f', 'd', 's', '2', 'f', 'd', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'u', 'r', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'p', 'a', 'r', 'o', 'a', 'c', 'h', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 'i', 't', 'o', 's', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'s', 'k', 'i', 'p', 'e', 'r', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'t', 'h', 'x', '1', 'm', 'a', 'n', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'c', 'h', 'u', 'l', 'a', '4', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'e', 'b', 'l', 'a', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'n', 'a', 'n', 'c', 'y', '2', '1', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := DuplicateBlockBack(tt.in, 2)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

// fix
// E
func TestTitle(t *testing.T) {
	//E
	var ruletest = []struct {
		in  []rune
		out []rune
	}{
		{[]rune{'o', 'w', 'n'}, []rune{'O', 'w', 'n'}},
		{[]rune{'s', 'p', '3', '3', 'd'}, []rune{'S', 'p', '3', '3', 'd'}},
		{[]rune{'r', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}, []rune{'R', 'o', 'e', 'a', 'r', 'o', 'e', 'a'}},
		{[]rune{'5', '5', '5', '6', '6', '6', '2'}, []rune{'5', '5', '5', '6', '6', '6', '2'}},
		{[]rune{'y', 'r', 'a', 'l', 'l', 'i', 'h'}, []rune{'Y', 'r', 'a', 'l', 'l', 'i', 'h'}},
		{[]rune{'g', 'e', 'm', 'm', 'a', '8'}, []rune{'G', 'e', 'm', 'm', 'a', '8'}},
		{[]rune{'l', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}, []rune{'L', 'o', 'v', 'e', '4', 'e', 'v', 'e', 'r', '2', '1'}},
		{[]rune{'t', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}, []rune{'T', 'e', 'l', 'e', 'p', 'h', 'o', 'n', 'e', '1', '2'}},
		{[]rune{'5', '1', '5', '1', '5'}, []rune{'5', '1', '5', '1', '5'}},
		{[]rune{'v', 'n', 'v', 'n', 'v', 'n'}, []rune{'V', 'n', 'v', 'n', 'v', 'n'}},
		{[]rune{'f', 'd', 's', '2', 'f', 'd'}, []rune{'F', 'd', 's', '2', 'f', 'd'}},
		{[]rune{'`', 'a', 'm', 'u', 'r'}, []rune{'`', 'a', 'm', 'u', 'r'}},
		{[]rune{'s', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}, []rune{'S', 'o', 'c', 'c', 'e', 'r', '0', '7', '2', '3'}},
		{[]rune{'p', 'a', 'r', 'o', 'a', 'c', 'h'}, []rune{'P', 'a', 'r', 'o', 'a', 'c', 'h'}},
		{[]rune{'D', 'o', 'r', 'i', 't', 'o', 's'}, []rune{'D', 'o', 'r', 'i', 't', 'o', 's'}},
		{[]rune{'s', 'k', 'i', 'p', 'e', 'r'}, []rune{'S', 'k', 'i', 'p', 'e', 'r'}},
		{[]rune{'t', 'h', 'x', '1', 'm', 'a', 'n'}, []rune{'T', 'h', 'x', '1', 'm', 'a', 'n'}},
		{[]rune{'c', 'h', 'u', 'l', 'a', '4'}, []rune{'C', 'h', 'u', 'l', 'a', '4'}},
		{[]rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}, []rune{'1', 't', 'r', 'e', 'b', 'l', 'a'}},
		{[]rune{'n', 'a', 'n', 'c', 'y', '2', '1'}, []rune{'N', 'a', 'n', 'c', 'y', '2', '1'}},
	}

	for _, tt := range ruletest {
		actual := Title(tt.in)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

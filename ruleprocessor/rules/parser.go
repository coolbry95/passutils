package rules

//import "fmt"

// hashcat rules
const (
	RULE_OP_MANGLE_NOOP           = ':'
	RULE_OP_MANGLE_LREST          = 'l'
	RULE_OP_MANGLE_UREST          = 'u'
	RULE_OP_MANGLE_LREST_UFIRST   = 'c'
	RULE_OP_MANGLE_UREST_LFIRST   = 'C'
	RULE_OP_MANGLE_TREST          = 't'
	RULE_OP_MANGLE_TOGGLE_AT      = 'T'
	RULE_OP_MANGLE_REVERSE        = 'r'
	RULE_OP_MANGLE_DUPEWORD       = 'd'
	RULE_OP_MANGLE_DUPEWORD_TIMES = 'p'
	RULE_OP_MANGLE_REFLECT        = 'f'
	RULE_OP_MANGLE_ROTATE_LEFT    = '{'
	RULE_OP_MANGLE_ROTATE_RIGHT   = '}'
	RULE_OP_MANGLE_APPEND         = '$'
	RULE_OP_MANGLE_PREPEND        = '^'
	RULE_OP_MANGLE_DELETE_FIRST   = '['
	RULE_OP_MANGLE_DELETE_LAST    = ']'
	RULE_OP_MANGLE_DELETE_AT      = 'D'
	RULE_OP_MANGLE_EXTRACT        = 'x'
	RULE_OP_MANGLE_OMIT           = 'O'
	RULE_OP_MANGLE_INSERT         = 'i'
	RULE_OP_MANGLE_OVERSTRIKE     = 'o'
	RULE_OP_MANGLE_TRUNCATE_AT    = '\''
	RULE_OP_MANGLE_REPLACE        = 's'
	RULE_OP_MANGLE_PURGECHAR      = '@'
	RULE_OP_MANGLE_DUPECHAR_FIRST = 'z'
	RULE_OP_MANGLE_DUPECHAR_LAST  = 'Z'
	RULE_OP_MANGLE_DUPECHAR_ALL   = 'q'
	RULE_OP_MANGLE_EXTRACT_MEMORY = 'X'
	RULE_OP_MANGLE_APPEND_MEMORY  = '4'
	RULE_OP_MANGLE_PREPEND_MEMORY = '6'

	RULE_OP_MEMORIZE_WORD = 'M'

	RULE_OP_REJECT_LESS        = '<'
	RULE_OP_REJECT_GREATER     = '>'
	RULE_OP_REJECT_CONTAIN     = '!'
	RULE_OP_REJECT_NOT_CONTAIN = '/'
	RULE_OP_REJECT_EQUAL_FIRST = '('
	RULE_OP_REJECT_EQUAL_LAST  = ')'
	RULE_OP_REJECT_EQUAL_AT    = '='
	RULE_OP_REJECT_CONTAINS    = '%'
	RULE_OP_REJECT_MEMORY      = 'Q'

	RULE_OP_MANGLE_SWITCH_FIRST    = 'k'
	RULE_OP_MANGLE_SWITCH_LAST     = 'K'
	RULE_OP_MANGLE_SWITCH_AT       = '*'
	RULE_OP_MANGLE_CHR_SHIFTL      = 'L'
	RULE_OP_MANGLE_CHR_SHIFTR      = 'R'
	RULE_OP_MANGLE_CHR_INCR        = '+'
	RULE_OP_MANGLE_CHR_DECR        = '-'
	RULE_OP_MANGLE_REPLACE_NP1     = '.'
	RULE_OP_MANGLE_REPLACE_NM1     = ','
	RULE_OP_MANGLE_DUPEBLOCK_FIRST = 'y'
	RULE_OP_MANGLE_DUPEBLOCK_LAST  = 'Y'
	RULE_OP_MANGLE_TITLE           = 'E'
)



// a rule is the entire line
// if one piece is not right then skip the rest of the line


// ParseLine parses a string for hashcat rules then returns a slice of the valid rule.
// One line is considered one rule.
func ParseLine(line string) (rules []string) {
// needs better rule checking.
// eg length of rule
// eg X008 doesnt make sense
	//fmt.Println(line)
	if len(line) != 0 && rune(line[0]) == '#' {
		return nil
	}
	lineRune := []rune(line)
	for i := 0; i < len(lineRune); i++ {
		switch lineRune[i] {
		case ' ':
			// just skip it
		case RULE_OP_MANGLE_NOOP:
			rules = append(rules, string(RULE_OP_MANGLE_NOOP))
		case RULE_OP_MANGLE_LREST:
			rules = append(rules, string(RULE_OP_MANGLE_LREST))
		case RULE_OP_MANGLE_UREST:
			rules = append(rules, string(RULE_OP_MANGLE_UREST))
		case RULE_OP_MANGLE_LREST_UFIRST:
			rules = append(rules, string(RULE_OP_MANGLE_LREST_UFIRST))
		case RULE_OP_MANGLE_UREST_LFIRST:
			rules = append(rules, string(RULE_OP_MANGLE_UREST_LFIRST))
		case RULE_OP_MANGLE_TREST:
			rules = append(rules, string(RULE_OP_MANGLE_TREST))
		case RULE_OP_MANGLE_TOGGLE_AT:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_MANGLE_REVERSE:
			rules = append(rules, string(RULE_OP_MANGLE_REVERSE))
		case RULE_OP_MANGLE_DUPEWORD:
			rules = append(rules, string(RULE_OP_MANGLE_DUPEWORD))
		case RULE_OP_MANGLE_DUPEWORD_TIMES:
			// if the insert position is not valid
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_MANGLE_REFLECT:
			rules = append(rules, string(RULE_OP_MANGLE_REFLECT))
		case RULE_OP_MANGLE_ROTATE_LEFT:
			rules = append(rules, string(RULE_OP_MANGLE_ROTATE_LEFT))
		case RULE_OP_MANGLE_ROTATE_RIGHT:
			rules = append(rules, string(RULE_OP_MANGLE_ROTATE_RIGHT))
		case RULE_OP_MANGLE_APPEND:
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_MANGLE_PREPEND:
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_MANGLE_DELETE_FIRST:
			rules = append(rules, string(RULE_OP_MANGLE_DELETE_FIRST))
		case RULE_OP_MANGLE_DELETE_LAST:
			rules = append(rules, string(RULE_OP_MANGLE_DELETE_LAST))
		case RULE_OP_MANGLE_DELETE_AT:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_MANGLE_EXTRACT:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			if line[i+2] > 90 || line[i+2] < 48 {
				return nil
			} else if line[i+2] < 65 && line[i+2] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+3])
			i += 2
		case RULE_OP_MANGLE_OMIT:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			if line[i+2] > 90 || line[i+2] < 48 {
				return nil
			} else if line[i+2] < 65 && line[i+2] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+3])
			i += 2
		case RULE_OP_MANGLE_INSERT:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+3])
			i += 2
		case RULE_OP_MANGLE_OVERSTRIKE:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+3])
			i += 2
		case RULE_OP_MANGLE_TRUNCATE_AT:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_MANGLE_REPLACE:
			rules = append(rules, line[i:i+3])
			i += 2
		case RULE_OP_MANGLE_PURGECHAR:
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_MANGLE_DUPECHAR_FIRST:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_MANGLE_DUPECHAR_LAST:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_MANGLE_DUPECHAR_ALL:
			rules = append(rules, string(RULE_OP_MANGLE_DUPECHAR_ALL))
		case RULE_OP_MANGLE_EXTRACT_MEMORY:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			if line[i+2] > 90 || line[i+2] < 48 {
				return nil
			} else if line[i+2] < 65 && line[i+2] > 57 {
				return nil
			}
			if line[i+3] > 90 || line[i+3] < 48 {
				return nil
			} else if line[i+3] < 65 && line[i+3] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+4])
			i += 3
		case RULE_OP_MANGLE_APPEND_MEMORY:
			rules = append(rules, string(RULE_OP_MANGLE_APPEND_MEMORY))
		case RULE_OP_MANGLE_PREPEND_MEMORY:
			rules = append(rules, string(RULE_OP_MANGLE_PREPEND_MEMORY))
		case RULE_OP_MEMORIZE_WORD:
			rules = append(rules, string(RULE_OP_MEMORIZE_WORD))
		case RULE_OP_REJECT_LESS:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_REJECT_GREATER:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_REJECT_CONTAIN:
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_REJECT_NOT_CONTAIN:
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_REJECT_EQUAL_FIRST:
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_REJECT_EQUAL_LAST:
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_REJECT_EQUAL_AT:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+3])
			i += 2
		case RULE_OP_REJECT_CONTAINS:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+3])
			i+=2
		case RULE_OP_REJECT_MEMORY:
			rules = append(rules, string(RULE_OP_REJECT_MEMORY))
		case RULE_OP_MANGLE_SWITCH_FIRST:
			rules = append(rules, string(RULE_OP_MANGLE_SWITCH_FIRST))
		case RULE_OP_MANGLE_SWITCH_LAST:
			rules = append(rules, string(RULE_OP_MANGLE_SWITCH_LAST))
		case RULE_OP_MANGLE_SWITCH_AT:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			if line[i+2] > 90 || line[i+2] < 48 {
				return nil
			} else if line[i+2] < 65 && line[i+2] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+3])
			i += 2
		case RULE_OP_MANGLE_CHR_SHIFTL:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_MANGLE_CHR_SHIFTR:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_MANGLE_CHR_INCR:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_MANGLE_CHR_DECR:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_MANGLE_REPLACE_NP1:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_MANGLE_REPLACE_NM1:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_MANGLE_DUPEBLOCK_FIRST:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_MANGLE_DUPEBLOCK_LAST:
			if line[i+1] > 90 || line[i+1] < 48 {
				return nil
			} else if line[i+1] < 65 && line[i+1] > 57 {
				return nil
			}
			rules = append(rules, line[i:i+2])
			i++
		case RULE_OP_MANGLE_TITLE:
			rules = append(rules, string(RULE_OP_MANGLE_TITLE))
			// if something is not right on the line
		default:
			return nil
		}
	}
	return rules
}


// ApplyRules applies a hashcat rule to a single word.
func ApplyRules(rules []string, word string) string {
// maybe make rules [][]rune so indexing is based on rune and not byte
	if rules == nil || word == "" {
		return ""
	}
	mangled := []rune(word)
	reject := false
	for _, rule := range rules {
		switch rule[0] {
		case RULE_OP_MANGLE_NOOP:
			mangled = Nothing(mangled)
		case RULE_OP_MANGLE_LREST:
			mangled = Lowercase(mangled)
		case RULE_OP_MANGLE_UREST:
			mangled = Uppercase(mangled)
		case RULE_OP_MANGLE_LREST_UFIRST:
			mangled = Capitalize(mangled)
		case RULE_OP_MANGLE_UREST_LFIRST:
			mangled = InvertCapitalize(mangled)
		case RULE_OP_MANGLE_TREST:
			mangled = ToggleCase(mangled)
		case RULE_OP_MANGLE_TOGGLE_AT:
			n := ToNumByte(rule[1])
			// cannot apply to rule because the managled word is not long enough
			// we are going to skip the entire rule
			// which is rules
			if n > len(mangled)-1 {
				break
			}
			mangled = ToggleAt(mangled, n)
		case RULE_OP_MANGLE_REVERSE:
			mangled = Reverse(mangled)
		case RULE_OP_MANGLE_DUPEWORD:
			mangled = Duplicate(mangled)
		case RULE_OP_MANGLE_DUPEWORD_TIMES:
			n := ToNumByte(rule[1])
			mangled = DuplicateN(mangled, n)
		case RULE_OP_MANGLE_REFLECT:
			mangled = Reflect(mangled)
		case RULE_OP_MANGLE_ROTATE_LEFT:
			mangled = RotateLeft(mangled)
		case RULE_OP_MANGLE_ROTATE_RIGHT:
			mangled = RotateRight(mangled)
		case RULE_OP_MANGLE_APPEND:
			mangled = AppendCharacter(mangled, rune(rule[1]))
		case RULE_OP_MANGLE_PREPEND:
			mangled = PrependCharacter(mangled, rune(rule[1]))
		case RULE_OP_MANGLE_DELETE_FIRST:
			mangled = TruncateLeft(mangled)
		case RULE_OP_MANGLE_DELETE_LAST:
			mangled = TruncateRight(mangled)
		case RULE_OP_MANGLE_DELETE_AT:
			n := ToNumByte(rule[1])
			if n > len(mangled)-1 {
				break
			}
			mangled = DeleteN(mangled, n)
		case RULE_OP_MANGLE_EXTRACT:
			n := ToNumByte(rule[1])
			m := ToNumByte(rule[2])
			if n > len(mangled)-1 || m > len(mangled)-1 {
				break
			}
			mangled = ExtractRange(mangled, n, m)
		case RULE_OP_MANGLE_OMIT:
			n := ToNumByte(rule[1])
			m := ToNumByte(rule[2])
			if n > len(mangled)-1 || m > len(mangled)-1 {
				break
			}
			mangled = OmitRange(mangled, n, m)
		case RULE_OP_MANGLE_INSERT:
			n := ToNumByte(rule[1])
			if n > len(mangled)-1 {
				break
			}
			mangled = InsertAtN(mangled, n, rune(rule[2]))
		case RULE_OP_MANGLE_OVERSTRIKE:
			n := ToNumByte(rule[1])
			if n > len(mangled)-1 {
				break
			}
			mangled = OverwriteAtN(mangled, n, rune(rule[2]))
		case RULE_OP_MANGLE_TRUNCATE_AT:
			n := ToNumByte(rule[1])
			if n > len(mangled)-1 {
				break
			}
			mangled = TruncateAtN(mangled, n)
		case RULE_OP_MANGLE_REPLACE:
			mangled = Replace(mangled, rune(rule[1]), rune(rule[2]))
		case RULE_OP_MANGLE_PURGECHAR:
			mangled = Purge(mangled, rune(rule[1]))
		case RULE_OP_MANGLE_DUPECHAR_FIRST:
			n := ToNumByte(rule[1])
			if n > len(mangled)-1 {
				break
			}
			mangled = DuplicateFirstN(mangled, n)
		case RULE_OP_MANGLE_DUPECHAR_LAST:
			n := ToNumByte(rule[1])
			if n > len(mangled)-1 {
				break
			}
			mangled = DuplicateLastN(mangled, n)
		case RULE_OP_MANGLE_DUPECHAR_ALL:
			mangled = DuplicateAll(mangled)
		case RULE_OP_MANGLE_EXTRACT_MEMORY:
			n := ToNumByte(rule[1])
			m := ToNumByte(rule[2])
			l := ToNumByte(rule[3])
			// this test is duplicated in ExtractMemory
			// -1 for the first two since they will be indexing the saved word
			/*
				if n > len(Saved) -1 || m > len(Saved) -1|| l > len(word) {
					mangled = ""
					break
				}
			*/
			mangled = ExtractMemory(mangled, n, m, l)
		case RULE_OP_MANGLE_APPEND_MEMORY:
			mangled = AppendMemory(mangled)
		case RULE_OP_MANGLE_PREPEND_MEMORY:
			mangled = PrependMemory(mangled)
		case RULE_OP_MEMORIZE_WORD:
			Memorize(mangled)
		case RULE_OP_REJECT_LESS:
			n := ToNumByte(rule[1])
			reject = RejectLess(mangled, n)
			if reject {
				return ""
			}
		case RULE_OP_REJECT_GREATER:
			n := ToNumByte(rule[1])
			reject = RejectGreater(mangled, n)
			if reject {
				return ""
			}
		case RULE_OP_REJECT_CONTAIN:
			reject = RejectContain(mangled, rune(rule[1]))
			if reject {
				return ""
			}
		case RULE_OP_REJECT_NOT_CONTAIN:
			reject = RejectNotContain(mangled, rune(rule[1]))
			if reject {
				return ""
			}
		case RULE_OP_REJECT_EQUAL_FIRST:
			reject = RejectEqualFirst(mangled, rune(rule[1]))
			if reject {
				return ""
			}
		case RULE_OP_REJECT_EQUAL_LAST:
			reject = RejectEqualLast(mangled, rune(rule[1]))
			if reject {
				return ""
			}
		case RULE_OP_REJECT_EQUAL_AT:
			n := ToNumByte(rule[1])
			if n > len(mangled)-1 {
				break
			}
			reject = RejectEqualAt(mangled, rune(rule[1]), n)
			if reject {
				return ""
			}
		case RULE_OP_REJECT_CONTAINS:
			n := ToNumByte(rule[1])
			if n > len(mangled)-1 {
				break
			}
			reject = RejectContains(mangled, rune(rule[1]), n)
			if reject {
				return ""
			}
		case RULE_OP_REJECT_MEMORY:
			reject = RejectMemory(mangled)
			if reject {
				return ""
			}
		case RULE_OP_MANGLE_SWITCH_FIRST:
			mangled = SwapFront(mangled)
		case RULE_OP_MANGLE_SWITCH_LAST:
			mangled = SwapBack(mangled)
		case RULE_OP_MANGLE_SWITCH_AT:
			n := ToNumByte(rule[1])
			m := ToNumByte(rule[2])
			if n > len(mangled)-1 || m > len(mangled)-1 {
				break
			}
			mangled = SwapAtN(mangled, n, m)
		case RULE_OP_MANGLE_CHR_SHIFTL:
			n := ToNumByte(rule[1])
			if n > len(mangled)-1 {
				break
			}
			mangled = BitwiseShiftLeft(mangled, n)
		case RULE_OP_MANGLE_CHR_SHIFTR:
			n := ToNumByte(rule[1])
			if n > len(mangled)-1 {
				break
			}
			mangled = BitwiseShiftRight(mangled, n)
		case RULE_OP_MANGLE_CHR_INCR:
			n := ToNumByte(rule[1])
			if n > len(mangled)-1 {
				break
			}
			mangled = ASCIIIncrementPlus(mangled, n)
		case RULE_OP_MANGLE_CHR_DECR:
			n := ToNumByte(rule[1])
			if n > len(mangled)-1 {
				break
			}
			mangled = ASCIIIncrementMinus(mangled, n)
		case RULE_OP_MANGLE_REPLACE_NP1:
			n := ToNumByte(rule[1])
			if n > len(mangled)-1 {
				break
			}
			mangled = ReplaceNPlus(mangled, n)
		case RULE_OP_MANGLE_REPLACE_NM1:
			n := ToNumByte(rule[1])
			if n > len(mangled)-1 {
				break
			}
			mangled = ReplaceNMinus(mangled, n)
		case RULE_OP_MANGLE_DUPEBLOCK_FIRST:
			n := ToNumByte(rule[1])
			if n > len(mangled)-1 {
				break
			}
			mangled = DuplicateBlockFront(mangled, n)
		case RULE_OP_MANGLE_DUPEBLOCK_LAST:
			n := ToNumByte(rule[1])
			if n > len(mangled)-1 {
				break
			}
			mangled = DuplicateBlockBack(mangled, n)
		case RULE_OP_MANGLE_TITLE:
			mangled = Title(mangled)
		//default:
		//	continue
		}
	}
	return string(mangled)
}

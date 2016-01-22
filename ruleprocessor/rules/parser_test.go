package rules

import (
	"testing"
)

// do generate rule
// then do some perl to make it like this

func TestParseLine(t *testing.T) {
	var ParseRules = []struct {
		rule    string
		spaced  string
		correct []string
	}{
		{"*69 *50 xB4", "*69*50xB4", []string{"*69", "*50", "xB4"}},
		{"o8X", "o8X", []string{"o8X"}},
		{"^6", "^6", []string{"^6"}},
		{"x3A z3", "x3Az3", []string{"x3A", "z3"}},
		{"o9L R1 ,9", "o9LR1,9", []string{"o9L", "R1", ",9"}},
		{"i7! XBB2", "i7!XBB2", []string{"i7!", "XBB2"}},
		{"p5 oA^", "p5oA^", []string{"p5", "oA^"}},
		{"x28 X114 z5", "x28X114z5", []string{"x28", "X114", "z5"}},
		{"x84", "x84", []string{"x84"}},
		{"sU, x67", "sU,x67", []string{"sU,", "x67"}},
		{"E d", "Ed", []string{"E", "d"}},
		{"X550 *83", "X550*83", []string{"X550", "*83"}},
		{"s!C", "s!C", []string{"s!C"}},
		{"^} E", "^}E", []string{"^}", "E"}},
		{"X663 6 z1", "X6636z1", []string{"X663", "6", "z1"}},
		{"swT ^e", "swT^e", []string{"swT", "^e"}},
		{"K i6 ", "Ki6 ", []string{"K", "i6 "}},
		{"-5 s{U ^Q", "-5s{U^Q", []string{"-5", "s{U", "^Q"}},
		{"E y5 X226", "Ey5X226", []string{"E", "y5", "X226"}},
		{"XBB3 @/", "XBB3@/", []string{"XBB3", "@/"}},
		{",8 @l o3i", ",8@lo3i", []string{",8", "@l", "o3i"}},
		{"*B2", "*B2", []string{"*B2"}},
		{"p4", "p4", []string{"p4"}},
		{"iAJ y4 $N", "iAJy4$N", []string{"iAJ", "y4", "$N"}},
		{"Y3 ^q XAA2", "Y3^qXAA2", []string{"Y3", "^q", "XAA2"}},
		{"Y1 ,5 i0W", "Y1,5i0W", []string{"Y1", ",5", "i0W"}},
		{"[ *8A @1", "[*8A@1", []string{"[", "*8A", "@1"}},
		{"x59", "x59", []string{"x59"}},
		{"sR1 X994", "sR1X994", []string{"sR1", "X994"}},
		{"x17", "x17", []string{"x17"}},
		{".8", ".8", []string{".8"}},
		{"*68 DB l", "*68DBl", []string{"*68", "DB", "l"}},
		{"s$\" o8S $3", "s$\"o8S$3", []string{"s$\"", "o8S", "$3"}},
		{"Z4", "Z4", []string{"Z4"}},
		{"*5A", "*5A", []string{"*5A"}},
		{"T7", "T7", []string{"T7"}},
		{"k", "k", []string{"k"}},
		{"xB8", "xB8", []string{"xB8"}},
		{"$0 s}]", "$0s}]", []string{"$0", "s}]"}},
		{"*49 sO' sV|", "*49sO'sV|", []string{"*49", "sO'", "sV|"}},
		{"Y3 @w", "Y3@w", []string{"Y3", "@w"}},
		{"x16 R2 i3F", "x16R2i3F", []string{"x16", "R2", "i3F"}},
		{"q", "q", []string{"q"}},
		{"^C E o4s", "^CEo4s", []string{"^C", "E", "o4s"}},
		{"Y3", "Y3", []string{"Y3"}},
		{"^3 T6 p4", "^3T6p4", []string{"^3", "T6", "p4"}},
		{"x94 x73 i1>", "x94x73i1>", []string{"x94", "x73", "i1>"}},
		{"{ f z4", "{fz4", []string{"{", "f", "z4"}},
		{"X009 f", "X009f", []string{"X009", "f"}},
		{"X661 *42 $R", "X661*42$R", []string{"X661", "*42", "$R"}},
		{"q c *A9", "qc*A9", []string{"q", "c", "*A9"}},
		{"C", "C", []string{"C"}},
		{"oAU ,7", "oAU,7", []string{"oAU", ",7"}},
		{"i3} x59 X889", "i3}x59X889", []string{"i3}", "x59", "X889"}},
		{"*B6", "*B6", []string{"*B6"}},
		{"4 *39", "4*39", []string{"4", "*39"}},
		{"x62 X227", "x62X227", []string{"x62", "X227"}},
		{"$\\ Y2 Y4", "$\\Y2Y4", []string{"$\\", "Y2", "Y4"}},
		{"^&", "^&", []string{"^&"}},
		{"X003", "X003", []string{"X003"}},
		{"Z2", "Z2", []string{"Z2"}},
		{"L2 T1 x64", "L2T1x64", []string{"L2", "T1", "x64"}},
		{"z4 @;", "z4@;", []string{"z4", "@;"}},
		{"'B x21 .B", "'Bx21.B", []string{"'B", "x21", ".B"}},
		{"R1 Y1 C", "R1Y1C", []string{"R1", "Y1", "C"}},
		{"x81", "x81", []string{"x81"}},
		{"sry x2B oBV", "sryx2BoBV", []string{"sry", "x2B", "oBV"}},
		{"f ,5", "f,5", []string{"f", ",5"}},
		{"Y1 o5< i9Q", "Y1o5<i9Q", []string{"Y1", "o5<", "i9Q"}},
		{"Y4 sS^ .B", "Y4sS^.B", []string{"Y4", "sS^", ".B"}},
		{"x75 o8A", "x75o8A", []string{"x75", "o8A"}},
		{"y5 f", "y5f", []string{"y5", "f"}},
		{"^b Y5", "^bY5", []string{"^b", "Y5"}},
		{"o3o @b l", "o3o@bl", []string{"o3o", "@b", "l"}},
		{"s9H sjO XAA0", "s9HsjOXAA0", []string{"s9H", "sjO", "XAA0"}},
		{"K Y2 $L", "KY2$L", []string{"K", "Y2", "$L"}},
		{"} +3", "}+3", []string{"}", "+3"}},
		{"$U", "$U", []string{"$U"}},
		{"i2{", "i2{", []string{"i2{"}},
		{"Z3 ^& *9B", "Z3^&*9B", []string{"Z3", "^&", "*9B"}},
		{"l i1U x4B", "li1Ux4B", []string{"l", "i1U", "x4B"}},
		{"X220 'B", "X220'B", []string{"X220", "'B"}},
		{"@`", "@`", []string{"@`"}},
		{"*13", "*13", []string{"*13"}},
		{"C p1", "Cp1", []string{"C", "p1"}},
		{"*84 X555", "*84X555", []string{"*84", "X555"}},
		{"X441", "X441", []string{"X441"}},
		{"p4", "p4", []string{"p4"}},
		{"s2F X990", "s2FX990", []string{"s2F", "X990"}},
		{"x78", "x78", []string{"x78"}},
		{"6", "6", []string{"6"}},
		{"*25 r ^h", "*25r^h", []string{"*25", "r", "^h"}},
		{",0 @} c", ",0@}c", []string{",0", "@}", "c"}},
		{"c o8$ *3A", "co8$*3A", []string{"c", "o8$", "*3A"}},
		{"@K", "@K", []string{"@K"}},
		{"o3f", "o3f", []string{"o3f"}},
		{"i6H *A5", "i6H*A5", []string{"i6H", "*A5"}},
		{"X991 Z3", "X991Z3", []string{"X991", "Z3"}},
		{"*01 ^J K", "*01^JK", []string{"*01", "^J", "K"}},
		{"X336", "X336", []string{"X336"}},

		// made by hand
		{"T[", "T[", []string{}},
		{"T/", "T/", []string{}},
		{"T:", "T:", []string{}},

		{"p[", "p[", []string{}},
		{"p/", "p/", []string{}},
		{"p:", "p:", []string{}},

		{"D[", "D[", []string{}},
		{"D/", "D/", []string{}},
		{"D:", "D:", []string{}},

		{"x[", "x[", []string{}},
		{"x/", "x/", []string{}},
		{"x:", "x:", []string{}},
		{"x[]", "x[]", []string{}},
		{"x2-", "x2-", []string{}},
		{"x2.", "x2>", []string{}},

		{"i[", "i[", []string{}},
		{"i/", "i/", []string{}},
		{"i:", "i:", []string{}},

		{"o[", "o[", []string{}},
		{"o/", "o/", []string{}},
		{"o:", "o:", []string{}},

		{"'[", "'[", []string{}},
		{"'/", "'/", []string{}},
		{"':", "':", []string{}},

		{"z[", "z[", []string{}},
		{"z/", "z/", []string{}},
		{"z:", "z:", []string{}},

		{"Z[", "Z[", []string{}},
		{"Z/", "Z/", []string{}},
		{"Z:", "Z:", []string{}},

		{"X[", "X[", []string{}},
		{"X/", "X/", []string{}},
		{"X:", "X:", []string{}},
		{"X2[", "X2[", []string{}},
		{"X2/", "X2/", []string{}},
		{"X2:", "X2:", []string{}},
		{"X22[", "X22[", []string{}},
		{"X22/", "X22/", []string{}},
		{"X22:", "X22:", []string{}},

		{"<[", "<[", []string{}},
		{"</", "</", []string{}},
		{"<:", "<:", []string{}},

		{">[", ">[", []string{}},
		{">/", ">/", []string{}},
		{">:", ">:", []string{}},

		{"=[a", "=[a", []string{}},
		{"=/a", "=/a", []string{}},
		{"=:a", "=:a", []string{}},

		{"%[a", "%[a", []string{}},
		{"%/a", "%/a", []string{}},
		{"%:a", "%:a", []string{}},

		{"*[", "*[", []string{}},
		{"*-", "*-", []string{}},
		{"*:", "*:", []string{}},

		{"2]", "*2]", []string{}},
		{"*2-", "*2-", []string{}},
		{"*2.", "*2>", []string{}},

		{"L[", "L[", []string{}},
		{"L/", "L/", []string{}},
		{"L:", "L:", []string{}},

		{"R[", "R[", []string{}},
		{"R/", "R/", []string{}},
		{"R:", "R:", []string{}},

		{"+[", "+[", []string{}},
		{"+/", "+/", []string{}},
		{"+:", "+:", []string{}},

		{"-[", "-[", []string{}},
		{"-/", "-/", []string{}},
		{"-:", "-:", []string{}},

		{".[", ".[", []string{}},
		{"./", "./", []string{}},
		{".:", ".:", []string{}},

		{",[", ",[", []string{}},
		{",/", ",/", []string{}},
		{",:", ",:", []string{}},

		{"y[", "y[", []string{}},
		{"y/", "y/", []string{}},
		{"y:", "y:", []string{}},

		{"Y[", "Y[", []string{}},
		{"Y/", "Y/", []string{}},
		{"Y:", "Y:", []string{}},


		// this is for the new `O' rule
		// should add better tests
		{"O12", "O12", []string{"O12"}},
	}

	for _, rule := range ParseRules {
		parsed := ParseLine(rule.rule)
		parsedSpaced := ParseLine(rule.spaced)
		/*
			if len(parsed) == 0 {
				t.Error("parsed is len 0")
			}
			if len(parsedSpaced) == 0 {
				t.Error("parsed is len 0")
			}
		*/
		// need to make a loop for comparison here
		// https://stackoverflow.com/questions/15311969/checking-the-equality-of-two-slices
		for i := 0; i < len(rule.correct); i++ {

			if parsed[i] != rule.correct[i] {
				t.Errorf("parsed not correct, got %s, expected %s, correct %s, %s", parsed[i], rule.correct[i], rule.correct, parsed)
			}

			if parsedSpaced[i] != rule.correct[i] {
				t.Errorf("parsed not correct, got %s, expected %s, correct %s, %s", parsedSpaced[i], rule.correct[i], rule.correct, parsedSpaced)
			}
		}
	}
}

func TestApplyRules(t *testing.T) {
	// this was just taken from the rules wiki page
	var RuleTest = []struct {
		plain   string
		rule    string
		mangled string
	}{
		{"p@ssW0rd", ":", "p@ssW0rd"},
		{"p@ssW0rd", "l", "p@ssw0rd"},
		{"p@ssW0rd", "u", "P@SSW0RD"},
		{"p@ssW0rd", "c", "P@ssw0rd"},
		{"p@ssW0rd", "C", "p@SSW0RD"},
		{"p@ssW0rd", "t", "P@SSw0RD"},
		{"p@ssW0rd", "T3", "p@sSW0rd"},
		{"p@ssW0rd", "r", "dr0Wss@p"},
		{"p@ssW0rd", "d", "p@ssW0rdp@ssW0rd"},
		{"p@ssW0rd", "p2", "p@ssW0rdp@ssW0rdp@ssW0rd"},
		{"p@ssW0rd", "f", "p@ssW0rddr0Wss@p"},
		{"p@ssW0rd", "{", "@ssW0rdp"},
		{"p@ssW0rd", "}", "dp@ssW0r"},
		{"p@ssW0rd", "$1", "p@ssW0rd1"},
		{"p@ssW0rd", "^1", "1p@ssW0rd"},
		{"p@ssW0rd", "[", "@ssW0rd"},
		{"p@ssW0rd", "]", "p@ssW0r"},
		{"p@ssW0rd", "D3", "p@sW0rd"},
		{"p@ssW0rd", "x04", "p@ss"},
		{"p@ssW0rd", "O12", "psW0rd"},
		{"p@ssW0rd", "i4!", "p@ss!W0rd"},
		{"p@ssW0rd", "o3$", "p@s$W0rd"},
		{"p@ssW0rd", "'6", "p@ssW0"},
		{"p@ssW0rd", "ss$", "p@$$W0rd"},
		{"p@ssW0rd", "@s", "p@W0rd"},
		{"p@ssW0rd", "z2", "ppp@ssW0rd"},
		{"p@ssW0rd", "Z2", "p@ssW0rddd"},
		{"p@ssW0rd", "q", "pp@@ssssWW00rrdd"},
		{"p@ssW0rd", "lMX428", "p@ssw0rdw0"},
		{"p@ssW0rd", "uMl4", "p@ssw0rdP@SSW0RD"},
		{"p@ssW0rd", "rMr6", "dr0Wss@pp@ssW0rd"},
		{"p@ssW0rd", "lMuX084", "P@SSp@ssw0rdW0RD"},

		// reject
		{"p@ssW0rd", "<1", ""},
		{"p@ssW0rd", ">7", "p@ssW0rd"},
		{"p@ssW0rd", "!z", "p@ssW0rd"},
		{"p@ssW0rd", "/e", ""},
		{"p@ssW0rd", "(h", ""},
		{"p@ssW0rd", ")t", ""},
		{"p@ssW0rd", "=1a", ""},
		{"p@ssW0rd", "%2a", ""},
		{"p@ssW0rd", "rMrQ", "p@ssW0rd"},

		// hashcat only
		{"p@ssW0rd", "k", "@pssW0rd"},
		{"p@ssW0rd", "K", "p@ssW0dr"},
		{"p@ssW0rd", "*34", "p@sWs0rd"},
		{"p@ssW0rd", "L2", "p@æsW0rd"},
		{"p@ssW0rd", "R2", "p@9sW0rd"},
		{"p@ssW0rd", "+2", "p@tsW0rd"},
		{"p@ssW0rd", "-1", "p?ssW0rd"},
		{"p@ssW0rd", ".1", "psssW0rd"},
		{"p@ssW0rd", ",1", "ppssW0rd"},
		{"p@ssW0rd", "y2", "p@p@ssW0rd"},
		{"p@ssW0rd", "Y2", "p@ssW0rdrd"},
		{"p@ssW0rd w0rld", "E", "P@ssw0rd W0rld"},

		// made by hand
		{"p@ssW0rd", "TF", "p@ssW0rd"},
		{"p@ssW0rd", "DF", "p@ssW0rd"},

		{"p@ssW0rd", "xF2", "p@ssW0rd"},
		{"p@ssW0rd", "x0F", "p@ssW0rd"},

		{"p@ssW0rd", "iF!", "p@ssW0rd"},
		{"p@ssW0rd", "oF$", "p@ssW0rd"},
		{"p@ssW0rd", "'F", "p@ssW0rd"},

		{"p@ssW0rd", "zF", "p@ssW0rd"},
		{"p@ssW0rd", "ZF", "p@ssW0rd"},

		{"p@ssW0rd", "lMXF84", ""},
		{"p@ssW0rd", "lMX0F4", ""},
		{"p@ssW0rd", "lMX02F", ""},

		{"p@ssW0rd", "<F", "p@ssW0rd"},
		{"p@ssW0rd", ">F", ""},

		{"p@ssW0rd", "!0", ""},

		{"p@ssW0rd", "=Fp", "p@ssW0rd"},
		{"p@ssW0rd", "%F0", "p@ssW0rd"},

		{"p@ssW0rd", "rMQ", ""},

		{"p@ssW0rd", "*F4", "p@ssW0rd"},
		{"p@ssW0rd", "*2F", "p@ssW0rd"},

		{"p@ssW0rd", "LF", "p@ssW0rd"},
		{"p@ssW0rd", "RF", "p@ssW0rd"},

		{"p@ssW0rd", "+F", "p@ssW0rd"},
		{"p@ssW0rd", "-F", "p@ssW0rd"},

		{"p@ssW0rd", ".F", "p@ssW0rd"},
		{"p@ssW0rd", ",F", "p@ssW0rd"},

		{"p@ssW0rd", "yF", "p@ssW0rd"},
		{"p@ssW0rd", "YF", "p@ssW0rd"},

		{"p@ssW0rd", "rMQ", ""},
	}

	for _, rule := range RuleTest {
		parsed := ParseLine(rule.rule)
		apply := ApplyRules(parsed, rule.plain)

		if apply != rule.mangled {
			t.Errorf("rule: %s\nword: %s\nmangled: %s\nshould be: %s\n", rule.rule, rule.plain, apply, rule.mangled)
		}
	}

	// this tests default
	var RuleTests = struct {
		plain   string
		rule    []string
		mangled string
	}{
		"p@ssW0rd", []string{"v"}, "p@ssW0rd",
	}
	apply := ApplyRules(RuleTests.rule, RuleTests.plain)
	if apply != RuleTests.mangled {
		t.Errorf("rule: %s\nword: %s\nmangled: %s\nshould be: %s\n", RuleTests.rule, RuleTests.plain, apply, RuleTests.mangled)
	}
}

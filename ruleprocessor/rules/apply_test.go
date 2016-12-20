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
	alpha := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	number := []rune("0123456789")

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
		{[]rune("own"), []rune("own")},
		{[]rune("sp33d"), []rune("sp33d")},
		{[]rune("roearoea"), []rune("roearoea")},
		{[]rune("5556662"), []rune("5556662")},
		{[]rune("yrallih"), []rune("yrallih")},
		{[]rune("gemma8"), []rune("gemma8")},
		{[]rune("love4ever21"), []rune("love4ever21")},
		{[]rune("telephone12"), []rune("telephone12")},
		{[]rune("51515"), []rune("51515")},
		{[]rune("vnvnvn"), []rune("vnvnvn")},
		{[]rune("fds2fd"), []rune("fds2fd")},
		{[]rune("`amur"), []rune("`amur")},
		{[]rune("soccer0723"), []rune("soccer0723")},
		{[]rune("paroach"), []rune("paroach")},
		{[]rune("Doritos"), []rune("Doritos")},
		{[]rune("skiper"), []rune("skiper")},
		{[]rune("thx1man"), []rune("thx1man")},
		{[]rune("chula4"), []rune("chula4")},
		{[]rune("1trebla"), []rune("1trebla")},
		{[]rune("nancy21"), []rune("nancy21")},
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
		{[]rune("own"), []rune("own")},
		{[]rune("sp33d"), []rune("sp33d")},
		{[]rune("roearoea"), []rune("roearoea")},
		{[]rune("5556662"), []rune("5556662")},
		{[]rune("yrallih"), []rune("yrallih")},
		{[]rune("gemma8"), []rune("gemma8")},
		{[]rune("love4ever21"), []rune("love4ever21")},
		{[]rune("telephone12"), []rune("telephone12")},
		{[]rune("51515"), []rune("51515")},
		{[]rune("vnvnvn"), []rune("vnvnvn")},
		{[]rune("fds2fd"), []rune("fds2fd")},
		{[]rune("`amur"), []rune("`amur")},
		{[]rune("soccer0723"), []rune("soccer0723")},
		{[]rune("paroach"), []rune("paroach")},
		{[]rune("Doritos"), []rune("doritos")},
		{[]rune("skiper"), []rune("skiper")},
		{[]rune("thx1man"), []rune("thx1man")},
		{[]rune("chula4"), []rune("chula4")},
		{[]rune("1trebla"), []rune("1trebla")},
		{[]rune("nancy21"), []rune("nancy21")},
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
		{[]rune("own"), []rune("OWN")},
		{[]rune("sp33d"), []rune("SP33D")},
		{[]rune("roearoea"), []rune("ROEAROEA")},
		{[]rune("5556662"), []rune("5556662")},
		{[]rune("yrallih"), []rune("YRALLIH")},
		{[]rune("gemma8"), []rune("GEMMA8")},
		{[]rune("love4ever21"), []rune("LOVE4EVER21")},
		{[]rune("telephone12"), []rune("TELEPHONE12")},
		{[]rune("51515"), []rune("51515")},
		{[]rune("vnvnvn"), []rune("VNVNVN")},
		{[]rune("fds2fd"), []rune("FDS2FD")},
		{[]rune("`amur"), []rune("`AMUR")},
		{[]rune("soccer0723"), []rune("SOCCER0723")},
		{[]rune("paroach"), []rune("PAROACH")},
		{[]rune("Doritos"), []rune("DORITOS")},
		{[]rune("skiper"), []rune("SKIPER")},
		{[]rune("thx1man"), []rune("THX1MAN")},
		{[]rune("chula4"), []rune("CHULA4")},
		{[]rune("1trebla"), []rune("1TREBLA")},
		{[]rune("nancy21"), []rune("NANCY21")},
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
		{[]rune("own"), []rune("Own")},
		{[]rune("sp33d"), []rune("Sp33d")},
		{[]rune("roearoea"), []rune("Roearoea")},
		{[]rune("5556662"), []rune("5556662")},
		{[]rune("yrallih"), []rune("Yrallih")},
		{[]rune("gemma8"), []rune("Gemma8")},
		{[]rune("love4ever21"), []rune("Love4ever21")},
		{[]rune("telephone12"), []rune("Telephone12")},
		{[]rune("51515"), []rune("51515")},
		{[]rune("vnvnvn"), []rune("Vnvnvn")},
		{[]rune("fds2fd"), []rune("Fds2fd")},
		{[]rune("`amur"), []rune("`amur")},
		{[]rune("soccer0723"), []rune("Soccer0723")},
		{[]rune("paroach"), []rune("Paroach")},
		{[]rune("Doritos"), []rune("Doritos")},
		{[]rune("skiper"), []rune("Skiper")},
		{[]rune("thx1man"), []rune("Thx1man")},
		{[]rune("chula4"), []rune("Chula4")},
		{[]rune("1trebla"), []rune("1trebla")},
		{[]rune("nancy21"), []rune("Nancy21")},
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
		{[]rune("own"), []rune("oWN")},
		{[]rune("sp33d"), []rune("sP33D")},
		{[]rune("roearoea"), []rune("rOEAROEA")},
		{[]rune("5556662"), []rune("5556662")},
		{[]rune("yrallih"), []rune("yRALLIH")},
		{[]rune("gemma8"), []rune("gEMMA8")},
		{[]rune("love4ever21"), []rune("lOVE4EVER21")},
		{[]rune("telephone12"), []rune("tELEPHONE12")},
		{[]rune("51515"), []rune("51515")},
		{[]rune("vnvnvn"), []rune("vNVNVN")},
		{[]rune("fds2fd"), []rune("fDS2FD")},
		{[]rune("`amur"), []rune("`AMUR")},
		{[]rune("soccer0723"), []rune("sOCCER0723")},
		{[]rune("paroach"), []rune("pAROACH")},
		{[]rune("Doritos"), []rune("dORITOS")},
		{[]rune("skiper"), []rune("sKIPER")},
		{[]rune("thx1man"), []rune("tHX1MAN")},
		{[]rune("chula4"), []rune("cHULA4")},
		{[]rune("1trebla"), []rune("1TREBLA")},
		{[]rune("nancy21"), []rune("nANCY21")},
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
		{[]rune("own"), []rune("OWN")},
		{[]rune("sp33d"), []rune("SP33D")},
		{[]rune("roearoea"), []rune("ROEAROEA")},
		{[]rune("5556662"), []rune("5556662")},
		{[]rune("yrallih"), []rune("YRALLIH")},
		{[]rune("gemma8"), []rune("GEMMA8")},
		{[]rune("love4ever21"), []rune("LOVE4EVER21")},
		{[]rune("telephone12"), []rune("TELEPHONE12")},
		{[]rune("51515"), []rune("51515")},
		{[]rune("vnvnvn"), []rune("VNVNVN")},
		{[]rune("fds2fd"), []rune("FDS2FD")},
		{[]rune("`amur"), []rune("`AMUR")},
		{[]rune("soccer0723"), []rune("SOCCER0723")},
		{[]rune("paroach"), []rune("PAROACH")},
		{[]rune("Doritos"), []rune("dORITOS")},
		{[]rune("skiper"), []rune("SKIPER")},
		{[]rune("thx1man"), []rune("THX1MAN")},
		{[]rune("chula4"), []rune("CHULA4")},
		{[]rune("1trebla"), []rune("1TREBLA")},
		{[]rune("nancy21"), []rune("NANCY21")},
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
		{[]rune("own"), []rune("own")},
		{[]rune("sp33d"), []rune("sp33d")},
		{[]rune("roearoea"), []rune("roeAroea")},
		{[]rune("5556662"), []rune("5556662")},
		{[]rune("yrallih"), []rune("yraLlih")},
		{[]rune("gemma8"), []rune("gemMa8")},
		{[]rune("love4ever21"), []rune("lovE4ever21")},
		{[]rune("telephone12"), []rune("telEphone12")},
		{[]rune("51515"), []rune("51515")},
		{[]rune("vnvnvn"), []rune("vnvNvn")},
		{[]rune("fds2fd"), []rune("fds2fd")},
		{[]rune("`amur"), []rune("`amUr")},
		{[]rune("soccer0723"), []rune("socCer0723")},
		{[]rune("paroach"), []rune("parOach")},
		{[]rune("Doritos"), []rune("DorItos")},
		{[]rune("skiper"), []rune("skiPer")},
		{[]rune("thx1man"), []rune("thx1man")},
		{[]rune("chula4"), []rune("chuLa4")},
		{[]rune("1trebla"), []rune("1trEbla")},
		{[]rune("nancy21"), []rune("nanCy21")},

		// made by hand
		{[]rune("THREASD"), []rune("THReASD")},
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
		{[]rune("own"), []rune("nwo")},
		{[]rune("sp33d"), []rune("d33ps")},
		{[]rune("roearoea"), []rune("aeoraeor")},
		{[]rune("5556662"), []rune("2666555")},
		{[]rune("yrallih"), []rune("hillary")},
		{[]rune("gemma8"), []rune("8ammeg")},
		{[]rune("love4ever21"), []rune("12reve4evol")},
		{[]rune("telephone12"), []rune("21enohpelet")},
		{[]rune("51515"), []rune("51515")},
		{[]rune("vnvnvn"), []rune("nvnvnv")},
		{[]rune("fds2fd"), []rune("df2sdf")},
		{[]rune("`amur"), []rune("ruma`")},
		{[]rune("soccer0723"), []rune("3270reccos")},
		{[]rune("paroach"), []rune("hcaorap")},
		{[]rune("Doritos"), []rune("sotiroD")},
		{[]rune("skiper"), []rune("repiks")},
		{[]rune("thx1man"), []rune("nam1xht")},
		{[]rune("chula4"), []rune("4aluhc")},
		{[]rune("1trebla"), []rune("albert1")},
		{[]rune("nancy21"), []rune("12ycnan")},
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
		{[]rune("own"), []rune("ownown")},
		{[]rune("sp33d"), []rune("sp33dsp33d")},
		{[]rune("roearoea"), []rune("roearoearoearoea")},
		{[]rune("5556662"), []rune("55566625556662")},
		{[]rune("yrallih"), []rune("yrallihyrallih")},
		{[]rune("gemma8"), []rune("gemma8gemma8")},
		{[]rune("love4ever21"), []rune("love4ever21love4ever21")},
		{[]rune("telephone12"), []rune("telephone12telephone12")},
		{[]rune("51515"), []rune("5151551515")},
		{[]rune("vnvnvn"), []rune("vnvnvnvnvnvn")},
		{[]rune("fds2fd"), []rune("fds2fdfds2fd")},
		{[]rune("`amur"), []rune("`amur`amur")},
		{[]rune("soccer0723"), []rune("soccer0723soccer0723")},
		{[]rune("paroach"), []rune("paroachparoach")},
		{[]rune("Doritos"), []rune("DoritosDoritos")},
		{[]rune("skiper"), []rune("skiperskiper")},
		{[]rune("thx1man"), []rune("thx1manthx1man")},
		{[]rune("chula4"), []rune("chula4chula4")},
		{[]rune("1trebla"), []rune("1trebla1trebla")},
		{[]rune("nancy21"), []rune("nancy21nancy21")},
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
		{[]rune("own"), []rune("ownownown")},
		{[]rune("sp33d"), []rune("sp33dsp33dsp33d")},
		{[]rune("roearoea"), []rune("roearoearoearoearoearoea")},
		{[]rune("5556662"), []rune("555666255566625556662")},
		{[]rune("yrallih"), []rune("yrallihyrallihyrallih")},
		{[]rune("gemma8"), []rune("gemma8gemma8gemma8")},
		{[]rune("love4ever21"), []rune("love4ever21love4ever21love4ever21")},
		{[]rune("telephone12"), []rune("telephone12telephone12telephone12")},
		{[]rune("51515"), []rune("515155151551515")},
		{[]rune("vnvnvn"), []rune("vnvnvnvnvnvnvnvnvn")},
		{[]rune("fds2fd"), []rune("fds2fdfds2fdfds2fd")},
		{[]rune("`amur"), []rune("`amur`amur`amur")},
		{[]rune("soccer0723"), []rune("soccer0723soccer0723soccer0723")},
		{[]rune("paroach"), []rune("paroachparoachparoach")},
		{[]rune("Doritos"), []rune("DoritosDoritosDoritos")},
		{[]rune("skiper"), []rune("skiperskiperskiper")},
		{[]rune("thx1man"), []rune("thx1manthx1manthx1man")},
		{[]rune("chula4"), []rune("chula4chula4chula4")},
		{[]rune("1trebla"), []rune("1trebla1trebla1trebla")},
		{[]rune("nancy21"), []rune("nancy21nancy21nancy21")},
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
		{[]rune("own"), []rune("ownnwo")},
		{[]rune("sp33d"), []rune("sp33dd33ps")},
		{[]rune("roearoea"), []rune("roearoeaaeoraeor")},
		{[]rune("5556662"), []rune("55566622666555")},
		{[]rune("yrallih"), []rune("yrallihhillary")},
		{[]rune("gemma8"), []rune("gemma88ammeg")},
		{[]rune("love4ever21"), []rune("love4ever2112reve4evol")},
		{[]rune("telephone12"), []rune("telephone1221enohpelet")},
		{[]rune("51515"), []rune("5151551515")},
		{[]rune("vnvnvn"), []rune("vnvnvnnvnvnv")},
		{[]rune("fds2fd"), []rune("fds2fddf2sdf")},
		{[]rune("`amur"), []rune("`amurruma`")},
		{[]rune("soccer0723"), []rune("soccer07233270reccos")},
		{[]rune("paroach"), []rune("paroachhcaorap")},
		{[]rune("Doritos"), []rune("DoritossotiroD")},
		{[]rune("skiper"), []rune("skiperrepiks")},
		{[]rune("thx1man"), []rune("thx1mannam1xht")},
		{[]rune("chula4"), []rune("chula44aluhc")},
		{[]rune("1trebla"), []rune("1treblaalbert1")},
		{[]rune("nancy21"), []rune("nancy2112ycnan")},
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
		{[]rune("own"), []rune("wno")},
		{[]rune("sp33d"), []rune("p33ds")},
		{[]rune("roearoea"), []rune("oearoear")},
		{[]rune("5556662"), []rune("5566625")},
		{[]rune("yrallih"), []rune("rallihy")},
		{[]rune("gemma8"), []rune("emma8g")},
		{[]rune("love4ever21"), []rune("ove4ever21l")},
		{[]rune("telephone12"), []rune("elephone12t")},
		{[]rune("51515"), []rune("15155")},
		{[]rune("vnvnvn"), []rune("nvnvnv")},
		{[]rune("fds2fd"), []rune("ds2fdf")},
		{[]rune("`amur"), []rune("amur`")},
		{[]rune("soccer0723"), []rune("occer0723s")},
		{[]rune("paroach"), []rune("aroachp")},
		{[]rune("Doritos"), []rune("oritosD")},
		{[]rune("skiper"), []rune("kipers")},
		{[]rune("thx1man"), []rune("hx1mant")},
		{[]rune("chula4"), []rune("hula4c")},
		{[]rune("1trebla"), []rune("trebla1")},
		{[]rune("nancy21"), []rune("ancy21n")},
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
		{[]rune("own"), []rune("now")},
		{[]rune("sp33d"), []rune("dsp33")},
		{[]rune("roearoea"), []rune("aroearoe")},
		{[]rune("5556662"), []rune("2555666")},
		{[]rune("yrallih"), []rune("hyralli")},
		{[]rune("gemma8"), []rune("8gemma")},
		{[]rune("love4ever21"), []rune("1love4ever2")},
		{[]rune("telephone12"), []rune("2telephone1")},
		{[]rune("51515"), []rune("55151")},
		{[]rune("vnvnvn"), []rune("nvnvnv")},
		{[]rune("fds2fd"), []rune("dfds2f")},
		{[]rune("`amur"), []rune("r`amu")},
		{[]rune("soccer0723"), []rune("3soccer072")},
		{[]rune("paroach"), []rune("hparoac")},
		{[]rune("Doritos"), []rune("sDorito")},
		{[]rune("skiper"), []rune("rskipe")},
		{[]rune("thx1man"), []rune("nthx1ma")},
		{[]rune("chula4"), []rune("4chula")},
		{[]rune("1trebla"), []rune("a1trebl")},
		{[]rune("nancy21"), []rune("1nancy2")},
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
		{[]rune("own"), []rune("own1")},
		{[]rune("sp33d"), []rune("sp33d1")},
		{[]rune("roearoea"), []rune("roearoea1")},
		{[]rune("5556662"), []rune("55566621")},
		{[]rune("yrallih"), []rune("yrallih1")},
		{[]rune("gemma8"), []rune("gemma81")},
		{[]rune("love4ever21"), []rune("love4ever211")},
		{[]rune("telephone12"), []rune("telephone121")},
		{[]rune("51515"), []rune("515151")},
		{[]rune("vnvnvn"), []rune("vnvnvn1")},
		{[]rune("fds2fd"), []rune("fds2fd1")},
		{[]rune("`amur"), []rune("`amur1")},
		{[]rune("soccer0723"), []rune("soccer07231")},
		{[]rune("paroach"), []rune("paroach1")},
		{[]rune("Doritos"), []rune("Doritos1")},
		{[]rune("skiper"), []rune("skiper1")},
		{[]rune("thx1man"), []rune("thx1man1")},
		{[]rune("chula4"), []rune("chula41")},
		{[]rune("1trebla"), []rune("1trebla1")},
		{[]rune("nancy21"), []rune("nancy211")},
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
		{[]rune("own"), []rune("1own")},
		{[]rune("sp33d"), []rune("1sp33d")},
		{[]rune("roearoea"), []rune("1roearoea")},
		{[]rune("5556662"), []rune("15556662")},
		{[]rune("yrallih"), []rune("1yrallih")},
		{[]rune("gemma8"), []rune("1gemma8")},
		{[]rune("love4ever21"), []rune("1love4ever21")},
		{[]rune("telephone12"), []rune("1telephone12")},
		{[]rune("51515"), []rune("151515")},
		{[]rune("vnvnvn"), []rune("1vnvnvn")},
		{[]rune("fds2fd"), []rune("1fds2fd")},
		{[]rune("`amur"), []rune("1`amur")},
		{[]rune("soccer0723"), []rune("1soccer0723")},
		{[]rune("paroach"), []rune("1paroach")},
		{[]rune("Doritos"), []rune("1Doritos")},
		{[]rune("skiper"), []rune("1skiper")},
		{[]rune("thx1man"), []rune("1thx1man")},
		{[]rune("chula4"), []rune("1chula4")},
		{[]rune("1trebla"), []rune("11trebla")},
		{[]rune("nancy21"), []rune("1nancy21")},
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
		{[]rune("own"), []rune("wn")},
		{[]rune("sp33d"), []rune("p33d")},
		{[]rune("roearoea"), []rune("oearoea")},
		{[]rune("5556662"), []rune("556662")},
		{[]rune("yrallih"), []rune("rallih")},
		{[]rune("gemma8"), []rune("emma8")},
		{[]rune("love4ever21"), []rune("ove4ever21")},
		{[]rune("telephone12"), []rune("elephone12")},
		{[]rune("51515"), []rune("1515")},
		{[]rune("vnvnvn"), []rune("nvnvn")},
		{[]rune("fds2fd"), []rune("ds2fd")},
		{[]rune("`amur"), []rune("amur")},
		{[]rune("soccer0723"), []rune("occer0723")},
		{[]rune("paroach"), []rune("aroach")},
		{[]rune("Doritos"), []rune("oritos")},
		{[]rune("skiper"), []rune("kiper")},
		{[]rune("thx1man"), []rune("hx1man")},
		{[]rune("chula4"), []rune("hula4")},
		{[]rune("1trebla"), []rune("trebla")},
		{[]rune("nancy21"), []rune("ancy21")},
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
		{[]rune("own"), []rune("ow")},
		{[]rune("sp33d"), []rune("sp33")},
		{[]rune("roearoea"), []rune("roearoe")},
		{[]rune("5556662"), []rune("555666")},
		{[]rune("yrallih"), []rune("yralli")},
		{[]rune("gemma8"), []rune("gemma")},
		{[]rune("love4ever21"), []rune("love4ever2")},
		{[]rune("telephone12"), []rune("telephone1")},
		{[]rune("51515"), []rune("5151")},
		{[]rune("vnvnvn"), []rune("vnvnv")},
		{[]rune("fds2fd"), []rune("fds2f")},
		{[]rune("`amur"), []rune("`amu")},
		{[]rune("soccer0723"), []rune("soccer072")},
		{[]rune("paroach"), []rune("paroac")},
		{[]rune("Doritos"), []rune("Dorito")},
		{[]rune("skiper"), []rune("skipe")},
		{[]rune("thx1man"), []rune("thx1ma")},
		{[]rune("chula4"), []rune("chula")},
		{[]rune("1trebla"), []rune("1trebl")},
		{[]rune("nancy21"), []rune("nancy2")},
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
		{[]rune("own"), []rune("own")},
		{[]rune("sp33d"), []rune("sp3d")},
		{[]rune("roearoea"), []rune("roeroea")},
		{[]rune("5556662"), []rune("555662")},
		{[]rune("yrallih"), []rune("yralih")},
		{[]rune("gemma8"), []rune("gema8")},
		{[]rune("love4ever21"), []rune("lov4ever21")},
		{[]rune("telephone12"), []rune("telphone12")},
		{[]rune("51515"), []rune("5155")},
		{[]rune("vnvnvn"), []rune("vnvvn")},
		{[]rune("fds2fd"), []rune("fdsfd")},
		{[]rune("`amur"), []rune("`amr")},
		{[]rune("soccer0723"), []rune("socer0723")},
		{[]rune("paroach"), []rune("parach")},
		{[]rune("Doritos"), []rune("Dortos")},
		{[]rune("skiper"), []rune("skier")},
		{[]rune("thx1man"), []rune("thxman")},
		{[]rune("chula4"), []rune("chua4")},
		{[]rune("1trebla"), []rune("1trbla")},
		{[]rune("nancy21"), []rune("nany21")},
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

		{[]rune("own"), []rune("ow")},
		{[]rune("sp33d"), []rune("sp")},
		{[]rune("roearoea"), []rune("ro")},
		{[]rune("5556662"), []rune("55")},
		{[]rune("yrallih"), []rune("yr")},
		{[]rune("gemma8"), []rune("ge")},
		{[]rune("love4ever21"), []rune("lo")},
		{[]rune("telephone12"), []rune("te")},
		{[]rune("51515"), []rune("51")},
		{[]rune("vnvnvn"), []rune("vn")},
		{[]rune("fds2fd"), []rune("fd")},
		{[]rune("`amur"), []rune("`a")},
		{[]rune("soccer0723"), []rune("so")},
		{[]rune("paroach"), []rune("pa")},
		{[]rune("Doritos"), []rune("Do")},
		{[]rune("skiper"), []rune("sk")},
		{[]rune("thx1man"), []rune("th")},
		{[]rune("chula4"), []rune("ch")},
		{[]rune("1trebla"), []rune("1t")},
		{[]rune("nancy21"), []rune("na")},
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
		{[]rune("own"), []rune("n")},
		{[]rune("sp33d"), []rune("33d")},
		{[]rune("roearoea"), []rune("earoea")},
		{[]rune("5556662"), []rune("56662")},
		{[]rune("yrallih"), []rune("allih")},
		{[]rune("gemma8"), []rune("mma8")},
		{[]rune("love4ever21"), []rune("ve4ever21")},
		{[]rune("telephone12"), []rune("lephone12")},
		{[]rune("51515"), []rune("515")},
		{[]rune("vnvnvn"), []rune("vnvn")},
		{[]rune("fds2fd"), []rune("s2fd")},
		{[]rune("`amur"), []rune("mur")},
		{[]rune("soccer0723"), []rune("ccer0723")},
		{[]rune("paroach"), []rune("roach")},
		{[]rune("Doritos"), []rune("ritos")},
		{[]rune("skiper"), []rune("iper")},
		{[]rune("thx1man"), []rune("x1man")},
		{[]rune("chula4"), []rune("ula4")},
		{[]rune("1trebla"), []rune("rebla")},
		{[]rune("nancy21"), []rune("ncy21")},
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
		{[]rune("own"), []rune("own")},
		{[]rune("sp33d"), []rune("sp33!d")},
		{[]rune("roearoea"), []rune("roea!roea")},
		{[]rune("5556662"), []rune("5556!662")},
		{[]rune("yrallih"), []rune("yral!lih")},
		{[]rune("gemma8"), []rune("gemm!a8")},
		{[]rune("love4ever21"), []rune("love!4ever21")},
		{[]rune("telephone12"), []rune("tele!phone12")},
		{[]rune("51515"), []rune("5151!5")},
		{[]rune("vnvnvn"), []rune("vnvn!vn")},
		{[]rune("fds2fd"), []rune("fds2!fd")},
		{[]rune("`amur"), []rune("`amu!r")},
		{[]rune("soccer0723"), []rune("socc!er0723")},
		{[]rune("paroach"), []rune("paro!ach")},
		{[]rune("Doritos"), []rune("Dori!tos")},
		{[]rune("skiper"), []rune("skip!er")},
		{[]rune("thx1man"), []rune("thx1!man")},
		{[]rune("chula4"), []rune("chul!a4")},
		{[]rune("1trebla"), []rune("1tre!bla")},
		{[]rune("nancy21"), []rune("nanc!y21")},
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
		{[]rune("own"), []rune("own")},
		{[]rune("sp33d"), []rune("sp3$d")},
		{[]rune("roearoea"), []rune("roe$roea")},
		{[]rune("5556662"), []rune("555$662")},
		{[]rune("yrallih"), []rune("yra$lih")},
		{[]rune("gemma8"), []rune("gem$a8")},
		{[]rune("love4ever21"), []rune("lov$4ever21")},
		{[]rune("telephone12"), []rune("tel$phone12")},
		{[]rune("51515"), []rune("515$5")},
		{[]rune("vnvnvn"), []rune("vnv$vn")},
		{[]rune("fds2fd"), []rune("fds$fd")},
		{[]rune("`amur"), []rune("`am$r")},
		{[]rune("soccer0723"), []rune("soc$er0723")},
		{[]rune("paroach"), []rune("par$ach")},
		{[]rune("Doritos"), []rune("Dor$tos")},
		{[]rune("skiper"), []rune("ski$er")},
		{[]rune("thx1man"), []rune("thx$man")},
		{[]rune("chula4"), []rune("chu$a4")},
		{[]rune("1trebla"), []rune("1tr$bla")},
		{[]rune("nancy21"), []rune("nan$y21")},
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
		{[]rune("own"), []rune("own")},
		{[]rune("sp33d"), []rune("sp33d")},
		{[]rune("roearoea"), []rune("roearo")},
		{[]rune("5556662"), []rune("555666")},
		{[]rune("yrallih"), []rune("yralli")},
		{[]rune("gemma8"), []rune("gemma8")},
		{[]rune("love4ever21"), []rune("love4e")},
		{[]rune("telephone12"), []rune("teleph")},
		{[]rune("51515"), []rune("51515")},
		{[]rune("vnvnvn"), []rune("vnvnvn")},
		{[]rune("fds2fd"), []rune("fds2fd")},
		{[]rune("`amur"), []rune("`amur")},
		{[]rune("soccer0723"), []rune("soccer")},
		{[]rune("paroach"), []rune("paroac")},
		{[]rune("Doritos"), []rune("Dorito")},
		{[]rune("skiper"), []rune("skiper")},
		{[]rune("thx1man"), []rune("thx1ma")},
		{[]rune("chula4"), []rune("chula4")},
		{[]rune("1trebla"), []rune("1trebl")},
		{[]rune("nancy21"), []rune("nancy2")},
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
		{[]rune("own"), []rune("own")},
		{[]rune("sp33d"), []rune("$p33d")},
		{[]rune("roearoea"), []rune("roearoea")},
		{[]rune("5556662"), []rune("5556662")},
		{[]rune("yrallih"), []rune("yrallih")},
		{[]rune("gemma8"), []rune("gemma8")},
		{[]rune("love4ever21"), []rune("love4ever21")},
		{[]rune("telephone12"), []rune("telephone12")},
		{[]rune("51515"), []rune("51515")},
		{[]rune("vnvnvn"), []rune("vnvnvn")},
		{[]rune("fds2fd"), []rune("fd$2fd")},
		{[]rune("`amur"), []rune("`amur")},
		{[]rune("soccer0723"), []rune("$occer0723")},
		{[]rune("paroach"), []rune("paroach")},
		{[]rune("Doritos"), []rune("Dorito$")},
		{[]rune("skiper"), []rune("$kiper")},
		{[]rune("thx1man"), []rune("thx1man")},
		{[]rune("chula4"), []rune("chula4")},
		{[]rune("1trebla"), []rune("1trebla")},
		{[]rune("nancy21"), []rune("nancy21")},
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
		{[]rune("own"), []rune("own")},
		{[]rune("sp33d"), []rune("p33d")},
		{[]rune("roearoea"), []rune("roearoea")},
		{[]rune("5556662"), []rune("5556662")},
		{[]rune("yrallih"), []rune("yrallih")},
		{[]rune("gemma8"), []rune("gemma8")},
		{[]rune("love4ever21"), []rune("love4ever21")},
		{[]rune("telephone12"), []rune("telephone12")},
		{[]rune("51515"), []rune("51515")},
		{[]rune("vnvnvn"), []rune("vnvnvn")},
		{[]rune("fds2fd"), []rune("fd2fd")},
		{[]rune("`amur"), []rune("`amur")},
		{[]rune("soccer0723"), []rune("occer0723")},
		{[]rune("paroach"), []rune("paroach")},
		{[]rune("Doritos"), []rune("Dorito")},
		{[]rune("skiper"), []rune("kiper")},
		{[]rune("thx1man"), []rune("thx1man")},
		{[]rune("chula4"), []rune("chula4")},
		{[]rune("1trebla"), []rune("1trebla")},
		{[]rune("nancy21"), []rune("nancy21")},
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
		{[]rune("own"), []rune("ooown")},
		{[]rune("sp33d"), []rune("sssp33d")},
		{[]rune("roearoea"), []rune("rrroearoea")},
		{[]rune("5556662"), []rune("555556662")},
		{[]rune("yrallih"), []rune("yyyrallih")},
		{[]rune("gemma8"), []rune("gggemma8")},
		{[]rune("love4ever21"), []rune("lllove4ever21")},
		{[]rune("telephone12"), []rune("tttelephone12")},
		{[]rune("51515"), []rune("5551515")},
		{[]rune("vnvnvn"), []rune("vvvnvnvn")},
		{[]rune("fds2fd"), []rune("fffds2fd")},
		{[]rune("`amur"), []rune("```amur")},
		{[]rune("soccer0723"), []rune("sssoccer0723")},
		{[]rune("paroach"), []rune("ppparoach")},
		{[]rune("Doritos"), []rune("DDDoritos")},
		{[]rune("skiper"), []rune("ssskiper")},
		{[]rune("thx1man"), []rune("ttthx1man")},
		{[]rune("chula4"), []rune("ccchula4")},
		{[]rune("1trebla"), []rune("111trebla")},
		{[]rune("nancy21"), []rune("nnnancy21")},
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
		{[]rune("own"), []rune("ownnn")},
		{[]rune("sp33d"), []rune("sp33ddd")},
		{[]rune("roearoea"), []rune("roearoeaaa")},
		{[]rune("5556662"), []rune("555666222")},
		{[]rune("yrallih"), []rune("yrallihhh")},
		{[]rune("gemma8"), []rune("gemma888")},
		{[]rune("love4ever21"), []rune("love4ever2111")},
		{[]rune("telephone12"), []rune("telephone1222")},
		{[]rune("51515"), []rune("5151555")},
		{[]rune("vnvnvn"), []rune("vnvnvnnn")},
		{[]rune("fds2fd"), []rune("fds2fddd")},
		{[]rune("`amur"), []rune("`amurrr")},
		{[]rune("soccer0723"), []rune("soccer072333")},
		{[]rune("paroach"), []rune("paroachhh")},
		{[]rune("Doritos"), []rune("Doritosss")},
		{[]rune("skiper"), []rune("skiperrr")},
		{[]rune("thx1man"), []rune("thx1mannn")},
		{[]rune("chula4"), []rune("chula444")},
		{[]rune("1trebla"), []rune("1treblaaa")},
		{[]rune("nancy21"), []rune("nancy2111")},
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
		{[]rune("own"), []rune("oowwnn")},
		{[]rune("sp33d"), []rune("sspp3333dd")},
		{[]rune("roearoea"), []rune("rrooeeaarrooeeaa")},
		{[]rune("5556662"), []rune("55555566666622")},
		{[]rune("yrallih"), []rune("yyrraalllliihh")},
		{[]rune("gemma8"), []rune("ggeemmmmaa88")},
		{[]rune("love4ever21"), []rune("lloovvee44eevveerr2211")},
		{[]rune("telephone12"), []rune("tteelleepphhoonnee1122")},
		{[]rune("51515"), []rune("5511551155")},
		{[]rune("vnvnvn"), []rune("vvnnvvnnvvnn")},
		{[]rune("fds2fd"), []rune("ffddss22ffdd")},
		{[]rune("`amur"), []rune("``aammuurr")},
		{[]rune("soccer0723"), []rune("ssoocccceerr00772233")},
		{[]rune("paroach"), []rune("ppaarrooaacchh")},
		{[]rune("Doritos"), []rune("DDoorriittooss")},
		{[]rune("skiper"), []rune("sskkiippeerr")},
		{[]rune("thx1man"), []rune("tthhxx11mmaann")},
		{[]rune("chula4"), []rune("cchhuullaa44")},
		{[]rune("1trebla"), []rune("11ttrreebbllaa")},
		{[]rune("nancy21"), []rune("nnaannccyy2211")},
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
		{[]rune("own"), []rune{}},
		{[]rune("sp33d"), []rune{}},
		{[]rune("roearoea"), []rune("roearoeroa")},
		{[]rune("5556662"), []rune("555666266")},
		{[]rune("yrallih"), []rune("yrallihli")},
		{[]rune("gemma8"), []rune{}},
		{[]rune("love4ever21"), []rune("love4ev4eer21")},
		{[]rune("telephone12"), []rune("telephophne12")},
		{[]rune("51515"), []rune{}},
		{[]rune("vnvnvn"), []rune{}},
		{[]rune("fds2fd"), []rune{}},
		{[]rune("`amur"), []rune{}},
		{[]rune("soccer0723"), []rune("soccer0er723")},
		{[]rune("paroach"), []rune("paroachac")},
		{[]rune("Doritos"), []rune("Doritosto")},
		{[]rune("skiper"), []rune{}},
		{[]rune("thx1man"), []rune("thx1manma")},
		{[]rune("chula4"), []rune{}},
		{[]rune("1trebla"), []rune("1treblabl")},
		{[]rune("nancy21"), []rune("nancy21y2")},
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
		{[]rune("own"), []rune("ownown")},
		{[]rune("sp33d"), []rune("sp33dsp33d")},
		{[]rune("roearoea"), []rune("roearoearoearoea")},
		{[]rune("5556662"), []rune("55566625556662")},
		{[]rune("yrallih"), []rune("yrallihyrallih")},
		{[]rune("gemma8"), []rune("gemma8gemma8")},
		{[]rune("love4ever21"), []rune("love4ever21love4ever21")},
		{[]rune("telephone12"), []rune("telephone12telephone12")},
		{[]rune("51515"), []rune("5151551515")},
		{[]rune("vnvnvn"), []rune("vnvnvnvnvnvn")},
		{[]rune("fds2fd"), []rune("fds2fdfds2fd")},
		{[]rune("`amur"), []rune("`amur`amur")},
		{[]rune("soccer0723"), []rune("soccer0723soccer0723")},
		{[]rune("paroach"), []rune("paroachparoach")},
		{[]rune("Doritos"), []rune("DoritosDoritos")},
		{[]rune("skiper"), []rune("skiperskiper")},
		{[]rune("thx1man"), []rune("thx1manthx1man")},
		{[]rune("chula4"), []rune("chula4chula4")},
		{[]rune("1trebla"), []rune("1trebla1trebla")},
		{[]rune("nancy21"), []rune("nancy21nancy21")},
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
		{[]rune("own"), []rune("ownown")},
		{[]rune("sp33d"), []rune("sp33dsp33d")},
		{[]rune("roearoea"), []rune("roearoearoearoea")},
		{[]rune("5556662"), []rune("55566625556662")},
		{[]rune("yrallih"), []rune("yrallihyrallih")},
		{[]rune("gemma8"), []rune("gemma8gemma8")},
		{[]rune("love4ever21"), []rune("love4ever21love4ever21")},
		{[]rune("telephone12"), []rune("telephone12telephone12")},
		{[]rune("51515"), []rune("5151551515")},
		{[]rune("vnvnvn"), []rune("vnvnvnvnvnvn")},
		{[]rune("fds2fd"), []rune("fds2fdfds2fd")},
		{[]rune("`amur"), []rune("`amur`amur")},
		{[]rune("soccer0723"), []rune("soccer0723soccer0723")},
		{[]rune("paroach"), []rune("paroachparoach")},
		{[]rune("Doritos"), []rune("DoritosDoritos")},
		{[]rune("skiper"), []rune("skiperskiper")},
		{[]rune("thx1man"), []rune("thx1manthx1man")},
		{[]rune("chula4"), []rune("chula4chula4")},
		{[]rune("1trebla"), []rune("1trebla1trebla")},
		{[]rune("nancy21"), []rune("nancy21nancy21")},
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
		{[]rune("own"), []rune("own")},
		{[]rune("sp33d"), []rune("sp33d")},
		{[]rune("roearoea"), []rune("roearoea")},
		{[]rune("5556662"), []rune("5556662")},
		{[]rune("yrallih"), []rune("yrallih")},
		{[]rune("gemma8"), []rune("gemma8")},
		{[]rune("love4ever21"), []rune("love4ever21")},
		{[]rune("telephone12"), []rune("telephone12")},
		{[]rune("51515"), []rune("51515")},
		{[]rune("vnvnvn"), []rune("vnvnvn")},
		{[]rune("fds2fd"), []rune("fds2fd")},
		{[]rune("`amur"), []rune("`amur")},
		{[]rune("soccer0723"), []rune("soccer0723")},
		{[]rune("paroach"), []rune("paroach")},
		{[]rune("Doritos"), []rune("Doritos")},
		{[]rune("skiper"), []rune("skiper")},
		{[]rune("thx1man"), []rune("thx1man")},
		{[]rune("chula4"), []rune("chula4")},
		{[]rune("1trebla"), []rune("1trebla")},
		{[]rune("nancy21"), []rune("nancy21")},
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
		{[]rune("own"), false},
		{[]rune("sp33d"), false},
		{[]rune("roearoea"), true},
		{[]rune("5556662"), false},
		{[]rune("yrallih"), false},
		{[]rune("gemma8"), false},
		{[]rune("love4ever21"), true},
		{[]rune("telephone12"), true},
		{[]rune("51515"), false},
		{[]rune("vnvnvn"), false},
		{[]rune("fds2fd"), false},
		{[]rune("`amur"), false},
		{[]rune("soccer0723"), true},
		{[]rune("paroach"), false},
		{[]rune("Doritos"), false},
		{[]rune("skiper"), false},
		{[]rune("thx1man"), false},
		{[]rune("chula4"), false},
		{[]rune("1trebla"), false},
		{[]rune("nancy21"), false},
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
		{[]rune("own"), true},
		{[]rune("sp33d"), true},
		{[]rune("roearoea"), false},
		{[]rune("5556662"), false},
		{[]rune("yrallih"), false},
		{[]rune("gemma8"), true},
		{[]rune("love4ever21"), false},
		{[]rune("telephone12"), false},
		{[]rune("51515"), true},
		{[]rune("vnvnvn"), true},
		{[]rune("fds2fd"), true},
		{[]rune("`amur"), true},
		{[]rune("soccer0723"), false},
		{[]rune("paroach"), false},
		{[]rune("Doritos"), false},
		{[]rune("skiper"), true},
		{[]rune("thx1man"), false},
		{[]rune("chula4"), true},
		{[]rune("1trebla"), false},
		{[]rune("nancy21"), false},
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
		{[]rune("own"), false},
		{[]rune("sp33d"), false},
		{[]rune("roearoea"), true},
		{[]rune("5556662"), false},
		{[]rune("yrallih"), false},
		{[]rune("gemma8"), true},
		{[]rune("love4ever21"), true},
		{[]rune("telephone12"), true},
		{[]rune("51515"), false},
		{[]rune("vnvnvn"), false},
		{[]rune("fds2fd"), false},
		{[]rune("`amur"), false},
		{[]rune("soccer0723"), true},
		{[]rune("paroach"), false},
		{[]rune("Doritos"), false},
		{[]rune("skiper"), true},
		{[]rune("thx1man"), false},
		{[]rune("chula4"), false},
		{[]rune("1trebla"), true},
		{[]rune("nancy21"), false},
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
		{[]rune("own"), true},
		{[]rune("sp33d"), true},
		{[]rune("roearoea"), false},
		{[]rune("5556662"), true},
		{[]rune("yrallih"), true},
		{[]rune("gemma8"), false},
		{[]rune("love4ever21"), false},
		{[]rune("telephone12"), false},
		{[]rune("51515"), true},
		{[]rune("vnvnvn"), true},
		{[]rune("fds2fd"), true},
		{[]rune("`amur"), true},
		{[]rune("soccer0723"), false},
		{[]rune("paroach"), true},
		{[]rune("Doritos"), true},
		{[]rune("skiper"), false},
		{[]rune("thx1man"), true},
		{[]rune("chula4"), true},
		{[]rune("1trebla"), false},
		{[]rune("nancy21"), true},
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
		{[]rune("own"), true},
		{[]rune("sp33d"), false},
		{[]rune("roearoea"), true},
		{[]rune("5556662"), true},
		{[]rune("yrallih"), true},
		{[]rune("gemma8"), true},
		{[]rune("love4ever21"), true},
		{[]rune("telephone12"), true},
		{[]rune("51515"), true},
		{[]rune("vnvnvn"), true},
		{[]rune("fds2fd"), true},
		{[]rune("`amur"), true},
		{[]rune("soccer0723"), false},
		{[]rune("paroach"), true},
		{[]rune("Doritos"), true},
		{[]rune("skiper"), false},
		{[]rune("thx1man"), true},
		{[]rune("chula4"), true},
		{[]rune("1trebla"), true},
		{[]rune("nancy21"), true},
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
		{[]rune("own"), false},
		{[]rune("sp33d"), true},
		{[]rune("roearoea"), true},
		{[]rune("5556662"), true},
		{[]rune("yrallih"), true},
		{[]rune("gemma8"), true},
		{[]rune("love4ever21"), true},
		{[]rune("telephone12"), true},
		{[]rune("51515"), true},
		{[]rune("vnvnvn"), false},
		{[]rune("fds2fd"), true},
		{[]rune("`amur"), true},
		{[]rune("soccer0723"), true},
		{[]rune("paroach"), true},
		{[]rune("Doritos"), true},
		{[]rune("skiper"), true},
		{[]rune("thx1man"), false},
		{[]rune("chula4"), true},
		{[]rune("1trebla"), true},
		{[]rune("nancy21"), true},
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
		{[]rune("own"), true},
		{[]rune("sp33d"), true},
		{[]rune("roearoea"), false},
		{[]rune("5556662"), true},
		{[]rune("yrallih"), true},
		{[]rune("gemma8"), true},
		{[]rune("love4ever21"), false},
		{[]rune("telephone12"), true},
		{[]rune("51515"), true},
		{[]rune("vnvnvn"), true},
		{[]rune("fds2fd"), true},
		{[]rune("`amur"), true},
		{[]rune("soccer0723"), false},
		{[]rune("paroach"), true},
		{[]rune("Doritos"), false},
		{[]rune("skiper"), true},
		{[]rune("thx1man"), true},
		{[]rune("chula4"), true},
		{[]rune("1trebla"), true},
		{[]rune("nancy21"), true},
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
		{[]rune("own"), true},
		{[]rune("sp33d"), true},
		{[]rune("roearoea"), false},
		{[]rune("5556662"), true},
		{[]rune("yrallih"), false},
		{[]rune("gemma8"), false},
		{[]rune("love4ever21"), true},
		{[]rune("telephone12"), true},
		{[]rune("51515"), true},
		{[]rune("vnvnvn"), true},
		{[]rune("fds2fd"), true},
		{[]rune("`amur"), false},
		{[]rune("soccer0723"), true},
		{[]rune("paroach"), false},
		{[]rune("Doritos"), true},
		{[]rune("skiper"), true},
		{[]rune("thx1man"), false},
		{[]rune("chula4"), false},
		{[]rune("1trebla"), false},
		{[]rune("nancy21"), false},
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
		{[]rune("own"), false},
		{[]rune("sp33d"), false},
		{[]rune("roearoea"), false},
		{[]rune("5556662"), false},
		{[]rune("yrallih"), false},
		{[]rune("gemma8"), true},
		{[]rune("love4ever21"), false},
		{[]rune("telephone12"), false},
		{[]rune("51515"), false},
		{[]rune("vnvnvn"), false},
		{[]rune("fds2fd"), false},
		{[]rune("`amur"), false},
		{[]rune("soccer0723"), false},
		{[]rune("paroach"), false},
		{[]rune("Doritos"), false},
		{[]rune("skiper"), false},
		{[]rune("thx1man"), false},
		{[]rune("chula4"), false},
		{[]rune("1trebla"), false},
		{[]rune("nancy21"), false},
	}

	Memorize([]rune("gemma8"))

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
		{[]rune("own"), []rune("won")},
		{[]rune("sp33d"), []rune("ps33d")},
		{[]rune("roearoea"), []rune("orearoea")},
		{[]rune("5556662"), []rune("5556662")},
		{[]rune("yrallih"), []rune("ryallih")},
		{[]rune("gemma8"), []rune("egmma8")},
		{[]rune("love4ever21"), []rune("olve4ever21")},
		{[]rune("telephone12"), []rune("etlephone12")},
		{[]rune("51515"), []rune("15515")},
		{[]rune("vnvnvn"), []rune("nvvnvn")},
		{[]rune("fds2fd"), []rune("dfs2fd")},
		{[]rune("`amur"), []rune("a`mur")},
		{[]rune("soccer0723"), []rune("osccer0723")},
		{[]rune("paroach"), []rune("aproach")},
		{[]rune("Doritos"), []rune("oDritos")},
		{[]rune("skiper"), []rune("ksiper")},
		{[]rune("thx1man"), []rune("htx1man")},
		{[]rune("chula4"), []rune("hcula4")},
		{[]rune("1trebla"), []rune("t1rebla")},
		{[]rune("nancy21"), []rune("anncy21")},
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
		{[]rune("own"), []rune("onw")},
		{[]rune("sp33d"), []rune("sp3d3")},
		{[]rune("roearoea"), []rune("roearoae")},
		{[]rune("5556662"), []rune("5556626")},
		{[]rune("yrallih"), []rune("yrallhi")},
		{[]rune("gemma8"), []rune("gemm8a")},
		{[]rune("love4ever21"), []rune("love4ever12")},
		{[]rune("telephone12"), []rune("telephone21")},
		{[]rune("51515"), []rune("51551")},
		{[]rune("vnvnvn"), []rune("vnvnnv")},
		{[]rune("fds2fd"), []rune("fds2df")},
		{[]rune("`amur"), []rune("`amru")},
		{[]rune("soccer0723"), []rune("soccer0732")},
		{[]rune("paroach"), []rune("paroahc")},
		{[]rune("Doritos"), []rune("Doritso")},
		{[]rune("skiper"), []rune("skipre")},
		{[]rune("thx1man"), []rune("thx1mna")},
		{[]rune("chula4"), []rune("chul4a")},
		{[]rune("1trebla"), []rune("1trebal")},
		{[]rune("nancy21"), []rune("nancy12")},
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
		{[]rune("own"), []rune("own")},
		{[]rune("sp33d"), []rune("sp3d3")},
		{[]rune("roearoea"), []rune("roeraoea")},
		{[]rune("5556662"), []rune("5556662")},
		{[]rune("yrallih"), []rune("yrallih")},
		{[]rune("gemma8"), []rune("gemam8")},
		{[]rune("love4ever21"), []rune("lov4eever21")},
		{[]rune("telephone12"), []rune("telpehone12")},
		{[]rune("51515"), []rune("51551")},
		{[]rune("vnvnvn"), []rune("vnvvnn")},
		{[]rune("fds2fd"), []rune("fdsf2d")},
		{[]rune("`amur"), []rune("`amru")},
		{[]rune("soccer0723"), []rune("socecr0723")},
		{[]rune("paroach"), []rune("paraoch")},
		{[]rune("Doritos"), []rune("Dortios")},
		{[]rune("skiper"), []rune("skiepr")},
		{[]rune("thx1man"), []rune("thxm1an")},
		{[]rune("chula4"), []rune("chual4")},
		{[]rune("1trebla"), []rune("1trbela")},
		{[]rune("nancy21"), []rune("nanyc21")},
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
		{[]rune("sp33d"), []rune("spf3d")},
		{[]rune("roearoea"), []rune("roÊaroea")},
		{[]rune("5556662"), []rune("55j6662")},
		{[]rune("yrallih"), []rune("yrÂllih")},
		{[]rune("gemma8"), []rune("geÚma8")},
		{[]rune("love4ever21"), []rune("loìe4ever21")},
		{[]rune("telephone12"), []rune("teØephone12")},
		{[]rune("51515"), []rune("51j15")},
		{[]rune("vnvnvn"), []rune("vnìnvn")},
		{[]rune("fds2fd"), []rune("fdæ2fd")},
		{[]rune("`amur"), []rune("`aÚur")},
		{[]rune("soccer0723"), []rune("soÆcer0723")},
		{[]rune("paroach"), []rune("paäoach")},
		{[]rune("Doritos"), []rune("Doäitos")},
		{[]rune("skiper"), []rune("skÒper")},
		{[]rune("thx1man"), []rune("thð1man")},
		{[]rune("chula4"), []rune("chêla4")},
		{[]rune("1trebla"), []rune("1täebla")},
		{[]rune("nancy21"), []rune("naÜcy21")},
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
		{[]rune("own"), []rune("ow7")},
		{[]rune("sp33d"), []rune("sp3d")},
		{[]rune("roearoea"), []rune("ro2aroea")},
		{[]rune("5556662"), []rune("556662")},
		{[]rune("yrallih"), []rune("yr0llih")},
		{[]rune("gemma8"), []rune("ge6ma8")},
		{[]rune("love4ever21"), []rune("lo;e4ever21")},
		{[]rune("telephone12"), []rune("te6ephone12")},
		{[]rune("51515"), []rune("5115")},
		{[]rune("vnvnvn"), []rune("vn;nvn")},
		{[]rune("fds2fd"), []rune("fd92fd")},
		{[]rune("`amur"), []rune("`a6ur")},
		{[]rune("soccer0723"), []rune("so1cer0723")},
		{[]rune("paroach"), []rune("pa9oach")},
		{[]rune("Doritos"), []rune("Do9itos")},
		{[]rune("skiper"), []rune("sk4per")},
		{[]rune("thx1man"), []rune("th<1man")},
		{[]rune("own"), []rune("ow7")},
		{[]rune("chula4"), []rune("ch:la4")},
		{[]rune("1trebla"), []rune("1t9ebla")},
		{[]rune("nancy21"), []rune("na7cy21")},
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
		{[]rune("own"), []rune("owo")},
		{[]rune("sp33d"), []rune("sp43d")},
		{[]rune("roearoea"), []rune("rofaroea")},
		{[]rune("5556662"), []rune("5566662")},
		{[]rune("yrallih"), []rune("yrbllih")},
		{[]rune("gemma8"), []rune("genma8")},
		{[]rune("love4ever21"), []rune("lowe4ever21")},
		{[]rune("telephone12"), []rune("temephone12")},
		{[]rune("51515"), []rune("51615")},
		{[]rune("vnvnvn"), []rune("vnwnvn")},
		{[]rune("fds2fd"), []rune("fdt2fd")},
		{[]rune("`amur"), []rune("`anur")},
		{[]rune("soccer0723"), []rune("sodcer0723")},
		{[]rune("paroach"), []rune("pasoach")},
		{[]rune("Doritos"), []rune("Dositos")},
		{[]rune("skiper"), []rune("skjper")},
		{[]rune("thx1man"), []rune("thy1man")},
		{[]rune("chula4"), []rune("chvla4")},
		{[]rune("1trebla"), []rune("1tsebla")},
		{[]rune("nancy21"), []rune("naocy21")},
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
		{[]rune("own"), []rune("ovn")},
		{[]rune("sp33d"), []rune("so33d")},
		{[]rune("roearoea"), []rune("rnearoea")},
		{[]rune("5556662"), []rune("5456662")},
		{[]rune("yrallih"), []rune("yqallih")},
		{[]rune("gemma8"), []rune("gdmma8")},
		{[]rune("love4ever21"), []rune("lnve4ever21")},
		{[]rune("telephone12"), []rune("tdlephone12")},
		{[]rune("51515"), []rune("50515")},
		{[]rune("vnvnvn"), []rune("vmvnvn")},
		{[]rune("fds2fd"), []rune("fcs2fd")},
		{[]rune("`amur"), []rune("``mur")},
		{[]rune("soccer0723"), []rune("snccer0723")},
		{[]rune("paroach"), []rune("p`roach")},
		{[]rune("Doritos"), []rune("Dnritos")},
		{[]rune("skiper"), []rune("sjiper")},
		{[]rune("thx1man"), []rune("tgx1man")},
		{[]rune("chula4"), []rune("cgula4")},
		{[]rune("1trebla"), []rune("1srebla")},
		{[]rune("nancy21"), []rune("n`ncy21")},
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
		{[]rune("own"), []rune("onn")},
		{[]rune("sp33d"), []rune("s333d")},
		{[]rune("roearoea"), []rune("reearoea")},
		{[]rune("5556662"), []rune("5556662")},
		{[]rune("yrallih"), []rune("yaallih")},
		{[]rune("gemma8"), []rune("gmmma8")},
		{[]rune("love4ever21"), []rune("lvve4ever21")},
		{[]rune("telephone12"), []rune("tllephone12")},
		{[]rune("51515"), []rune("55515")},
		{[]rune("vnvnvn"), []rune("vvvnvn")},
		{[]rune("fds2fd"), []rune("fss2fd")},
		{[]rune("`amur"), []rune("`mmur")},
		{[]rune("soccer0723"), []rune("scccer0723")},
		{[]rune("paroach"), []rune("prroach")},
		{[]rune("Doritos"), []rune("Drritos")},
		{[]rune("skiper"), []rune("siiper")},
		{[]rune("thx1man"), []rune("txx1man")},
		{[]rune("chula4"), []rune("cuula4")},
		{[]rune("1trebla"), []rune("1rrebla")},
		{[]rune("nancy21"), []rune("nnncy21")},
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
		{[]rune("own"), []rune("oon")},
		{[]rune("sp33d"), []rune("ss33d")},
		{[]rune("roearoea"), []rune("rrearoea")},
		{[]rune("5556662"), []rune("5556662")},
		{[]rune("yrallih"), []rune("yyallih")},
		{[]rune("gemma8"), []rune("ggmma8")},
		{[]rune("love4ever21"), []rune("llve4ever21")},
		{[]rune("telephone12"), []rune("ttlephone12")},
		{[]rune("51515"), []rune("55515")},
		{[]rune("vnvnvn"), []rune("vvvnvn")},
		{[]rune("fds2fd"), []rune("ffs2fd")},
		{[]rune("`amur"), []rune("``mur")},
		{[]rune("soccer0723"), []rune("ssccer0723")},
		{[]rune("paroach"), []rune("pproach")},
		{[]rune("Doritos"), []rune("DDritos")},
		{[]rune("skiper"), []rune("ssiper")},
		{[]rune("thx1man"), []rune("ttx1man")},
		{[]rune("chula4"), []rune("ccula4")},
		{[]rune("1trebla"), []rune("11rebla")},
		{[]rune("nancy21"), []rune("nnncy21")},
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
		{[]rune("own"), []rune("owown")},
		{[]rune("sp33d"), []rune("spsp33d")},
		{[]rune("roearoea"), []rune("roroearoea")},
		{[]rune("5556662"), []rune("555556662")},
		{[]rune("yrallih"), []rune("yryrallih")},
		{[]rune("gemma8"), []rune("gegemma8")},
		{[]rune("love4ever21"), []rune("lolove4ever21")},
		{[]rune("telephone12"), []rune("tetelephone12")},
		{[]rune("51515"), []rune("5151515")},
		{[]rune("vnvnvn"), []rune("vnvnvnvn")},
		{[]rune("fds2fd"), []rune("fdfds2fd")},
		{[]rune("`amur"), []rune("`a`amur")},
		{[]rune("soccer0723"), []rune("sosoccer0723")},
		{[]rune("paroach"), []rune("paparoach")},
		{[]rune("Doritos"), []rune("DoDoritos")},
		{[]rune("skiper"), []rune("skskiper")},
		{[]rune("thx1man"), []rune("ththx1man")},
		{[]rune("chula4"), []rune("chchula4")},
		{[]rune("1trebla"), []rune("1t1trebla")},
		{[]rune("nancy21"), []rune("nanancy21")},
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
		{[]rune("own"), []rune("ownwn")},
		{[]rune("sp33d"), []rune("sp33d3d")},
		{[]rune("roearoea"), []rune("roearoeaea")},
		{[]rune("5556662"), []rune("555666262")},
		{[]rune("yrallih"), []rune("yrallihih")},
		{[]rune("gemma8"), []rune("gemma8a8")},
		{[]rune("love4ever21"), []rune("love4ever2121")},
		{[]rune("telephone12"), []rune("telephone1212")},
		{[]rune("51515"), []rune("5151515")},
		{[]rune("vnvnvn"), []rune("vnvnvnvn")},
		{[]rune("fds2fd"), []rune("fds2fdfd")},
		{[]rune("`amur"), []rune("`amurur")},
		{[]rune("soccer0723"), []rune("soccer072323")},
		{[]rune("paroach"), []rune("paroachch")},
		{[]rune("Doritos"), []rune("Doritosos")},
		{[]rune("skiper"), []rune("skiperer")},
		{[]rune("thx1man"), []rune("thx1manan")},
		{[]rune("chula4"), []rune("chula4a4")},
		{[]rune("1trebla"), []rune("1treblala")},
		{[]rune("nancy21"), []rune("nancy2121")},
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
		{[]rune("own"), []rune("Own")},
		{[]rune("sp33d"), []rune("Sp33d")},
		{[]rune("roearoea"), []rune("Roearoea")},
		{[]rune("5556662"), []rune("5556662")},
		{[]rune("yrallih"), []rune("Yrallih")},
		{[]rune("gemma8"), []rune("Gemma8")},
		{[]rune("love4ever21"), []rune("Love4ever21")},
		{[]rune("telephone12"), []rune("Telephone12")},
		{[]rune("51515"), []rune("51515")},
		{[]rune("vnvnvn"), []rune("Vnvnvn")},
		{[]rune("fds2fd"), []rune("Fds2fd")},
		{[]rune("`amur"), []rune("`amur")},
		{[]rune("soccer0723"), []rune("Soccer0723")},
		{[]rune("paroach"), []rune("Paroach")},
		{[]rune("Doritos"), []rune("Doritos")},
		{[]rune("skiper"), []rune("Skiper")},
		{[]rune("thx1man"), []rune("Thx1man")},
		{[]rune("chula4"), []rune("Chula4")},
		{[]rune("1trebla"), []rune("1trebla")},
		{[]rune("nancy21"), []rune("Nancy21")},
	}

	for _, tt := range ruletest {
		actual := Title(tt.in)
		if compare(actual, tt.out) {
			t.Errorf("expected %s, got %s", string(tt.out), string(actual))
		}
	}

}

//Saved=[]rune("aavwlo.trw")

var result int

//func ToNum(char uint8) int {
func BenchmarkToNum(b *testing.B) {
	var r int
	// read comment about single word
	// same applies here
	var i uint8 = 12
	for n := 0; n < b.N; n++ {
		r = ToNumByte(i)
	}

	result = r
}

// dont really know how else i should do this
// just going to use this one word for benchmarking
// maybe have a slice of slices to do more realistic benchmarking?
var mangled []rune
var plain = []rune("golang")

//func Nothing(word []rune) []rune {
func BenchmarkNothing(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = Nothing(plain)
	}

	mangled = r
}

//func Lowercase(word []rune) []rune {
func BenchmarkLowercase(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = Lowercase(plain)
	}

	mangled = r
}

//func Uppercase(word []rune) []rune {
func BenchmarkUppercase(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = Uppercase(plain)
	}

	mangled = r
}

//func Capitalize(word []rune) []rune {
func BenchmarkCapitalize(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = Capitalize(plain)
	}

	mangled = r
}

//func InvertCapitalize(word []rune) []rune {
func BenchmarkInvertCapitalize(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = InvertCapitalize(plain)
	}

	mangled = r
}

//func ToggleCase(word []rune) []rune {
func BenchmarkToggleCase(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = ToggleCase(plain)
	}

	mangled = r
}

//func ToggleAt(word []rune, n int) []rune {
func BenchmarkToggleAt(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = ToggleAt(plain, 2)
	}

	mangled = r
}

//func Reverse(word []rune) []rune {
func BenchmarkReverse(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = Reverse(plain)
	}

	mangled = r
}

//func Duplicate(word []rune) []rune {
func BenchmarkDuplicate(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = Duplicate(plain)
	}

	mangled = r
}

//func DuplicateN(word []rune, n int) []rune {
func BenchmarkDuplicateN(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = DuplicateN(plain, 2)
	}

	mangled = r
}

//func Reflect(word []rune) []rune {
func BenchmarkReflect(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = Reflect(plain)
	}

	mangled = r
}

//func RotateLeft(word []rune) []rune {
func BenchmarkRotateLeft(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = RotateLeft(plain)
	}

	mangled = r
}

//func RotateRight(word []rune) []rune {
func BenchmarkRotateRight(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = RotateRight(plain)
	}

	mangled = r
}

//func AppendCharacter(word []rune, char rune) []rune {
func BenchmarkAppendCharacter(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = AppendCharacter(plain, 'a')
	}

	mangled = r
}

//func PrependCharacter(word []rune, char rune) []rune {
func BenchmarkPrependCharacter(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = PrependCharacter(plain, 'a')
	}

	mangled = r
}

//func TruncateLeft(word []rune) []rune {
func BenchmarkTruncateLeft(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = TruncateLeft(plain)
	}

	mangled = r
}

//func TruncateRight(word []rune) []rune {
func BenchmarkTruncateRight(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = TruncateRight(plain)
	}

	mangled = r
}

//func DeleteN(word []rune, n int) []rune {
func BenchmarkDeleteN(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = DeleteN(plain, 2)
	}

	mangled = r
}

//func ExtractRange(word []rune, n, m int) []rune {
func BenchmarkExtractRange(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = ExtractRange(plain, 1, 3)
	}

	mangled = r
}

//func OmitRange(word []rune, n, m int) []rune {
func BenchmarkOmitRange(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = OmitRange(plain, 1, 3)
	}

	mangled = r
}

//func InsertAtN(word []rune, n int, char rune) []rune {
func BenchmarkInsertAtN(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = InsertAtN(plain, 2, 'a')
	}

	mangled = r
}

//func OverwriteAtN(word []rune, n int, char rune) []rune {
func BenchmarkOverwriteAtN(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = OverwriteAtN(plain, 2, 'a')
	}

	mangled = r
}

//func TruncateAtN(word []rune, n int) []rune {
func BenchmarkTruncateAtN(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = TruncateAtN(plain, 2)
	}

	mangled = r
}

//func Replace(word []rune, x, y rune) []rune {
func BenchmarkReplace(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = Replace(plain, 'o', 'a')
	}

	mangled = r
}

//func Purge(word []rune, x rune) []rune {
func BenchmarkPurge(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = Purge(plain, 'o')
	}

	mangled = r
}

//func DuplicateFirstN(word []rune, n int) []rune {
func BenchmarkDuplicateFirstN(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = DuplicateFirstN(plain, 2)
	}

	mangled = r
}

//func DuplicateLastN(word []rune, n int) []rune {
func BenchmarkDuplicateLastN(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = DuplicateLastN(plain, 2)
	}

	mangled = r
}

//func DuplicateAll(word []rune) []rune {
func BenchmarkDuplicateAll(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = DuplicateAll(plain)
	}

	mangled = r
}

//func ExtractMemory(word []rune, n, m, i int) []rune {
func BenchmarkExtractMemory(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = ExtractMemory(plain, 1, 2, 3)
	}

	mangled = r
}

//func AppendMemory(word []rune) []rune {
func BenchmarkAppendMemory(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = AppendMemory(plain)
	}

	mangled = r
}

//func PrependMemory(word []rune) []rune {
func BenchmarkPrependMemory(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = PrependMemory(plain)
	}

	mangled = r
}

//func Memorize(word []rune) {
func BenchmarkMemorize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Memorize(plain)
	}
}

var truth bool

//func RejectLess(word []rune, n int) bool {
func BenchmarkRejectLess(b *testing.B) {
	var r bool

	for n := 0; n < b.N; n++ {
		r = RejectLess(plain, 2)
	}

	truth = r
}

//func RejectGreater(word []rune, n int) bool {
func BenchmarkRejectGreater(b *testing.B) {
	var r bool

	for n := 0; n < b.N; n++ {
		r = RejectGreater(plain, 2)
	}

	truth = r
}

//func RejectContain(word []rune, char rune) bool {
func BenchmarkRejectContain(b *testing.B) {
	var r bool

	for n := 0; n < b.N; n++ {
		r = RejectContain(plain, 'a')
	}

	truth = r
}

//func RejectNotContain(word []rune, char rune) bool {
func BenchmarkRejectNotContain(b *testing.B) {
	var r bool

	for n := 0; n < b.N; n++ {
		r = RejectNotContain(plain, 'a')
	}

	truth = r
}

//func RejectEqualFirst(word []rune, char rune) bool {
func BenchmarkRejectEqualFirst(b *testing.B) {
	var r bool

	for n := 0; n < b.N; n++ {
		r = RejectEqualFirst(plain, 'g')
	}

	truth = r
}

//func RejectEqualLast(word []rune, char rune) bool {
func BenchmarkRejectEqualLast(b *testing.B) {
	var r bool

	for n := 0; n < b.N; n++ {
		r = RejectEqualLast(plain, 'g')
	}

	truth = r
}

//func RejectEqualAt(word []rune, char rune, n int) bool {
func BenchmarkRejectEqualAt(b *testing.B) {
	var r bool

	for n := 0; n < b.N; n++ {
		r = RejectEqualAt(plain, 'g', 0)
	}

	truth = r
}

//func RejectContains(word []rune, char rune, n int) bool {
func BenchmarkRejectContains(b *testing.B) {
	var r bool

	for n := 0; n < b.N; n++ {
		r = RejectContains(plain, 'g', 1)
	}

	truth = r
}

//func RejectMemory(word []rune) bool {
func BenchmarkRejectMemory(b *testing.B) {
	var r bool

	for n := 0; n < b.N; n++ {
		r = RejectMemory(plain)
	}

	truth = r
}

//func SwapFront(word []rune) []rune {
func BenchmarkSwapFront(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = SwapFront(plain)
	}

	mangled = r
}

//func SwapBack(word []rune) []rune {
func BenchmarkSwapBack(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = SwapBack(plain)
	}

	mangled = r
}

//func SwapAtN(word []rune, x, y int) []rune {
func BenchmarkSwapAtN(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = SwapAtN(plain, 2, 0)
	}

	mangled = r
}

//func BitwiseShiftLeft(word []rune, n int) []rune {
func BenchmarkBitwiseShiftLeft(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = BitwiseShiftLeft(plain, 2)
	}

	mangled = r
}

//func BitwiseShiftRight(word []rune, n int) []rune {
func BenchmarkBitwiseShiftRight(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = BitwiseShiftRight(plain, 2)
	}

	mangled = r
}

//func ASCIIIncrementPlus(word []rune, n int) []rune {
func BenchmarkASCIIIncrementPlus(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = ASCIIIncrementPlus(plain, 2)
	}

	mangled = r
}

//func ASCIIIncrementMinus(word []rune, n int) []rune {
func BenchmarkASCIIIncrementMinus(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = ASCIIIncrementMinus(plain, 2)
	}

	mangled = r
}

//func ReplaceNPlus(word []rune, n int) []rune {
func BenchmarkReplaceNPlus(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = ReplaceNPlus(plain, 2)
	}

	mangled = r
}

//func ReplaceNMinus(word []rune, n int) []rune {
func BenchmarkReplaceNMinus(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = ReplaceNMinus(plain, 2)
	}

	mangled = r
}

//func DuplicateBlockFront(word []rune, n int) []rune {
func BenchmarkDuplicateBlockFront(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = DuplicateBlockFront(plain, 2)
	}

	mangled = r
}

//func DuplicateBlockBack(word []rune, n int) []rune {
func BenchmarkDuplicateBlockBack(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = DuplicateBlockBack(plain, 2)
	}

	mangled = r
}

//func Title(word []rune) []rune {
func BenchmarkTitle(b *testing.B) {
	var r []rune

	for n := 0; n < b.N; n++ {
		r = Title(plain)
	}

	mangled = r
}

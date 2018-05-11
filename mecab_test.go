package mecab_test

import (
	"testing"

	mecab "github.com/breezymind/go-mecab"
)

// NOTE: Tests
func Test_Nouns(t *testing.T) {

	m, e := mecab.New()
	if e != nil {
		t.Error(e)
	}
	defer m.Destroy()
	rawtext := "택시 운전사는 어두운 창밖으로 고개를 내밀어 이따금 고함을 친다."
	t.Log(
		m.Nouns(rawtext),
	)
}
func Test_Morphs(t *testing.T) {

	m, e := mecab.New()
	if e != nil {
		t.Error(e)
	}
	defer m.Destroy()
	rawtext := "택시 운전사는 어두운 창밖으로 고개를 내밀어 이따금 고함을 친다."
	t.Log(
		m.Morphs(rawtext),
	)
}
func Test_Pos(t *testing.T) {

	m, e := mecab.New()
	if e != nil {
		t.Error(e)
	}
	defer m.Destroy()
	rawtext := "택시 운전사는 어두운 창밖으로 고개를 내밀어 이따금 고함을 친다."
	t.Log(
		m.Pos(rawtext),
	)
}
func Test_GetJSONPretty(t *testing.T) {

	m, e := mecab.New()
	if e != nil {
		t.Error(e)
	}
	defer m.Destroy()
	rawtext := "택시 운전사는 어두운 창밖으로 고개를 내밀어 이따금 고함을 친다."
	t.Log(
		m.Parse(rawtext, "NN*").GetJSONPretty(),
	)
}

// // NOTE: Examples

// func ExampleRequireJSONFile() {
// 	parse, err := RequireJSONFile("test.json")
// 	if err != nil {
// 		fmt.Printf("Error : %s", err.Error())
// 	}
// 	fmt.Printf(parse.GetJSONPretty())
// 	// Output:
// 	//  {
// 	// 	"obj": {
// 	// 		"arr-attr": [
// 	// 			"gml",
// 	// 			"xml"
// 	// 		],
// 	// 		"int-attr": 100,
// 	// 		"obj-attr": {
// 	// 			"para": "a meta-markup language, used to create markup languages such as docbook."
// 	// 		},
// 	// 		"str-attr": "sgml"
// 	// 	}
// 	// }
// }
// func ExampleGoroutineID() {
// 	fmt.Printf("GoroutineID %d", GoroutineID())
// 	// Output:
// 	// GoroutineID 7
// }

// func ExampleSetTimeout() {
// 	SetTimeout(func() {
// 		fmt.Printf("after 2 seconds")
// 	}, 2000)
// 	// Output:
// 	// after 2 seconds
// }

// func ExampleSetInterval() {
// 	val := 3
// 	SetInterval(func() bool {
// 		fmt.Printf("loop %d sec\n", val)
// 		val--
// 		if val < 1 {
// 			fmt.Print("end")
// 			return true
// 		}
// 		return false
// 	}, 1000)
// 	// Output:
// 	// loop 3 sec
// 	// loop 2 sec
// 	// loop 1 sec
// 	// end
// }

// func ExampleIsJSON() {
// 	parse, err := RequireJSONFile("test.json")
// 	if err != nil {
// 		fmt.Printf("Error : %s", err.Error())
// 	}
// 	rawjson := parse.GetJSONString()
// 	fmt.Print(rawjson)
// 	fmt.Printf("IsJSON %t", IsJSON(rawjson))
// 	// Output:
// 	// {"obj":{"arr-attr":["gml","xml"],"int-attr":100,"obj-attr":{"para":"a meta-markup language, used to create markup languages such as docbook."},"str-attr":"sgml"}}
// 	// IsJSON true
// }

// func ExampleInArray() {
// 	strs := []string{"apple", "banana", "orange"}
// 	fmt.Println(InArray("apple", strs))

// 	ints := []int{1, 2, 3}
// 	fmt.Println(InArray(2, ints))

// 	infs := []interface{}{222, "breezy"}
// 	fmt.Println(InArray("breezy", infs))

// 	// Output:
// 	// true 0
// 	// true 1
// 	// true 1
// }

// func ExampleStringSplitApply() {
// 	strs := "박기호 기자\n     이후민 기자\n     박응진 기자\n     정상훈 기자"
// 	fmt.Println(
// 		StringSplitApply(strs, "\n", "/", func(part string) string {
// 			return strings.TrimSpace(strings.Replace(part, " 기자", "", -1))
// 		}),
// 	)

// 	// Output:
// 	// 박기호/이후민/박응진/정상훈
// }

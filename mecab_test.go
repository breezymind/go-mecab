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
func Test_Morphs2Str(t *testing.T) {
	m, e := mecab.New()
	if e != nil {
		t.Error(e)
	}
	defer m.Destroy()
	rawtext := "택시 운전사는 어두운 창밖으로 고개를 내밀어 이따금 고함을 친다."
	t.Log(
		m.Morphs2Str(rawtext),
	)
}

func Test_Nouns2Str(t *testing.T) {
	m, e := mecab.New()
	if e != nil {
		t.Error(e)
	}
	defer m.Destroy()
	rawtext := "택시 운전사는 어두운 창밖으로 고개를 내밀어 이따금 고함을 친다."
	t.Log(
		m.Nouns2Str(rawtext),
	)
}

// NOTE: Examples

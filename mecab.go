package mecab

// #include <mecab.h>
// #include <stdio.h>
// #include <stdlib.h>
import "C"
import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"unsafe"
)

/*
1. libmecab doc
http://bbs.nicklib.com/library/MeCab/libmecab.html

2. mecab-doxygen
http://taku910.github.io/mecab/doxygen/structmecab__node__t.html#aff3d4ffe36157f2f2e9bc10f744e6782

3. mecab-ko-dic :
https://bitbucket.org/eunjeon/mecab-ko-dic
https://docs.google.com/spreadsheets/d/1-9blXKjtjeKZqsf4NzHeYJCrr49-nXeRF6D80udfcwY/edit#gid=589544265
*/

// Mecab struct 는 libmecab 바인딩을 위한 인스턴스를 관리하며, New() 또는 NewWithDicPath() 로 인스턴스를 생성하고 사용 후 Destroy()로 포인터를 해제 해주어야 합니다.
type Mecab struct {
	model   *C.struct_mecab_model_t
	tagger  *C.struct_mecab_t
	lattice *C.struct_mecab_lattice_t
}

// Nodes struct 는 NodeItem 의 묶음이며, GetSurfaceSlice와 같이 형태소 분석 후 결과값을 일괄 처리하는 메소드를 제공합니다.
type Nodes struct {
	List []*NodeItem `json:"list"`
}

// NodeItem struct 는 형태소분석 후 하나의 결과값을 표현합니다.
type NodeItem struct {
	Surface string   `json:"surface"`
	TagID   string   `json:"tag_id"`
	Etc     []string `json:"etc"`
}

// Version 는 설치된 mecab 의 버전을 리턴 합니다.
func Version() string {
	return C.GoString(C.mecab_version())
}

// GetLastError 는 mecab 의 마지막 발생 오류값을 리턴합니다.
func GetLastError() string {
	return C.GoString(C.mecab_strerror(nil))
}

// New 는 libmecab 바인딩을 위한 인스턴스를 생성합니다.
func New(args ...string) (*Mecab, error) {
	initArg := C.CString("")
	if len(args) > 0 {
		initArg = C.CString(args[0])
	}

	defer C.free(unsafe.Pointer(initArg))

	m := &Mecab{}
	m.model = C.mecab_model_new2(initArg)
	if m.model == nil {
		return nil, fmt.Errorf("mecab error: %s", GetLastError())
	}
	m.tagger = C.mecab_model_new_tagger(m.model)
	if m.tagger == nil {
		return nil, fmt.Errorf("mecab error: %s", GetLastError())
	}
	m.lattice = C.mecab_model_new_lattice(m.model)
	if m.lattice == nil {
		return nil, fmt.Errorf("mecab error: %s", GetLastError())
	}
	return m, nil
}

// NewWithDicPath 는 사전경로를 표기하여 libmecab 바인딩을 위한 인스턴스를 생성합니다.
// ex) NewWithDicPath("/usr/local/Cellar/mecab-ko-dic-2.0.3-20170922")
func NewWithDicPath(dicpath string) (*Mecab, error) {
	return New(fmt.Sprintf("-d %s", dicpath))
}

// Destroy 는 점유된 libmecab 의 포인터를 해제 합니다.
func (t *Mecab) Destroy() {
	C.mecab_lattice_destroy(t.lattice)
	C.mecab_destroy(t.tagger)
	C.mecab_model_destroy(t.model)
}

// Parse 는 형태소를 분석하여 Nodes 타입으로 리턴하며, args[0] 은 분석할 문장, args[1] 는 원하는 품사만 필터링 할때 포함하여 호출합니다.
// ex) t.Parse("이 문장을 형태소 분석합니다.", "NN*").GetSurfaceSlice()
func (t *Mecab) Parse(args ...string) *Nodes {
	text := C.CString(args[0])
	defer C.free(unsafe.Pointer(text))

	var srch *regexp.Regexp

	if len(args) > 1 {
		srch = regexp.MustCompile(args[1])
	}

	C.mecab_lattice_set_sentence(t.lattice, text)
	C.mecab_parse_lattice(t.tagger, t.lattice)

	res := &Nodes{List: []*NodeItem{}}

	for _, line := range strings.Split(
		C.GoString(C.mecab_lattice_tostr(t.lattice)),
		"\n",
	) {
		if line != "EOS" && len(line) > 0 {
			tmp := strings.Split(line, "\t")
			feature := strings.Split(tmp[1], ",")

			node := &NodeItem{
				Surface: tmp[0],
				TagID:   feature[0],
				Etc:     feature[1:],
			}

			if (srch != nil && srch.MatchString(node.TagID)) || (srch == nil) {
				res.List = append(res.List, node)
			}
		}
	}
	return res
}

// Pos 는 형태소와 품사를 붙여 [][]string 형태로 리턴합니다. ex) [[택시 NNG] [운전사 NNG] [는 JX]]
func (t *Mecab) Pos(text string) [][]string {
	res := [][]string{}
	for _, v := range t.Parse(text).List {
		res = append(res, []string{
			v.Surface,
			v.TagID,
		})
	}
	return res
}

// Morphs 는 형태소만 추출하여 []string 형태로 리턴합니다. ex) [택시 운전사 는 어두운 창밖 으로]
func (t *Mecab) Morphs(text string) []string {
	return t.Parse(text).GetSurfaceSlice()
}

// Nouns 는 명사만 추출하여 []string 형태로 리턴합니다. ex) [택시 운전사 창밖 고개 고함]
func (t *Mecab) Nouns(text string) []string {
	return t.Parse(text, "NN*").GetSurfaceSlice()
}

// GetJSONPretty 는 Nodes 타입의 List 데이터를 JSON string 으로 보기좋게 리턴합니다.
func (t *Nodes) GetJSONPretty() string {
	res, e := json.MarshalIndent(t.List, "", "\t")
	if e != nil {
		return ""
	}
	return string(res)
}

// GetJSONString 는 Nodes 타입의 List 데이터를 JSON string 으로 리턴합니다
func (t *Nodes) GetJSONString() string {
	res, e := json.Marshal(t.List)
	if e != nil {
		return ""
	}
	return string(res)
}

// GetJSONString 는 Nodes 타입의 List 데이터를 JSON []byte 으로 리턴합니다.
func (t *Nodes) GetJSONByte() []byte {
	res, e := json.Marshal(t.List)
	if e != nil {
		return nil
	}
	return res
}

// GetSurfaceSlice 는 Nodes 타입의 List 데이터를 []string 로 리턴합니다.
func (t *Nodes) GetSurfaceSlice() []string {
	res := []string{}
	for _, v := range t.List {
		res = append(res, v.Surface)
	}
	return res
}

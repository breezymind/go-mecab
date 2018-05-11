# go-mecab [![GoDoc](https://godoc.org/github.com/breezymind/go-mecab?status.svg)](https://godoc.org/github.com/breezymind/go-mecab)

> Mecab binding for Go


## Dependency

> mecab-ko 0.996 이상 설치
> 
> [eunjeon/mecab-ko 설치](https://bitbucket.org/eunjeon/mecab-ko)  
> [eunjeon/mecab-ko-dic 설치](https://bitbucket.org/eunjeon/mecab-ko-dic)

```bash
# OSX
brew install mecab-ko
brew install mecab-ko-dic
```

## Reference
    1) mecab-doxygen
    
    http://taku910.github.io/mecab/doxygen/functions.html
    
    2) mecab-ko, mecab-ko-dic
    
    - mecab-ko
    https://bitbucket.org/eunjeon/mecab-ko
    
    - mecab-ko-dic 품사 태그 설명
    https://docs.google.com/spreadsheets/d/1-9blXKjtjeKZqsf4NzHeYJCrr49-nXeRF6D80udfcwY/edit#gid=589544265

## Installation

```bash
export CGO_LDFLAGS="`mecab-config --libs`"
export CGO_CFLAGS="-I`mecab-config --inc-dir`"

go get "github.com/breezymind/go-mecab"
```

## Usage
```go
fmt.Printf("Mecab Version: %s\n", mecab.Version())
// Mecab Version: 0.996/ko-0.9.2

m, _ := mecab.New()
// m, _ := mecab.NewWithDicPath("/usr/local/Cellar/mecab-ko-dic-2.0.3-20170922")
defer m.Destroy()

rawtext := `
    택시 운전사는 어두운 창밖으로 고개를 내밀어
    이따금 고함을 친다. 그때마다 새들이 날아간다 
    이곳은 처음 지나는 벌판과 황혼,
    나는 한번도 만난 적 없는 그를 생각한다

    기형도 / 입속의 검은 잎
`

// 1. 명사 추출 예시
fmt.Println(
    m.Nouns(rawtext),
)
// [택시 운전사 창밖 고개 고함 그때 새 들 이곳 벌판 황혼 나 번 적 그 생각 기형 입속 잎]

// 2. Parse 품사 정규식 필터링 예시
fmt.Println(
    // 명사만 뽑아서 Surface 값을 슬라이스로 리턴
    m.Parse(rawtext, "NN*").GetSurfaceSlice(),    
)
// [택시 운전사 창밖 고개 고함 그때 새 들 이곳 벌판 황혼 나 번 적 그 생각 기형 입속 잎]

// 3. Morphs 형태소 추출 예시
fmt.Println(
    m.Morphs(rawtext),
)
// [택시 운전사 는 어두운 창밖 으로 고개 를 내밀 어 이따금 고함 을 친다 . 그때 마다 새 들 이 날아간다 이곳 은 처음 지나 는 벌판 과 황혼 , 나 는 한 번 도 만난 적 없 는 그 를 생각 한다 기형 / 입속 의 검 은 잎]

// 4. Pos [형태소,품사] 추출 예시
fmt.Println(
    m.Pos(rawtext),
)
// [[택시 NNG] [운전사 NNG] [는 JX] [어두운 VA+ETM] [창밖 NNG] [으로 JKB] [고개 NNG] [를 JKO] [내밀 VV] [어 EC] [이따금 MAG] [고함 NNG] [을 JKO] [친다 VV+EF] [. SF] [그때 NNG] [마다 JX] [새 NNG] [들 XSN] [이 JKS] [이곳 NP] [은 JX] [처음 MAG] [지나 VV] [는 ETM] [벌판 NNG] [과 JC] [황혼 NNG] [, SC] [나 NP] [는 JX] [한 MM] [번 NNBC] [도 JX] [만난 VV+ETM] [적 NNB] [없 VA] [는 ETM] [그 NP] [를 JKO] [생각 NNG] [한다 XSV+EC] [기형 NNG] [도 JX] [/ SC] [입속 NNG] [의 JKG] [검 VA] [은 ETM] [잎 NNG]]

// 5. GetJSONPretty 사용 예시
fmt.Println(
    m.Parse(rawtext, "NN*").GetJSONPretty(),
)
// [
// 	{
// 		"surface": "택시",
// 		"tag_id": "NNG",
// 		"etc": [
// 			"*",
// 			"F",
// 			"택시",
// 			"*",
// 			"*",
// 			"*",
//     },
//     (...생략...)
// ]
```

## Todos

- [ ] go-mecab, test example, godoc 작성

## License
[MIT license](https://opensource.org/licenses/MIT)
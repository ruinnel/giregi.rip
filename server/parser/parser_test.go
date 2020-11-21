package parser

import (
	"fmt"
	"net/url"
	"testing"
	"time"
)

func makeTime(year, month, day, hour, min int) *time.Time {
	t := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)
	return &t
}

func TestParser_Parse(t *testing.T) {
	urls := []string{
		"https://news.v.daum.net/v/20190907114344810",
		"https://news.naver.com/main/read.nhn?mode=LSD&mid=shm&sid1=102&oid=055&aid=0000757196",
		"https://news.v.daum.net/v/20190909040807819",
		"https://news.naver.com/main/read.nhn?mode=LSD&mid=sec&sid1=102&oid=005&aid=0001237528",
		"https://news.v.daum.net/v/20190909200038581",
		"http://news.chosun.com/site/data/html_dir/2019/08/27/2019082700128.html",
	}

	expect := []Result{
		{
			URL:       "https://news.v.daum.net/v/20190907114344810",
			Title:     "[현장연결] 정의당, 조국에 사실상 적격 판정..\"대통령 임명권 존중\"",
			Agency:    "연합뉴스TV",
			Writer:    "양찬주",
			CreatedAt: makeTime(2019, 9, 7, 11, 43),
			UpdatedAt: makeTime(2019, 9, 10, 15, 36),
		},
		{
			URL:       "https://news.naver.com/main/read.nhn?oid=055&aid=0000757196",
			Title:     "장제원 아들, 만취 음주운전…장제원 \"국민께 사과\"",
			Agency:    "SBS",
			Writer:    "한소희",
			Email:     "han@sbs.co.kr",
			CreatedAt: makeTime(2019, 9, 7, 21, 26),
			UpdatedAt: makeTime(2019, 9, 9, 11, 2),
		},
		{
			URL:       "https://news.v.daum.net/v/20190909040807819",
			Title:     "하나둘 늘어나는 개입 정황.. 검찰 수사, 조국 턱밑까지",
			Agency:    "국민일보",
			Writer:    "허경구",
			Cowriter:  "구승은",
			CreatedAt: makeTime(2019, 9, 9, 4, 8),
			UpdatedAt: makeTime(2019, 9, 9, 9, 11),
		},
		{
			URL:       "https://news.naver.com/main/read.nhn?oid=005&aid=0001237528",
			Title:     "하나둘 늘어나는 개입 정황… 검찰 수사, 조국 턱밑까지",
			Agency:    "국민일보",
			Writer:    "허경구",
			Cowriter:  "구승은",
			CreatedAt: makeTime(2019, 9, 9, 4, 8),
			UpdatedAt: nil,
		},
		{
			URL:       "https://news.v.daum.net/v/20190909200038581",
			Title:     "'장제원 아들'의 무서웠던 폭주..CCTV 확인해보니",
			Agency:    "MBC",
			Writer:    "윤상문",
			Email:     "sangmoon@mbc.co.kr",
			CreatedAt: makeTime(2019, 9, 9, 20, 0),
			UpdatedAt: makeTime(2019, 9, 9, 20, 3),
		},
		{
			URL: "http://news.chosun.com/site/data/html_dir/2019/08/27/2019082700128.html",
		},
		nil, // error
	}

	for idx, rawUrl := range urls {
		parsedUrl, _ := url.Parse(rawUrl)
		strippedUrl := StripUrl(parsedUrl)
		result, err := Parse(strippedUrl)
		exp := expect[idx]
		if err != nil {
			if exp != nil {
				t.Error("error occurred.", err)
				return
			} else {
				continue
			}
		}
		if result[URL] != exp[URL] {
			t.Errorf("url not match - %v, %v", result[URL], exp[URL])
			return
		}
		if exp[Title] != nil {
			if result[Title] != exp[Title] {
				t.Error("title not match")
				return
			}
		}
		if exp[Agency] != nil {
			if result[Agency] != exp[Agency] {
				t.Errorf("agency not match - %v, %v", result[Agency], exp[Agency])
				return
			}
		}
		if exp[Writer] != nil {
			writer := exp[Writer].(string)
			if len(writer) > 0 && result[Writer] != exp[Writer] {
				t.Error("writer not match")
				return
			}
		}
		if exp[Email] != nil {
			email := exp[Email].(string)
			if len(email) > 0 && result[Email] != exp[Email] {
				t.Errorf("email not match - %v, %v", result[Email], exp[Email])
				return
			}
		}

		if exp[Cowriter] != nil {
			cowriter := exp[Cowriter].(string)
			if len(cowriter) > 0 && result[Cowriter] != exp[Cowriter] {
				t.Error("co-reporter not match")
				return
			}
		}
		if exp[CreatedAt] != nil {
			createdAt := exp[CreatedAt].(*time.Time)
			if createdAt != nil && !createdAt.Equal(*result[CreatedAt].(*time.Time)) {
				t.Errorf("created At not match, exp - %v, ret - %v\n", createdAt, result[CreatedAt])
				return
			}
		}
		if exp[UpdatedAt] != nil {
			updatedAt := exp[UpdatedAt].(*time.Time)
			if updatedAt != nil && !updatedAt.Equal(*result[UpdatedAt].(*time.Time)) {
				t.Errorf("updated At not match, exp - %v, ret - %v\n", updatedAt, result[UpdatedAt])
				return
			}
		}
	}
}

func TestParser(t *testing.T) {
	//rawUrl, _ := url.Parse("https://news.naver.com/main/read.nhn?oid=055&aid=0000757196")
	rawUrl, _ := url.Parse("https://news.v.daum.net/v/20190909200038581")
	parsed, err := Parse(rawUrl)
	if err != nil {
		t.Error("parse fail", err)
	}
	fmt.Println(parsed)
}

package parser

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"time"
)

const (
	HostDaum  = "news.v.daum.net"
	HostNaver = "news.naver.com"
)

type Key int

const (
	URL Key = iota
	Title
	Writer
	CreatedAt
	UpdatedAt
	Email
	Agency
	Cowriter
)

var keys = []string{
	"url",
	"title",
	"writer",
	"createdAt",
	"updatedAt",
	"email",
	"agency",
	"cowriter",
}

func (k Key) Key() string {
	return keys[k]
}

type Extractor func(selection *goquery.Selection) interface{}
type Applier func(selection *goquery.Selection, result *Result)
type FieldExtractor struct {
	Selector  string
	Extractor Extractor
	Applier   Applier
}

type Parser interface {
	StripUrl(url *url.URL) *url.URL
	Fields() map[Key]FieldExtractor
}

type Result map[Key]interface{}

func (r Result) ToList() []map[string]interface{} {
	var list []map[string]interface{}
	keys := make([]Key, len(r))
	i := 0
	for k := range r {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, key := range keys {
		value := r[key]
		name := key.Key()

		if value == nil {
			continue
		}

		switch value.(type) {
		case *time.Time:
			if value.(*time.Time) != nil {
				list = append(list, map[string]interface{}{"name": name, "value": value.(*time.Time).UnixNano() / int64(time.Millisecond)})
			}
		case time.Time:
			list = append(list, map[string]interface{}{"name": name, "value": value.(*time.Time).UnixNano() / int64(time.Millisecond)})
		default:
			list = append(list, map[string]interface{}{"name": name, "value": value})
		}
	}
	return list
}

func (r Result) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.ToList())
}

func Parse(targetUrl *url.URL) (Result, error) {
	host := targetUrl.Host
	if host != HostDaum && host != HostNaver {
		result := Result{}
		result[URL] = targetUrl.String()
		return result, nil
	}

	res, err := http.Get(targetUrl.String())
	if err != nil {
		return nil, errors.New("request fail")
	}
	defer res.Body.Close()
	contentType := res.Header.Get("Content-Type")
	regex, _ := regexp.Compile("^.*?charset=(.+)$")
	match := regex.FindStringSubmatch(contentType)
	var charset = "UTF-8"
	if len(match) > 0 {
		charset = match[1]
	}

	var data []byte
	if charset == "euc-kr" || charset == "EUC-KR" {
		reader := transform.NewReader(res.Body, korean.EUCKR.NewDecoder())
		data, err = ioutil.ReadAll(reader)
	} else {
		data, err = ioutil.ReadAll(res.Body)
	}

	if err != nil {
		return nil, errors.New("get body fail")
	}

	switch host {
	case HostDaum:
		return processParse(targetUrl, data, NewDaumParser().Fields())
	case HostNaver:
		return processParse(targetUrl, data, NewNaverParser().Fields())
	default:
		// unreached
		return nil, errors.New("not supported")
	}
}

func StripUrl(targetUrl *url.URL) *url.URL {
	host := targetUrl.Host
	switch host {
	case HostDaum:
		return NewDaumParser().StripUrl(targetUrl)
	case HostNaver:
		return NewNaverParser().StripUrl(targetUrl)
	default:
		return targetUrl
	}
}

func SplitByRegex(text string, delimiter string) []string {
	reg := regexp.MustCompile(delimiter)
	indexes := reg.FindAllStringIndex(text, -1)
	lastIdx := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = text[lastIdx:element[0]]
		lastIdx = element[1]
	}
	result[len(indexes)] = text[lastIdx:]
	return result
}

func processParse(url *url.URL, data []byte, fields map[Key]FieldExtractor) (Result, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	result := Result{URL: StripUrl(url).String()}

	for key, field := range fields {
		selector := field.Selector
		extractor := field.Extractor
		applier := field.Applier
		doc.Find(selector).Each(func(idx int, selection *goquery.Selection) {
			if applier != nil {
				applier(selection, &result)
			} else if extractor == nil {
				val := selection.Text()
				if len(val) > 0 {
					result[key] = val
				}
			} else {
				val := extractor(selection)
				if val != nil {
					result[key] = val
				}
			}
		})
	}

	return result, nil
}

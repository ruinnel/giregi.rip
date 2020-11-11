package parser

import (
	"encoding/json"
	"errors"
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

type Parser interface {
	StripUrl(url *url.URL) *url.URL
	Parse(url *url.URL, data []byte) (*Result, error)
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
		return NewDaumParser().Parse(targetUrl, data)
	case HostNaver:
		return NewNaverParser().Parse(targetUrl, data)
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

var emailRegex = regexp.MustCompile(`^.*?([a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}).*?$`)

func ExtractEmail(text string) string {
	match := emailRegex.FindStringSubmatch(text)
	if len(match) > 1 {
		return match[1]
	}
	return ""
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

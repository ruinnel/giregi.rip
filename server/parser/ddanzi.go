package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type DdanziParser struct {
	result     Result
	dateLayout string
	dateRegex  *regexp.Regexp
	emailRegex *regexp.Regexp
}

func NewDdanziParser() *DdanziParser {
	return &DdanziParser{
		result:     Result{},
		dateLayout: "2006-01-02 15:04",
		dateRegex:  regexp.MustCompile(`([0-9]{4}\.[0-9]{2}\.[0-9]{2}\. [0-9]{2}:[0-9]{2})`),
		emailRegex: regexp.MustCompile(`^.*?([a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}).*?$`),
	}
}

func (p *DdanziParser) StripUrl(url *url.URL) *url.URL {
	articleId := url.Query().Get("document_srl")
	if strings.HasSuffix(url.Path, "index.php") {
		url.Path = fmt.Sprintf("/free/%s", articleId)
	}
	url.RawQuery = ""
	return url
}

func (p *DdanziParser) Fields() map[Key]FieldExtractor {
	return map[Key]FieldExtractor{
		Title: {
			Selector: "div.read_header > div.top_title > h1 > a.link",
			Extractor: func(selection *goquery.Selection) interface{} {
				return strings.TrimSpace(selection.Text())
			},
		},
		Writer: {
			Selector: "div.read_header > div.meta > a.author",
			Applier: func(selection *goquery.Selection, result *Result) {
				(*result)[Writer] = strings.TrimSpace(selection.Text())
				classes, exists := selection.Attr("class")
				if exists {
					for _, className := range strings.Split(classes, " ") {
						if strings.HasPrefix(className, "member_") {
							memberId := strings.Replace(className, "member_", "", 1)
							(*result)[WriterId] = strings.TrimSpace(memberId)
						}
					}
				}
			},
		},
		CreatedAt: {
			Selector:  "div.read_header > div.top_title > div.right > p.time",
			Extractor: p.extractCreatedAt,
		},
	}
}

func (p *DdanziParser) extractCreatedAt(selection *goquery.Selection) interface{} {
	text := selection.Text()
	date, err := time.ParseInLocation(p.dateLayout, text, time.Local)
	if err == nil {
		return &date
	}
	return nil
}

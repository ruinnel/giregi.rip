package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type NaverParser struct {
	result      Result
	dateLayout  string
	dateRegex   *regexp.Regexp
	writerRegex *regexp.Regexp
}

func NewNaverParser() *NaverParser {
	return &NaverParser{
		result:      Result{},
		dateLayout:  "2006.01.02. PM 15:04",
		dateRegex:   regexp.MustCompile(`([0-9]{4}\.[0-9]{2}\.[0-9]{2}\. (오전|오후) [0-9]{1,2}:[0-9]{2})`),
		writerRegex: regexp.MustCompile(`^.*?([가-힣]{2,})(.*?)([가-힣]{2,})*(\s*)(기자)*(\s*)([a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}).*?$`),
	}
}

func (np *NaverParser) StripUrl(url *url.URL) *url.URL {
	oid := url.Query().Get("oid")
	aid := url.Query().Get("aid")
	url.RawQuery = fmt.Sprintf("oid=%s&aid=%s", oid, aid)
	return url
}

func (np *NaverParser) Fields() map[Key]FieldExtractor {
	return map[Key]FieldExtractor{
		Title: {
			Selector: "div.article_info > h3#articleTitle",
		},
		Writer: {
			Selector: "div#articleBodyContents",
			Applier: func(selection *goquery.Selection, result *Result) {
				writer, cowriter, email := np.extractWriter(selection)
				if len(writer) > 0 {
					(*result)[Writer] = writer
				}
				if len(cowriter) > 0 {
					(*result)[Cowriter] = cowriter
				}
				if len(email) > 0 {
					(*result)[Email] = email
				}
			},
		},
		CreatedAt: {
			Selector: "div.article_info > div.sponsor",
			Extractor: func(selection *goquery.Selection) interface{} {
				return np.extractDate(selection, "기사입력")
			},
		},
		UpdatedAt: {
			Selector: "div.article_info > div.sponsor",
			Extractor: func(selection *goquery.Selection) interface{} {
				return np.extractDate(selection, "최종수정")
			},
		},
		Agency: {
			Selector:  "meta[property=\"me2:category1\"]",
			Extractor: np.extractAgency,
		},
	}
}

func (np *NaverParser) parseDate(text string) (*time.Time, error) {
	dateText := strings.Replace(text, "오전", "AM", 1)
	dateText = strings.Replace(dateText, "오후", "PM", 1)
	date, err := time.ParseInLocation(np.dateLayout, dateText, time.Local)
	if err != nil {
		return nil, err
	} else {
		return &date, nil
	}
}

func (np *NaverParser) extractWriter(selection *goquery.Selection) (writer, cowriter, email string) {
	for _, node := range selection.Nodes {
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			if child.Type == html.TextNode {
				text := strings.TrimSpace(child.Data)
				if len(text) > 0 {
					match := np.writerRegex.FindStringSubmatch(child.Data)
					if len(match) == 8 {
						writer := match[1]
						cowriter := match[3]
						email := match[7]
						return writer, cowriter, email
					}
				}
			}
		}
	}
	return "", "", ""
}

func (np *NaverParser) extractAgency(selection *goquery.Selection) interface{} {
	for _, node := range selection.Nodes {
		if node.Type == html.ElementNode && node.Data == "meta" {
			for _, attr := range node.Attr {
				if attr.Key == "property" && attr.Val == "me2:category1" {
					for _, at := range node.Attr {
						if at.Key == "content" {
							return at.Val
						}
					}
				}
			}
		}
	}
	return nil
}

func (np *NaverParser) extractDate(selection *goquery.Selection, prefix string) interface{} {
	for _, node := range selection.Nodes {
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			if child.Type == html.TextNode {
				text := strings.TrimSpace(child.Data)
				if text == prefix && child.NextSibling.Type == html.ElementNode && child.NextSibling.Data == "span" {
					dateText := child.NextSibling.FirstChild.Data
					match := np.dateRegex.FindStringSubmatch(dateText)
					if strings.HasPrefix(text, prefix) && len(match) > 1 {
						dateText := match[0]
						if len(dateText) > 0 {
							date, err := np.parseDate(dateText)
							if err == nil {
								return date
							}
						}
					}
				}
			}
		}
	}
	return nil
}

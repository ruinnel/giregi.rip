package parser

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type DaumParser struct {
	result     Result
	dateLayout string
	dateRegex  *regexp.Regexp
	emailRegex *regexp.Regexp
}

func NewDaumParser() *DaumParser {
	return &DaumParser{
		result:     Result{},
		dateLayout: "2006.01.02. 15:04",
		dateRegex:  regexp.MustCompile(`([0-9]{4}\.[0-9]{2}\.[0-9]{2}\. [0-9]{2}:[0-9]{2})`),
		emailRegex: regexp.MustCompile(`^.*?([a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}).*?$`),
	}
}

func (p *DaumParser) StripUrl(url *url.URL) *url.URL {
	url.RawQuery = ""
	return url
}

func (p *DaumParser) Fields() map[Key]FieldExtractor {
	return map[Key]FieldExtractor{
		Title: {
			Selector: "div.head_view > h3.tit_view",
		},
		Writer: {
			Selector: "div.head_view > span.info_view > span.txt_info:first-child",
			Applier:  p.applyWriter,
		},
		CreatedAt: {
			Selector: "div.head_view > span.info_view > span.txt_info > span.num_date",
			Extractor: func(selection *goquery.Selection) interface{} {
				return p.extractDate(selection, "입력")
			},
		},
		UpdatedAt: {
			Selector: "div.head_view > span.info_view > span.txt_info > span.num_date",
			Extractor: func(selection *goquery.Selection) interface{} {
				return p.extractDate(selection, "수정")
			},
		},
		Email: {
			Selector:  "section > p",
			Extractor: p.extractEmail,
		},
		Agency: {
			Selector: "div.head_view > em.info_cp > a.link_cp > img.thumb_g",
			Extractor: func(selection *goquery.Selection) interface{} {
				val, exists := selection.Attr("alt")
				if exists && len(val) > 0 {
					return val
				}
				return nil
			},
		},
	}
}

func (p *DaumParser) splitByRegex(text string, delimiter string) []string {
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

func (p *DaumParser) applyWriter(selection *goquery.Selection, result *Result) {
	reporter := selection.Text()
	splits := p.splitByRegex(reporter, "[^가-힣]+")
	if len(splits) == 1 {
		reporter := strings.TrimSpace(reporter)
		if len(reporter) > 0 {
			(*result)[Writer] = reporter
		}
	} else if len(splits) > 1 {
		last := splits[len(splits)-1]
		if last == "기자" {
			splits = splits[:len(splits)-1]
		}
		if len(splits) > 0 {
			writer := strings.TrimSpace(splits[0])
			if len(writer) > 0 {
				(*result)[Writer] = writer
			}
		}
		if len(splits) > 1 {
			cowriter := strings.TrimSpace(splits[1])
			if len(cowriter) > 0 {
				(*result)[Cowriter] = cowriter
			}
		}
	}
}

func (p *DaumParser) extractDate(selection *goquery.Selection, prevText string) interface{} {
	for _, node := range selection.Nodes {
		if strings.TrimSpace(node.PrevSibling.Data) == prevText {
			text := strings.TrimSpace(node.FirstChild.Data)
			if len(text) > 0 {
				date, err := time.ParseInLocation(p.dateLayout, text, time.Local)
				if err == nil {
					return &date
				}
			}
		}
	}
	return nil
}

func (p *DaumParser) extractEmail(selection *goquery.Selection) interface{} {
	for _, node := range selection.Nodes {
		if node.Type == html.ElementNode && node.Data == "p" {
			child := node.FirstChild
			if child.Type == html.TextNode {
				text := strings.TrimSpace(child.Data)
				if len(text) > 0 {
					match := p.emailRegex.FindStringSubmatch(text)
					if len(match) > 1 && len(match[0]) > 0 {
						return match[1]
					}
				}
			}
		}
	}
	return nil
}

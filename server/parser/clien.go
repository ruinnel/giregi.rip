package parser

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"net/url"
	"strings"
	"time"
)

type ClienParser struct {
	result     Result
	dateLayout string
}

func NewClienParser() *ClienParser {
	return &ClienParser{
		result:     Result{},
		dateLayout: "2006-01-02 15:04:05",
	}
}

func (p *ClienParser) StripUrl(url *url.URL) *url.URL {
	url.RawQuery = ""
	return url
}

func (p *ClienParser) Fields() map[Key]FieldExtractor {
	return map[Key]FieldExtractor{
		Title: {
			Selector:  "div.post_title > h3.post_subject > span",
			Extractor: p.extractTitle,
		},
		WriterId: {
			Selector: "input#writer",
			Extractor: func(selection *goquery.Selection) interface{} {
				val, exists := selection.Attr("value")
				if exists && len(val) > 0 {
					return val
				}
				return nil
			},
		},
		Writer: {
			Selector: "div.post_info > div.post_contact > span.contact_name > span.nickname > span",
		},
		CreatedAt: {
			Selector:  "div.post_view > div.post_author > span > span.fa-clock-o",
			Extractor: p.extractCreatedAt,
		},
	}
}

func (p *ClienParser) extractTitle(selection *goquery.Selection) interface{} {
	for _, node := range selection.Nodes {
		if len(node.Attr) == 0 && node.FirstChild.Type == html.TextNode {
			return strings.TrimSpace(node.FirstChild.Data)
		}
	}
	return nil
}

func (p *ClienParser) extractCreatedAt(selection *goquery.Selection) interface{} {
	for _, node := range selection.Nodes {
		text := strings.TrimSpace(node.NextSibling.Data)
		date, err := time.ParseInLocation(p.dateLayout, text, time.Local)
		if err == nil {
			return &date
		}
	}
	return nil
}

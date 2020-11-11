package parser

import (
	"bytes"
	"fmt"
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
	return new(NaverParser)
}

func (np *NaverParser) getTitle(node *html.Node) string {
	if node.Type == html.ElementNode && node.Data == "h3" {
		for _, attr := range node.Attr {
			if attr.Key == "id" && attr.Val == "articleTitle" {
				return node.FirstChild.Data
			}
		}
	}
	return ""
}

func (np *NaverParser) getNewsAgency(node *html.Node) string {
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
	return ""
}

func (np *NaverParser) getReporterInfo(node *html.Node) (reporter, coreporter, email string) {
	if node.Type == html.ElementNode && node.Data == "div" {
		for _, attr := range node.Attr {
			if attr.Key == "id" && attr.Val == "articleBodyContents" {
				for child := node.FirstChild; child != nil; child = child.NextSibling {
					if child.Type == html.TextNode {
						text := strings.TrimSpace(child.Data)
						if len(text) > 0 {
							extract := np.extractWriter(text)
							if len(extract) == 8 {
								return extract[1], extract[3], extract[7]
							}
						}
					}
				}
			}
		}
	}
	return "", "", ""
}

func (np *NaverParser) extractWriter(text string) []string {
	match := np.writerRegex.FindStringSubmatch(text)
	if len(match) > 1 {
		return match
	}
	return nil
}

func (np *NaverParser) getDateText(node *html.Node, prefix string) string {
	if node.Type == html.ElementNode && node.Data == "div" {
		for _, attr := range node.Attr {
			if attr.Key == "class" && attr.Val == "sponsor" {
				for child := node.FirstChild; child != nil; child = child.NextSibling {
					if child.Type == html.TextNode {
						text := strings.TrimSpace(child.Data)
						if text == prefix && child.NextSibling.Type == html.ElementNode && child.NextSibling.Data == "span" {
							dateText := child.NextSibling.FirstChild.Data
							match := np.dateRegex.FindStringSubmatch(dateText)
							if strings.HasPrefix(text, prefix) && len(match) > 1 {
								return match[0]
							}
						}
					}
				}
			}
		}
	}
	return ""
}

func (np *NaverParser) parse(node *html.Node) {
	if _, ok := np.result[Title]; !ok {
		title := strings.TrimSpace(np.getTitle(node))
		if len(title) > 0 {
			np.result[Title] = title
		}
	}
	if _, ok := np.result[Agency]; !ok {
		agency := strings.TrimSpace(np.getNewsAgency(node))
		if len(agency) > 0 {
			np.result[Agency] = agency
		}
	}
	if _, ok := np.result[Writer]; !ok {
		reporter, coreporter, email := np.getReporterInfo(node)

		if len(reporter) > 0 {
			writer := strings.TrimSpace(reporter)
			if len(writer) > 0 {
				np.result[Writer] = writer
			}
			cowriter := strings.TrimSpace(coreporter)
			if len(cowriter) > 0 {
				np.result[Cowriter] = cowriter
			}
			// 기자가 2명인 경우 email 주소가 누구껀지 특정 하기 힘듬.
			email := strings.TrimSpace(email)
			if len(email) > 0 {
				np.result[Email] = email
			}
		}
	}

	if _, ok := np.result[CreatedAt]; !ok {
		reportedAt := np.getDateText(node, "기사입력")
		if len(reportedAt) > 0 {
			reportedAt = strings.Replace(reportedAt, "오전", "AM", 1)
			reportedAt = strings.Replace(reportedAt, "오후", "PM", 1)
			date, err := time.ParseInLocation(np.dateLayout, reportedAt, time.Local)
			if err == nil {
				np.result[CreatedAt] = &date
			}
		}
	}
	if _, ok := np.result[UpdatedAt]; !ok {
		updatedAt := np.getDateText(node, "최종수정")
		if len(updatedAt) > 0 {
			updatedAt = strings.Replace(updatedAt, "오전", "AM", 1)
			updatedAt = strings.Replace(updatedAt, "오후", "PM", 1)
			date, err := time.ParseInLocation(np.dateLayout, updatedAt, time.Local)
			if err == nil {
				np.result[UpdatedAt] = &date
			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		np.parse(child)
	}
}

func (np *NaverParser) StripUrl(url *url.URL) *url.URL {
	oid := url.Query().Get("oid")
	aid := url.Query().Get("aid")
	url.RawQuery = fmt.Sprintf("oid=%s&aid=%s", oid, aid)
	return url
}

func (np *NaverParser) Parse(url *url.URL, data []byte) (Result, error) {
	node, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	np.result = Result{}
	np.dateLayout = "2006.01.02. PM 15:04"
	np.dateRegex = regexp.MustCompile(`([0-9]{4}\.[0-9]{2}\.[0-9]{2}\. (오전|오후) [0-9]{1,2}:[0-9]{2})`)
	np.writerRegex = regexp.MustCompile(`^.*?([가-힣]{2,})(.*?)([가-힣]{2,})*(\s*)(기자)*(\s*)([a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}).*?$`)
	np.parse(node)
	np.result[URL] = np.StripUrl(url).String()
	return np.result, nil
}

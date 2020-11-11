package parser

import (
	"bytes"
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
	inSection  bool
}

func NewDaumParser() *DaumParser {
	return new(DaumParser)
}

func (dp *DaumParser) getTitle(node *html.Node) string {
	if node.Type == html.ElementNode && node.Data == "h3" {
		for _, attr := range node.Attr {
			if attr.Key == "class" && attr.Val == "tit_view" {
				return node.FirstChild.Data
			}
		}
	}
	return ""
}

func (dp *DaumParser) getNewsAgency(node *html.Node) string {
	if node.Type == html.ElementNode && node.Data == "meta" {
		for _, attr := range node.Attr {
			if attr.Key == "name" && attr.Val == "article:media_name" {
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

func (dp *DaumParser) getReporter(node *html.Node) string {
	if node.Type == html.ElementNode && node.Data == "span" {
		for _, attr := range node.Attr {
			if attr.Key == "class" && attr.Val == "info_view" {
				for child := node.FirstChild; child != nil; child = child.NextSibling {
					if child.Type == html.ElementNode && child.Data == "span" {
						for _, childAttr := range child.Attr {
							if childAttr.Key == "class" && childAttr.Val == "txt_info" {
								return child.FirstChild.Data
							}
						}
					}
				}
			}
		}
	}
	return ""
}

func (dp *DaumParser) getDateText(node *html.Node) string {
	if node.Type == html.ElementNode && node.Data == "span" {
		for _, attr := range node.Attr {
			if attr.Key == "class" && attr.Val == "num_date" {
				text := node.FirstChild.Data
				match := dp.dateRegex.FindStringSubmatch(text)
				if len(match) > 1 {
					return match[0]
				}
			}
		}
	}
	return ""
}

func (dp *DaumParser) getEmail(node *html.Node) string {
	parent := node.Parent
	if parent != nil &&
		parent.Type == html.ElementNode && parent.Data == "section" &&
		node.Type == html.ElementNode && node.Data == "p" {
		child := node.FirstChild
		if child.Type == html.TextNode {
			text := strings.TrimSpace(child.Data)
			if len(text) > 0 {
				email := ExtractEmail(text)
				if len(email) > 0 {
					return email
				}
			}
		}
	}
	return ""
}

func (dp *DaumParser) parse(node *html.Node) {
	if _, ok := dp.result[Title]; !ok {
		title := strings.TrimSpace(dp.getTitle(node))
		if len(title) > 0 {
			dp.result[Title] = title
		}
	}
	if _, ok := dp.result[CreatedAt]; !ok {
		reportedAt := dp.getDateText(node)
		if len(reportedAt) > 0 {
			date, err := time.ParseInLocation(dp.dateLayout, reportedAt, time.Local)
			if err == nil {
				dp.result[CreatedAt] = &date
			}
		}
	} else { // updatedAt은 createdAt 보다 뒤에 나옴.
		if _, ok := dp.result[UpdatedAt]; !ok {
			updatedAt := dp.getDateText(node)
			if len(updatedAt) > 0 {
				date, err := time.ParseInLocation(dp.dateLayout, updatedAt, time.Local)
				if err == nil {
					dp.result[UpdatedAt] = &date
				}
			}
		}
	}

	if _, ok := dp.result[Agency]; !ok {
		agency := strings.TrimSpace(dp.getNewsAgency(node))
		if len(agency) > 0 {
			dp.result[Agency] = agency
		}
	}
	if _, ok := dp.result[Writer]; !ok {
		reporter := strings.TrimSpace(dp.getReporter(node))
		splits := SplitByRegex(reporter, "[^가-힣]+")
		if len(splits) == 1 {
			reporter := strings.TrimSpace(reporter)
			if len(reporter) > 0 {
				dp.result[Writer] = reporter
			}
		} else if len(splits) > 1 {
			last := splits[len(splits)-1]
			if last == "기자" {
				splits = splits[:len(splits)-1]
			}
			if len(splits) > 0 {
				writer := strings.TrimSpace(splits[0])
				if len(writer) > 0 {
					dp.result[Writer] = writer
				}
			}
			if len(splits) > 1 {
				cowriter := strings.TrimSpace(splits[1])
				if len(cowriter) > 0 {
					dp.result[Cowriter] = cowriter
				}
			}
		}
	}
	// 기자가 2명인 경우 email 주소가 누구껀지 특정 하기 힘듬.
	if _, ok := dp.result[Email]; !ok {
		email := dp.getEmail(node)
		if len(email) > 0 {
			dp.result[Email] = email
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		dp.parse(child)
	}
}

func (dp *DaumParser) StripUrl(url *url.URL) *url.URL {
	url.RawQuery = ""
	return url
}

func (dp *DaumParser) Parse(url *url.URL, data []byte) (Result, error) {
	node, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	dp.result = Result{}
	dp.dateLayout = "2006.01.02. 15:04"
	dp.dateRegex = regexp.MustCompile(`([0-9]{4}\.[0-9]{2}\.[0-9]{2}\. [0-9]{2}:[0-9]{2})`)
	dp.parse(node)
	dp.result[URL] = dp.StripUrl(url).String()
	return dp.result, nil
}

package util

import (
	"fmt"
	"regexp"
	"strconv"
)

//find the first group matches of the text of the leftmost match of the regular expression in html
func ParseHtmlReg(reg string, html string) (string, error) {
	return ParseHtmlRegIndex(reg, html, 1)
}

//find the index group matches of the text of the leftmost match of the regular expression in html
//filter html tags
func ParseHtmlRegIndex(reg string, html string, index int) (string, error) {
	compile := regexp.MustCompile(reg)
	subMatch := compile.FindStringSubmatch(html)
	if subMatch != nil && len(subMatch) > index {
		//fmt.Println(subMatch)
		return HtmlTagFilter(subMatch[index]), nil
	}
	return "", fmt.Errorf("no match")
}

func ParseHtmlRegSlice(reg string, html string) ([]string, error) {
	compile := regexp.MustCompile(reg)
	subMatch := compile.FindStringSubmatch(html)
	if subMatch != nil {
		return subMatch, nil
	}
	return nil, fmt.Errorf("no match")
}

func HtmlTagFilter(html string) string {
	compile := regexp.MustCompile(`<[\s\S]*?>`)
	html = compile.ReplaceAllString(html, "")
	return html
}

func ParseStringToInt64(str string) int64 {
	ret, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return ret
}

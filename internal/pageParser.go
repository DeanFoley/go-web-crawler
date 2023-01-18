package app

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

func ExtractAnchors(webpage []byte) ([]string, error) {
	stringifiedWebpage := string(webpage)

	tkn := html.NewTokenizer(strings.NewReader(stringifiedWebpage))

    var anchors []string

    for {
        tt := tkn.Next()

        switch {
        case tt == html.ErrorToken:
			if tkn.Err() == io.EOF {
				return anchors, nil
			} else {
				return nil, tkn.Err()
			}

		case tt == html.StartTagToken:
            t := tkn.Token()
            if t.Data == "a" {
				for _, value := range t.Attr {
					if value.Key == "href" {
						anchors = append(anchors, value.Val)
					}
				}
			}
        }
    }
}

func FormatAnchors(anchors []string) []byte {
	bytes := make([]byte, 0)
	for _, url := range anchors {
		bytes = append(bytes, []byte(url + "\n")...)
	}
	return bytes
}
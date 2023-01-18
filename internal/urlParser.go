package app

import "net/url"

func ValidateUrl(requestedUrl string) (*url.URL, error) {
	return url.ParseRequestURI(requestedUrl)
}
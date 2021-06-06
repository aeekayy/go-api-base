package handlers

import (
	"errors"
	"fmt"
	"net/url"
)

const (
	allowedOriginHosts = "app.aeekay.co"
)

func ReturnAccessControlAllowOrigin(cors map[string]bool, checkUrl string) (string, error) {
	if len(cors) == 0 {
		return "", errors.New("Invalid Access-Control-Allow-Origin string entered")
	}

	u, err := url.Parse(checkUrl)

	if err != nil {
		return "", err
	}

	// check CORS
	// if the URL is local.veritione.com, allow http
	if u.Hostname() == allowedOriginHosts {
		port := ""
		if u.Port() != "80" && u.Port() != "443" {
			port = ":" + u.Port()
		}
		return fmt.Sprintf("%s://%s:%s", "http", allowedOriginHosts, string(port)), nil
	}

	if cors[u.Hostname()] {
		host := "https://" + u.Hostname()
		if u.Port() != "80" && u.Port() != "443" {
			host = host + ":" + u.Port()
		}
		return host, nil
	}

	return "", errors.New("Access-Control-Allow-Origin not found")
}

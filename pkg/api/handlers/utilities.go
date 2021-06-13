package handlers

import (
	"errors"
	"fmt"
	"net/url"
)

const (
	allowedOriginHosts = "app.aeekay.co"
)

// ReturnAccessControlAllowOrigin returns the appropriate Access-Control-Allow-Origin
// based on a map of acceptable domains
// TODO: Account for when the origin is not passed in
func ReturnAccessControlAllowOrigin(cors map[string]bool, checkURL string) (string, error) {
	if len(cors) == 0 {
		return "", errors.New("Invalid Access-Control-Allow-Origin string entered")
	}

	u, err := url.Parse(checkURL)

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

	// ignore for localhost
	if u.Hostname() == "localhost" {
		return checkURL, nil
	}

	fmt.Println(checkURL)
	fmt.Println(u.Hostname())

	return "", errors.New("Access-Control-Allow-Origin not found")
}

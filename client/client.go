package client

import (
	"net/http"
	"net/http/cookiejar"
)

var BaseURL = "https://cskh.npc.com.vn"

type Client struct {
	client *http.Client
}

func New() (*Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: &http.Client{
			Jar: jar,
		},
	}, nil
}


package main

import (
	"encoding/xml"

	"github.com/go-resty/resty/v2"
)

const CbrDailyUrl = "https://www.cbr-xml-daily.ru/daily_utf8.xml"

type Client struct {
	resty *resty.Client
}

func NewClient() *Client {
	return &Client{
		resty: resty.New(),
	}
}

func (c *Client) GetCbrDaily() (*ValCurs, error) {
	resp, err := c.resty.R().Get(CbrDailyUrl)
	if err != nil {
		return nil, err
	}
	var valCurs *ValCurs
	if err := xml.Unmarshal(resp.Body(), &valCurs); err != nil {
		return nil, err
	}
	return valCurs, nil
}

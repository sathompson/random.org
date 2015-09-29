package randomorg

import (
    "net/http"
)

type Client interface {
    URL() string
    SetURL(url string)
    NextId() int
    SetNextId(id int)
    IncrementId() int
}

type client struct {
    apiKey string
    url string
    httpClient *http.Client
    id int
}

func NewClient(apiKey string) Client {
    return &client{apiKey,URL,new(http.Client),0}
}

func (c *client) URL() string {
    return c.url
}

func (c *client) SetURL(url string) {
    c.url = url
}

func (c *client) NextId() int {
    return c.id
}

func (c *client) SetNextId(id int) {
    c.id = id
}

func (c *client) IncrementId() (id int) {
    id = c.id
    c.id += 1
    return c.id
}
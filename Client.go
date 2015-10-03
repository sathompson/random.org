package randomorg

import (
    "net/http"
    "encoding/json"
    "bytes"
    "errors"
    "io/ioutil"
)

const (
  URL = "https://api.random.org/json-rpc/1/invoke"
  ContentType = "application/json-rpc"
)

var httpclient *http.Client = new(http.Client)

// Client interface

type Client interface {
  URL() string
  SetURL(url string)
  HTTPClient() *http.Client
  SetHTTPClient(httpClient *http.Client)
  NextId() int
  SetNextId(id int)
  IncrementId() int
  BitsUsed() int      //By last request
  BitsLeft() int
  RequestsLeft() int
  AdvisoryDelay() int //From last request
  // Requests to random.org
  GenerateIntegers(n, min, max int) ([]int, error)
}

// Client implementation

type client struct {
  apiKey string
  url string
  httpClient *http.Client
  id int
  bitsUsed, bitsLeft, requestsLeft, advisoryDelay int
}

func NewClient(apiKey string) Client {
    return &client{apiKey,URL,httpclient,0,-1,-1,-1,-1}
}

// Utility functions

func (c *client) makeHTTPRequest(r *request) (httpRequest *http.Request, e error) {
    jsonRequest, e := json.Marshal(r)
    if (e != nil) {
        return
    }
    
    body := bytes.NewReader(jsonRequest)
    httpRequest, e = http.NewRequest("POST", c.url, body)
    if (e != nil) {
        return
    }
    
    httpRequest.Header.Add("Content-Type", ContentType)
    
    return
}

func (c *client) sendRequest(r *request, response interface{}) (e error) {
    httpRequest, e := c.makeHTTPRequest(r)
    
    httpResponse, e := c.httpClient.Do(httpRequest)
    if (httpResponse.StatusCode != 200) {
        e = errors.New(httpResponse.Status)
    }
    
    body, e := ioutil.ReadAll(httpResponse.Body)
    if (e != nil) {
        return
    }
    
    json.Unmarshal(body, response)
    return
}

func (c *client) updateMetaData(bitsUsed, bitsLeft, requestsLeft, advisoryDelay int) {
  c.bitsUsed = bitsUsed
  c.bitsLeft = bitsLeft
  c.requestsLeft = requestsLeft
  c.advisoryDelay = advisoryDelay
}

// Fulfilling Client interface

func (c *client) URL() string {
    return c.url
}

func (c *client) SetURL(url string) {
    c.url = url
}

func (c *client) HTTPClient() *http.Client {
    return c.httpClient
}

func (c *client) SetHTTPClient(httpClient *http.Client) {
    c.httpClient = httpClient
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

func (c *client) BitsUsed() int {
  return c.bitsUsed
}

func (c *client) BitsLeft() int {
  return c.bitsLeft
}

func (c *client) RequestsLeft() int {
  return c.requestsLeft
}

func (c *client) AdvisoryDelay() int {
  return c.advisoryDelay
}
package randomorg

import (
    "testing"
    "fmt"
    "bytes"
)

var testClient = NewClient(apiKey).(*client)

func TestMakeHTTPRequest(t *testing.T) {
    params := &generateIntegersParams{testClient.apiKey, 5, 1, 10}
    r := &request{generateIntegersMethod, params, testClient.IncrementId()}
    httpRequest, e := testClient.makeHTTPRequest(r)
    if (e != nil) {
        t.Error(e)
    }
    
    fmt.Println(httpRequest)
    
    buf := new(bytes.Buffer)
    buf.ReadFrom(httpRequest.Body)
    s := buf.String()
    fmt.Println(s)
}

func TestGenerateIntegers(t *testing.T) {
  nums, e := testClient.GenerateIntegers(4, 1, 10)
  if (e != nil) {
    t.Error(e)
  }
  
  fmt.Println(nums)
}
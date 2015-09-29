package randomorg

import (
    "fmt"
)

type generateIntegersParams struct {
    apiKey string
    n, min, max int
}

func (params *generateIntegersParams) MarshalJSON() ([]byte, error) {
    return []byte(fmt.Sprintf("%+v",params)), nil
}

func (c *client) GenerateIntegers(n, min, max int) (result []int, e error) {
    params := &generateIntegersParams{c.apiKey, n, min, max}
    r := &request{generateIntegersMethod, params, c.IncrementId()}
    
    httpResponse, e := c.sendRequest(r)
    if (e != nil) {
        return
    }
    
    // Some stuff goes here
    
    return
}
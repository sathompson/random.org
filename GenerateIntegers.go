package randomorg

import (
    "fmt"
)

type generateIntegersParams struct {
    apiKey string
    n, min, max int
}

func (params *generateIntegersParams) MarshalJSON() (jsonBytes []byte, e error) {
    jsonBytes = []byte(fmt.Sprintf("{%q:%q,%q:%v,%q:%v,%q:%v}",
                                  "apiKey", params.apiKey,
                                  "n", params.n,
                                  "min", params.min,
                                  "max", params.max))
    return
}

func (c *client) GenerateIntegers(n, min, max int) (result []int, e error) {
    params := &generateIntegersParams{c.apiKey, n, min, max}
    r := &request{generateIntegersMethod, params, c.IncrementId()}
    
    httpResponse, e := c.sendRequest(r)
    if (e != nil) {
        return
    }
    
    //stuff goes here
    
    return
}
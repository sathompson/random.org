package randomorg

import (
    "fmt"
    "errors"
)

const generateIntegersMethod = "generateIntegers"

// Actual function

func (c *client) GenerateIntegers(n, min, max int) (result []int, e error) {
    params := &generateIntegersParams{c.apiKey, n, min, max}
    req := &request{generateIntegersMethod, params, c.IncrementId()}
    
    response := new(generateIntegersResponse)
    
    e = c.sendRequest(req, response)
    if (e != nil) {
        return
    }
    if (response.Error.Message != "") {
        e = errors.New(response.Error.Message)
        return
    }
    
    result = response.Result.Random.Data
    return
}

// Params

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

// Results

type generateIntegersResponse struct {
    Jsonrpc string
    Result struct {
        Random struct {
            Data []int
            CompletionTime string
        }
        BitsUsed, BitsLeft, RequestsLeft, AdvisoryDelay int
    }
    Error struct {
        Code int
        Message string
        Data interface{}
    }
    Id int
}
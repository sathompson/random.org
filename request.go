package randomorg

import (
    "encoding/json"
    "fmt"
)

const JSONRPC = "2.0"

type request struct {
    method string
    params json.Marshaler
    id int
}

func (r *request) MarshalJSON() (jsonbytes []byte, e error) {
    params, e := json.Marshal(r.params)
    if (e != nil) {
        return
    }
    jsonbytes = []byte(
        fmt.Sprintf("{%q:%q,%q:%q,%q:%s,%q:%v}",
                    "jsonrpc", JSONRPC,
                    "method", r.method,
                    "params", params,
                    "id", r.id))
    return
}
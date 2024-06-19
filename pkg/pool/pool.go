package pool

import (
    "sync"

    "github.com/valyala/fasthttp"
)

var connPool = sync.Pool{
    New: func() interface{} {
        return &fasthttp.Client{
            // Allow HTTP and HTTPS connections
            Dial: fasthttp.Dial,
        }
    },
}

func GetConnection() *fasthttp.Client {
    return connPool.Get().(*fasthttp.Client)
}

func PutConnection(client *fasthttp.Client) {
    connPool.Put(client)
}

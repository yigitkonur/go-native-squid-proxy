package test

import (
    "testing"
    "net/http/httptest"
    "github.com/stretchr/testify/assert"
    "proxy-server/pkg/handler"
    "github.com/valyala/fasthttp"
)

func TestHandleRequest(t *testing.T) {
    req := httptest.NewRequest("GET", "http://example.com/foo", nil)
    resp := httptest.NewRecorder()

    fastReq := &fasthttp.Request{}
    fastReq.SetRequestURI(req.URL.String())

    fastCtx := &fasthttp.RequestCtx{}
    fastCtx.Init(fastReq, nil, nil)
    handler.HandleRequest(fastCtx)

    // Ensure the response status code is OK (200)
    assert.Equal(t, fasthttp.StatusOK, fastCtx.Response.StatusCode())
}

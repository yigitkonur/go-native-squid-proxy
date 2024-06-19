package handler

import (
    "io"
    "net"

    "github.com/valyala/fasthttp"
    "proxy-server/pkg/metrics"
    "proxy-server/pkg/pool"
)

// HandleRequest handles both normal HTTP requests and HTTP CONNECT requests for HTTPS
func HandleRequest(ctx *fasthttp.RequestCtx) {
    metrics.IncrementRequestCounter()

    // Handle HTTP CONNECT method for HTTPS proxying
    if string(ctx.Method()) == fasthttp.MethodConnect {
        handleTunneling(ctx)
    } else {
        handleHTTP(ctx)
    }
}

// handleHTTP handles standard HTTP requests
func handleHTTP(ctx *fasthttp.RequestCtx) {
    client := pool.GetConnection()
    defer pool.PutConnection(client)

    req := &ctx.Request
    resp := &ctx.Response

    if err := client.Do(req, resp); err != nil {
        ctx.Error("Failed to process request: "+err.Error(), fasthttp.StatusInternalServerError)
        return
    }
}

// handleTunneling handles HTTP CONNECT requests
func handleTunneling(ctx *fasthttp.RequestCtx) {
    destinationConn, err := net.Dial("tcp", string(ctx.Host()))
    if err != nil {
        ctx.Error("Failed to connect to destination", fasthttp.StatusServiceUnavailable)
        return
    }

    // Send 200 OK response to the client
    ctx.SetStatusCode(fasthttp.StatusOK)
    ctx.Response.SetBodyRaw([]byte{})

    ctx.Hijack(func(clientConn net.Conn) {
        defer clientConn.Close()
        defer destinationConn.Close()

        go func() {
            _, _ = io.Copy(destinationConn, clientConn)
        }()
        _, _ = io.Copy(clientConn, destinationConn)
    })
}

// transfer copies data between source and destination connections
func transfer(destination net.Conn, source net.Conn) {
    _, _ = io.Copy(destination, source)
}

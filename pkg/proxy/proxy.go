package proxy

import (
    "go.uber.org/zap"
    "github.com/valyala/fasthttp"
    "proxy-server/pkg/config"
    "proxy-server/pkg/handler"
)

type ProxyServer struct {
    cfg    *config.Config
    logger *zap.SugaredLogger
    server *fasthttp.Server
}

func NewProxyServer(cfg *config.Config, logger *zap.SugaredLogger) *ProxyServer {
    return &ProxyServer{
        cfg:    cfg,
        logger: logger,
        server: &fasthttp.Server{
            Handler:            handler.HandleRequest,
            MaxConnsPerIP:      cfg.MaxConnections,
            MaxRequestsPerConn: 10000, // Example configuration
        },
    }
}

func (p *ProxyServer) Start() error {
    return p.server.ListenAndServe(p.cfg.ServerAddress)
}

func (p *ProxyServer) Shutdown() error {
    return p.server.Shutdown()
}

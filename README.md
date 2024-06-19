# GoNativeSquidProxy

GoNativeSquidProxy is a high-performance, scalable proxy server fully written in Go, designed to efficiently handle HTTP/HTTPS requests as a modern alternative to Squid.

## Features

- **High Performance**: Leveraging Go's goroutines for efficient handling of multiple connections concurrently.
- **HTTP/1.1 and HTTP/2 Support**: Handles both HTTP/1.1 and HTTP/2 protocols.
- **HTTPS Proxying**: Supports tunneling for HTTPS connections.
- **Basic Authentication and Access Control**: (Planned for future development)
- **Logging**: Structured logging using Uber's `zap` library.
- **Monitoring**: Integrates with Prometheus for metrics collection.
- **Configuration Management**: Uses `viper` for flexible configuration.

## Requirements

- Go 1.18 or higher
- Access to a Unix-like environment (Linux, macOS)

## Installation

### Clone the Repository

```bash
git clone https://github.com/yourusername/go-native-squid-proxy.git
cd go-native-squid-proxy
```

### Initialize the Go Module

```bash
go mod init go-native-squid-proxy
```

### Download Dependencies

```bash
go get github.com/spf13/viper
go get go.uber.org/zap
go get github.com/prometheus/client_golang/prometheus/promhttp
go get github.com/stretchr/testify/assert
go get github.com/valyala/fasthttp
```

### Build the Project

```bash
cd cmd/proxy
go build -o proxy-server
mv proxy-server ../../
```

## Usage

### Configuration

Create a `config.yaml` file in the root directory of the project with the following content:

```yaml
ServerAddress: ":8080"
MaxConnections: 10000
LogLevel: "info"
```

### Run the Proxy Server

```bash
./proxy-server
```

### Test the Proxy Server

#### HTTP Request

```bash
curl -x 127.0.0.1:8080 http://ifconfig.io
```

#### HTTPS Request

```bash
curl -x 127.0.0.1:8080 https://ifconfig.io
```

## Project Structure

```plaintext
go-native-squid-proxy/
├── cmd/
│   └── proxy/
│       └── main.go
├── config.yaml
├── go.mod
├── go.sum
├── pkg/
│   ├── config/
│   │   └── config.go
│   ├── handler/
│   │   └── handler.go
│   ├── log/
│   │   └── log.go
│   ├── metrics/
│   │   └── metrics.go
│   ├── pool/
│   │   └── pool.go
│   └── proxy/
│       └── proxy.go
└── test/
    └── proxy_test.go
```

## Code Overview

- **cmd/proxy/main.go**: Entry point of the application. Loads configuration, sets up logging, initializes the proxy server, and handles graceful shutdown.
- **pkg/config/config.go**: Handles loading and validating configuration using `viper`.
- **pkg/handler/handler.go**: Contains the request handling logic for both HTTP and HTTPS requests.
- **pkg/log/log.go**: Sets up structured logging using Uber's `zap` library.
- **pkg/metrics/metrics.go**: Sets up Prometheus metrics for monitoring.
- **pkg/pool/pool.go**: Implements a connection pool to reuse `fasthttp.Client` instances.
- **pkg/proxy/proxy.go**: Initializes and starts the `fasthttp` server, and handles server shutdown.
- **test/proxy_test.go**: Contains unit tests for the proxy server.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you have any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [fasthttp](https://github.com/valyala/fasthttp): High-performance HTTP package for Go.
- [zap](https://github.com/uber-go/zap): High-performance logging library.
- [viper](https://github.com/spf13/viper): Configuration management library.
- [Prometheus](https://prometheus.io/): Monitoring and alerting toolkit.
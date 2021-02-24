// Package quic implements the QUIC API for Client-to-Server Connections
// https://w3c.github.io/webrtc-quic/
package quic

import (
	"fmt"
	"io"

	"github.com/pion/logging"
	"github.com/sumeet-tiwari-samespace/quic/internal/wrapper"
)

// Transport is a quic transport focused on client/server use cases.
type Transport struct {
	TransportBase
}

// NewTransport creates a new Transport
func NewTransport(url string, config *Config) (*Transport, error) {
	if config.LoggerFactory == nil {
		config.LoggerFactory = logging.NewDefaultLoggerFactory()
	}

	cfg := config.clone()
	cfg.SkipVerify = true // Using self signed certificates for now

	s, err := wrapper.Dial(url, cfg)
	if err != nil {
		return nil, err
	}

	t := &Transport{}
	t.TransportBase.log = config.LoggerFactory.NewLogger("quic")
	return t, t.TransportBase.startBase(s)
}

// NewServer accept listen for testing
func NewServer(url string, config *Config) (io.Closer, error) {
	cfg := config.clone()
	cfg.SkipVerify = true // Using self signed certificates for now

	l, err := wrapper.Listen(url, cfg)
	if err != nil {
		return nil, err
	}
        return l, nil
}

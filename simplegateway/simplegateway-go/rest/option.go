// Code generated by jzero. DO NOT EDIT.
// type: rest_frame_option

package rest

import (
	"net/http"
	"time"
)

func WithClient(client *http.Client) Opt {
	return func(c *RESTClient) error {
		c.client = client
		return nil
	}
}

func WithHeaders(headers http.Header) Opt {
	return func(c *RESTClient) error {
		c.headers = headers
		return nil
	}
}

func WithProtocol(protocol string) Opt {
	return func(c *RESTClient) error {
		c.protocol = protocol
		return nil
	}
}

func WithAddr(addr string) Opt {
	return func(c *RESTClient) error {
		c.addr = addr
		return nil
	}
}

func WithPort(port string) Opt {
	return func(c *RESTClient) error {
		c.port = port
		return nil
	}
}

func WithRetryTimes(times int) Opt {
	return func(c *RESTClient) error {
		c.retryTimes = times
		return nil
	}
}

func WithRetryDelay(time time.Duration) Opt {
	return func(c *RESTClient) error {
		c.retryDelay = time
		return nil
	}
}

func WithGatewayPrefix(prefix string) Opt {
	return func(c *RESTClient) error {
		c.gatewayPrefix = prefix
		return nil
	}
}
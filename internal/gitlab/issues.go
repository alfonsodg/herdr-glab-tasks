package gitlab

import "errors"

var (
	ErrGlabMissing  = errors.New("glab not found")
	ErrUnauthorized = errors.New("glab is not authenticated")
)

type Client struct {
	glab string
	host string
}

func NewClient(glab, host string) *Client {
	return &Client{glab: glab, host: host}
}

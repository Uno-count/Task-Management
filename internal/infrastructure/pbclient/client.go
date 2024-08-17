package pbclient

import "github.com/pocketbase/pocketbase"

type Client struct {
	App *pocketbase.PocketBase
}

func NewClient(app *pocketbase.PocketBase) *Client {
	return &Client{App: app}
}

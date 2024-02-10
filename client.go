package spreads

import (
	"context"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Client struct {
	ctx    context.Context
	Sheets *sheets.Service

	spreadID  string
	sheetName string
	rangeKey  string
}

func New(ctx context.Context, credFIleOrByteData any) *Client {
	var op option.ClientOption
	switch v := credFIleOrByteData.(type) {
	case string:
		op = option.WithCredentialsFile(v)
	case []byte:
		op = option.WithCredentialsJSON(v)

	default:
		log.Fatal("invalid type")
	}

	config, err := sheets.NewService(ctx, op)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &Client{
		ctx:    ctx,
		Sheets: config,

		sheetName: "",
		rangeKey:  "",
	}
}

func (c *Client) SetSpreadID(spreadID string) *Client {
	c.spreadID = spreadID
	return c
}

func (c *Client) SetSheetName(sheetName string) *Client {
	c.sheetName = sheetName
	return c
}

func (c *Client) SetRangeKey(rangeKey string) *Client {
	c.rangeKey = rangeKey
	return c
}

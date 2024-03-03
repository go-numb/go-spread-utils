package spreads

import (
	"context"
	"fmt"
	"log"
	"os"

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
	var (
		config *sheets.Service
		err    error
	)
	switch v := credFIleOrByteData.(type) {
	case string:
		if f, err := os.Stat(v); err == nil && !f.IsDir() {
			config, err = sheets.NewService(ctx, option.WithCredentialsFile(v))
			if err != nil {
				log.Fatal(fmt.Errorf("failed to create new service with credential file, error: %v", err))
				return nil
			}
		} else {
			config, err = sheets.NewService(ctx)
			if err != nil {
				log.Fatal(fmt.Errorf("failed to create new service with gcloud auth, error: %v", err))
				return nil
			}
		}

	case []byte:
		config, err = sheets.NewService(ctx, option.WithCredentialsJSON(v))
		if err != nil {
			log.Fatal(fmt.Errorf("failed to create new service with credential byte data, error: %v", err))
			return nil
		}

	case nil:
		// production or gcloud auth
		config, err = sheets.NewService(ctx)
		if err != nil {
			log.Fatal(fmt.Errorf("failed to create new service with gcloud auth, error: %v", err))
			return nil
		}

	default:
		config, err = sheets.NewService(ctx)
		if err != nil {
			log.Println(err)
			return nil
		}
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

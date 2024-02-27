package spreads

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	var testData = []interface{}{
		"test-credential.json",
		// []byte("xxxxxxxxxxxxxxxxxxxxxxxx"),
		nil,
	}

	for _, v := range testData {
		c := New(context.Background(), v)
		assert.NotNil(t, c, "client should not be nil")
	}
}

package context

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRequestId(t *testing.T) {

	c := context.Background()
	id := "1"
	c = SetRequestId(c, id)
	var pick string
	require.NotPanics(t, func() {
		pick = RequestId(c)
	})
	require.Equal(t, id, pick)
}
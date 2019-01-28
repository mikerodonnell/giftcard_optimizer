package gift

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const prices = `taffy, 111
chocolate, 390
mint, 11`

func TestOptimize(t *testing.T) {
	items, err := NewGiftList(strings.Split(prices, "\n"))
	require.Nil(t, err)

	cheap, expensive := items.Optimize(400)

	require.Equal(t, "mint", cheap.Description)
	require.Equal(t, 11, cheap.Cents)
	require.Equal(t, "taffy", expensive.Description)
	require.Equal(t, 111, expensive.Cents)

	cheap, expensive = items.Optimize(600)

	require.Equal(t, "taffy", cheap.Description)
	require.Equal(t, 111, cheap.Cents)
	require.Equal(t, "chocolate", expensive.Description)
	require.Equal(t, 390, expensive.Cents)

	// exact change
	cheap, expensive = items.Optimize(501)

	require.Equal(t, "taffy", cheap.Description)
	require.Equal(t, 111, cheap.Cents)
	require.Equal(t, "chocolate", expensive.Description)
	require.Equal(t, 390, expensive.Cents)
}

func TestOptimizeNotPossible(t *testing.T) {
	cheap, expensive := GiftList{}.Optimize(400)

	require.Nil(t, cheap)
	require.Nil(t, expensive)
	items, err := NewGiftList(strings.Split(prices, "\n"))
	require.Nil(t, err)

	cheap, expensive = items.Optimize(121)

	require.Nil(t, cheap)
	require.Nil(t, expensive)
}

package wirebson

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncodeScalarField(t *testing.T) {
	buf := make([]byte, 13)
	encodeScalarField(0, buf, "foo", "bar")

	expected := []byte{0x02, 0x66, 0x6f, 0x6f, 0x0, 0x4, 0x0, 0x0, 0x0, 0x62, 0x61, 0x72, 0x0}
	actual := buf

	assert.Equal(t, expected, actual)
}

func TestEncodeField(t *testing.T) {
	buf := make([]byte, 25)

	var i int

	var err error
	i, err = encodeField(i, buf, "foo", "bar")
	require.NoError(t, err)

	actual := buf

	expected := []byte{0x2, 0x66, 0x6f, 0x6f, 0x0, 0x4, 0x0, 0x0, 0x0, 0x62, 0x61, 0x72, 0x0}
	assert.Equal(t, expected, actual)

	i, err = encodeField(i, buf, "foo", int32(1))
	require.NoError(t, err)

	actual = buf

	expected = append(expected, []byte{0x10, 0x66, 0x6f, 0x6f, 0x0, 0x1, 0x0, 0x0, 0x0}...)
	assert.Equal(t, expected, actual)
}

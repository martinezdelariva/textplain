package textplain_test

import (
	"strings"
	"testing"

	"github.com/mailproto/textplain"
	"github.com/stretchr/testify/assert"
)

func TestWrappingInvalidLength(t *testing.T) {
	body := `.stylesheet {
		color: white;
		background-image: url('data:image/png;base64,123456789012345678901234567890');
		font-weight: bold;
		margin: 0px;
	}`

	wrapped := textplain.WordWrap(body, -1)
	assert.Equal(t, body, wrapped)
}

func TestWrappingTrailingWhitespace(t *testing.T) {
	body := "1 23 45\n67\n1234567890 1   "

	wrapped := textplain.WordWrap(body, 13)
	assert.Equal(t, "1 23 45\n67\n1234567890 1 \n", wrapped)

	wrapped = textplain.WordWrap("1234567890"+strings.Repeat(" ", 20), 10)
	assert.Equal(t, "1234567890\n", wrapped)
}

package airplane

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncodeOutputHTMLEscape(t *testing.T) {
	// Note this example uses a JSON-format string, but it's a string nonetheless!
	got, err := encodeOutput(map[string]interface{}{
		"url": "https://airplane.dev?a=b&c=d",
		"another": `new
line`,
	})
	require.NoError(t, err)
	// Go unfortunately alphabetizes the keys:
	want := `{"another":"new\nline","url":"https://airplane.dev?a=b&c=d"}`
	require.Equal(t, want, got)
}

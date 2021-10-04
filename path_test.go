package airplane

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToJS(tt *testing.T) {
	for _, test := range []struct {
		s string
		c []interface{}
	}{
		{"", []interface{}{}},
		{"outputs", []interface{}{"outputs"}},
		{`outputs[""]`, []interface{}{"outputs", ""}},
		{"[0]", []interface{}{0}},
		{`["an output"]`, []interface{}{"an output"}},
		{"outputs.foo", []interface{}{"outputs", "foo"}},
		{"outputs[0]", []interface{}{"outputs", 0}},
		{"outputs[0][1]", []interface{}{"outputs", 0, 1}},
		{`a["hello world"][10].foo.bar`, []interface{}{"a", "hello world", 10, "foo", "bar"}},
		{`["[]"]`, []interface{}{"[]"}},
		{`["]["]`, []interface{}{"]["}},
		{`["\""]`, []interface{}{`"`}},
		{`["\\"]`, []interface{}{`\`}},
		{`[""]`, []interface{}{``}},
		{`a["]\"\\["][10].b`, []interface{}{"a", `]"\[`, 10, "b"}},
		{`a[".b"][10]["[\".c\"]"]`, []interface{}{`a`, `.b`, 10, `[".c"]`}},
	} {
		tt.Run(test.s, func(t *testing.T) {
			require := require.New(t)

			require.Equal(test.s, toJS(test.c...))
		})
	}
}

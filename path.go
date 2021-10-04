package airplane

import (
	"regexp"
	"strconv"
	"strings"
)

var canDot = regexp.MustCompile(`^\w+$`)
// toJS converts a path to a JSONPath-style string representation.
//
// For example:
//   ["foo", 0, "bar", "ba z"] -> `foo[0].bar["ba z"]`
func toJS(path ...interface{}) string {
  var b strings.Builder
  for _, c := range path {
    switch v := c.(type) {
    case string:
      if canDot.MatchString(v) {
        if b.Len() > 0 {
          b.WriteRune('.')
        }
        b.WriteString(v)
      } else {
        b.WriteRune('[')
        b.WriteString(strconv.Quote(v))
        b.WriteRune(']')
      }
    case int:
      b.WriteRune('[')
      b.WriteString(strconv.FormatInt(int64(v), 10))
      b.WriteRune(']')
    }
  }

  return b.String()
}



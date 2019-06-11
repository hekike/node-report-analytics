package classify

import (
	"crypto/sha1"
	"fmt"
	"strings"

	"github.com/hekike/node-report-analytics/pkg/model"
	"github.com/hekike/node-report-analytics/pkg/sanitize"
)

// HashStack returns the hash for a JavaScript stack
func HashStack(stack model.JavaScriptStack) (string, error) {
	var tohash string

	// Sanitize error, remove all numbers
	stack, err := sanitize.Stack(stack)
	if err != nil {
		return "", err
	}

	tohash = fmt.Sprintf(
		"%s%s",
		stack.Message,
		strings.Join(stack.Stack, ","),
	)

	// SHA1
	h := sha1.New()
	h.Write([]byte(tohash))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs), nil
}

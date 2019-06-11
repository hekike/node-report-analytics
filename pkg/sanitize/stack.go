package sanitize

import (
	"regexp"

	"github.com/hekike/node-report-analytics/pkg/model"
)

// Stack sanitize stack
func Stack(stack model.JavaScriptStack) (model.JavaScriptStack, error) {
	reg, err := regexp.Compile("[0-9]+")
	if err != nil {
		return stack, err
	}

	stack.Message = reg.ReplaceAllString(stack.Message, "")

	return stack, nil
}

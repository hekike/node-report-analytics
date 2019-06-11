package sanitize

import (
	"testing"
	"github.com/hekike/node-report-analytics/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := model.JavaScriptStack{
		Message: "My 12Error45",
		Stack: []string{"foo", "bar"},
	}

	result, err := Stack(stack)
	if err != nil {
		panic(err)
	}

	expected := model.JavaScriptStack{
		Message: "My Error",
		Stack: []string{"foo", "bar"},
	}
	assert.Equal(t, expected, result)
}

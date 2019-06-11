package classify

import (
	"testing"
	"github.com/hekike/node-report-analytics/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestHashStack(t *testing.T) {
	stack := model.JavaScriptStack{
		Message: "My Error",
		Stack: []string{"foo", "bar"},
	}
	hash, err := HashStack(stack)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "a1f27373c3e1e0ec9cd897a3a2290ec0a138ce50", hash)
}

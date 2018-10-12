package dry

import (
	"context"
	"errors"
	"testing"
)

func testContext() context.Context {
	return context.Background()
}

func TestError(t *testing.T) {
	t.Run("when the context contains an error", func(t *testing.T) {
		err := errors.New("watnow")
		actual := Error(AddError(testContext(), err))

		if actual != err {
			t.Errorf("expected '%s', got '%s'", err, actual)
		}
	})

	t.Run("when the context does not contain an error", func(t *testing.T) {
		actual := Error(testContext())

		if actual != nil {
			t.Errorf("expected no error, got '%s'", actual)
		}
	})
}

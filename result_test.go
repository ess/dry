package dry

import (
	"context"
	"errors"
	"testing"
)

func successContext() context.Context {
	return testContext()
}

func failureContext() context.Context {
	return context.WithValue(successContext(), dryError, errors.New("failure"))
}

func TestResult_Wrapped(t *testing.T) {
	t.Run("when the result is a failure", func(t *testing.T) {
		ctx := testContext()
		result := Failure(ctx)

		t.Run("it is the wrapped context", func(t *testing.T) {
			if result.Wrapped() != ctx {
				t.Errorf("expected the wrapped context")
			}
		})
	})

	t.Run("when the result is a success", func(t *testing.T) {
		ctx := testContext()
		result := Success(ctx)

		t.Run("it is the wrapped context", func(t *testing.T) {
			if result.Wrapped() != ctx {
				t.Errorf("expected the wrapped context")
			}
		})
	})
}

func TestResult_Value(t *testing.T) {
	k := key(1)
	value := "turned to gold"

	t.Run("when the result is a failure", func(t *testing.T) {
		ctx := context.WithValue(testContext(), k, value)
		result := Failure(ctx)

		t.Run("it is nil", func(t *testing.T) {
			actual := result.Value()

			if actual != nil {
				t.Errorf("expected nil, got %s", actual)
			}
		})

	})

	t.Run("when the result is a success", func(t *testing.T) {
		ctx := context.WithValue(testContext(), k, value)
		result := Success(ctx)

		t.Run("it is the wrapped value", func(t *testing.T) {
			expected := result.Wrapped()
			actual := result.Value()

			if actual != expected {
				t.Errorf("expected '%s', got '%s'", expected, actual)
			}
		})

	})
}

func TestResult_Error(t *testing.T) {
	t.Run("when the result is a failure", func(t *testing.T) {
		ctx := testContext()
		result := Failure(ctx)

		t.Run("it is the wrapped value", func(t *testing.T) {
			expected := result.Wrapped()
			actual := result.Error()

			if actual != expected {
				t.Errorf("expected the context error, got '%s'", actual)
			}
		})

	})

	t.Run("when the result is a success", func(t *testing.T) {
		ctx := testContext()
		result := Success(ctx)

		t.Run("it is nil", func(t *testing.T) {
			actual := result.Error()

			if actual != nil {
				t.Errorf("expected no error, got '%s'", actual)
			}
		})

	})
}

func TestResult_Success(t *testing.T) {
	t.Run("when the result is a failure", func(t *testing.T) {
		ctx := testContext()
		result := Failure(ctx)

		t.Run("it is false", func(t *testing.T) {
			if result.Success() {
				t.Errorf("expected a failure not to be successful")
			}
		})

	})

	t.Run("when the result is a success", func(t *testing.T) {
		ctx := testContext()
		result := Success(ctx)

		t.Run("it is true", func(t *testing.T) {
			if !result.Success() {
				t.Errorf("expected a success to be successful")
			}
		})

	})
}

func TestResult_Failure(t *testing.T) {
	t.Run("when the result is a failure", func(t *testing.T) {
		ctx := testContext()
		result := Failure(ctx)

		t.Run("it is true", func(t *testing.T) {
			if !result.Failure() {
				t.Errorf("expected a failure not to be successful")
			}
		})

	})

	t.Run("when the result is a success", func(t *testing.T) {
		ctx := testContext()
		result := Success(ctx)

		t.Run("it is false", func(t *testing.T) {
			if result.Failure() {
				t.Errorf("expected a success to be successful")
			}
		})

	})
}

func TestSuccess(t *testing.T) {
	ctx := testContext()

	result, isSuccess := Success(ctx).(*success)
	if !isSuccess {
		t.Errorf("expected a success, got %T", result)
	}
}

func TestFailure(t *testing.T) {
	ctx := testContext()

	result, isFailure := Failure(ctx).(*failure)
	if !isFailure {
		t.Errorf("expected a failure, got %T", result)
	}
}

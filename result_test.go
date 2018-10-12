package dry

import (
	"testing"
)

func TestResult_Wrapped(t *testing.T) {
	t.Run("when the result is a failure", func(t *testing.T) {
		value := "holy failure, batman"
		result := Failure(value)

		t.Run("it is the wrapped value", func(t *testing.T) {
			actual := result.Wrapped()

			if actual != value {
				t.Errorf("expected '%s', got '%s'", value, actual)
			}
		})
	})

	t.Run("when the result is a success", func(t *testing.T) {
		value := "what great success"
		result := Success(value)

		t.Run("it is the wrapped value", func(t *testing.T) {
			actual := result.Wrapped()

			if actual != value {
				t.Errorf("expected '%s', got '%s'", value, actual)
			}
		})
	})
}

func TestResult_Value(t *testing.T) {
	value := "my sausages turned to gold"

	t.Run("when the result is a failure", func(t *testing.T) {
		result := Failure(value)

		t.Run("it is nil", func(t *testing.T) {
			actual := result.Value()

			if actual != nil {
				t.Errorf("expected nil, got %s", actual)
			}
		})

	})

	t.Run("when the result is a success", func(t *testing.T) {
		result := Success(value)

		t.Run("it is the wrapped value", func(t *testing.T) {
			actual := result.Value()

			if actual != value {
				t.Errorf("expected '%s', got '%s'", value, actual)
			}
		})

	})
}

func TestResult_Error(t *testing.T) {
	value := "hello i am an error how are you"

	t.Run("when the result is a failure", func(t *testing.T) {
		result := Failure(value)

		t.Run("it is the wrapped value", func(t *testing.T) {
			actual := result.Error()

			if actual != value {
				t.Errorf("expected '%s', got '%s'", value, actual)
			}
		})

	})

	t.Run("when the result is a success", func(t *testing.T) {
		result := Success(value)

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
		result := Failure(nil)

		t.Run("it is false", func(t *testing.T) {
			if result.Success() {
				t.Errorf("expected a failure not to be successful")
			}
		})

	})

	t.Run("when the result is a success", func(t *testing.T) {
		result := Success(nil)

		t.Run("it is true", func(t *testing.T) {
			if !result.Success() {
				t.Errorf("expected a success to be successful")
			}
		})

	})
}

func TestResult_Failure(t *testing.T) {
	t.Run("when the result is a failure", func(t *testing.T) {
		result := Failure(nil)

		t.Run("it is true", func(t *testing.T) {
			if !result.Failure() {
				t.Errorf("expected a failure not to be successful")
			}
		})

	})

	t.Run("when the result is a success", func(t *testing.T) {
		result := Success(nil)

		t.Run("it is false", func(t *testing.T) {
			if result.Failure() {
				t.Errorf("expected a success to be successful")
			}
		})

	})
}

func TestSuccess(t *testing.T) {
	result, isSuccess := Success(nil).(*success)
	if !isSuccess {
		t.Errorf("expected a success, got %T", result)
	}
}

func TestFailure(t *testing.T) {
	result, isFailure := Failure(nil).(*failure)
	if !isFailure {
		t.Errorf("expected a failure, got %T", result)
	}
}

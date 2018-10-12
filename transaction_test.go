package dry

import (
	"reflect"
	"testing"
)

const (
	firstKey string = "first"
	lastKey  string = "last"
	errKey   string = "error"
)

var firstRan = "first ran"
var lastRan = "last ran"
var errFail = "failing step fails"

var first = func(input Value) Result {
	data := input.(map[string]string)
	data[firstKey] = firstRan

	return Success(data)
}

var last = func(input Value) Result {
	data := input.(map[string]string)
	data[lastKey] = lastRan

	return Success(data)
}

var failing = func(input Value) Result {
	data := input.(map[string]string)
	data[errKey] = errFail

	return Failure(data)
}

func TestNewTransaction(t *testing.T) {
	t.Run("when no steps are given", func(t *testing.T) {
		transaction := NewTransaction()

		t.Run("it has no steps", func(t *testing.T) {
			count := len(transaction.steps)

			if count > 0 {
				t.Errorf("expected no steps, got %d steps", count)
			}
		})
	})

	t.Run("when steps are given", func(t *testing.T) {
		steps := []Step{first, last}
		transaction := NewTransaction(steps...)

		t.Run("it contains exactly the given steps in-order", func(t *testing.T) {
			candidates := transaction.steps

			if len(candidates) != len(steps) {
				t.Errorf("the step lists are not the same length")
			}

			for i, step := range steps {
				actual := reflect.ValueOf(candidates[i]).Pointer()
				expected := reflect.ValueOf(step).Pointer()

				if actual != expected {
					t.Errorf("the steps are not in the proper order")
				}
			}
		})
	})
}

func TestTransact(t *testing.T) {
	t.Run("when no steps are given", func(t *testing.T) {
		ctx := make(map[string]string)
		result := Transact(ctx)

		t.Run("it is a success", func(t *testing.T) {
			if !result.Success() {
				t.Errorf("expected a success")
			}
		})

		t.Run("it does not alter its input", func(t *testing.T) {
			actual := result.Wrapped()

			if !reflect.DeepEqual(actual, ctx) {
				t.Errorf("expected '%s', got '%s'", ctx, actual)
			}
		})
	})

	t.Run("when all steps succeed", func(t *testing.T) {
		ctx := make(map[string]string)
		result := Transact(ctx, first, last)

		t.Run("it is a success", func(t *testing.T) {
			if !result.Success() {
				t.Errorf("expected a success")
			}
		})

		t.Run("it executes all of the steps", func(t *testing.T) {
			firstValue := result.Wrapped().(map[string]string)[firstKey]
			lastValue := result.Wrapped().(map[string]string)[lastKey]

			if firstValue != firstRan {
				t.Errorf("first did not run!")
			}

			if lastValue != lastRan {
				t.Errorf("last did not run!")
			}
		})
	})

	t.Run("when a step fails", func(t *testing.T) {
		ctx := make(map[string]string)
		result := Transact(ctx, first, failing, last)

		t.Run("it is a failure", func(t *testing.T) {
			if !result.Failure() {
				t.Errorf("expected a failure")
			}
		})

		t.Run("it runs steps before the failure", func(t *testing.T) {
			firstValue := result.Wrapped().(map[string]string)[firstKey]

			if firstValue != firstRan {
				t.Errorf("first did not run!")
			}
		})

		t.Run("it skips steps after the failure", func(t *testing.T) {
			lastValue := result.Wrapped().(map[string]string)[lastKey]

			if lastValue == lastRan {
				t.Errorf("Expected last not to run")
			}
		})
	})
}

func TestTransaction_Call(t *testing.T) {
	t.Run("when no steps are given", func(t *testing.T) {
		ctx := make(map[string]string)
		result := NewTransaction().Call(ctx)

		t.Run("it is a success", func(t *testing.T) {
			if !result.Success() {
				t.Errorf("expected a success")
			}
		})

		t.Run("it does not alter its input", func(t *testing.T) {
			actual := result.Wrapped()

			if !reflect.DeepEqual(actual, ctx) {
				t.Errorf("expected '%s', got '%s'", ctx, actual)
			}
		})
	})

	t.Run("when all steps succeed", func(t *testing.T) {
		ctx := make(map[string]string)
		result := NewTransaction(first, last).Call(ctx)

		t.Run("it is a success", func(t *testing.T) {
			if !result.Success() {
				t.Errorf("expected a success")
			}
		})

		t.Run("it executes all of the steps", func(t *testing.T) {
			firstValue := result.Value().(map[string]string)[firstKey]
			lastValue := result.Value().(map[string]string)[lastKey]

			if firstValue != firstRan {
				t.Errorf("first did not run!")
			}

			if lastValue != lastRan {
				t.Errorf("last did not run!")
			}
		})
	})

	t.Run("when a step fails", func(t *testing.T) {
		ctx := make(map[string]string)
		result := NewTransaction(first, failing, last).Call(ctx)

		t.Run("it is a failure", func(t *testing.T) {
			if !result.Failure() {
				t.Errorf("expected a failure")
			}
		})

		t.Run("it runs steps before the failure", func(t *testing.T) {
			firstValue := result.Wrapped().(map[string]string)[firstKey]

			if firstValue != firstRan {
				t.Errorf("first did not run!")
			}
		})

		t.Run("it skips steps after the failure", func(t *testing.T) {
			lastValue := result.Wrapped().(map[string]string)[lastKey]

			if lastValue == lastRan {
				t.Errorf("Expected last not to run")
			}
		})
	})
}

func TestTransaction_Step(t *testing.T) {
	steps := []Step{first, last}
	extra := func(input Value) Result {
		data := input.(map[string]string)
		data["extra"] = "read all about it"

		return Success(data)
	}

	transaction := NewTransaction(steps...)
	result := transaction.Step(extra)

	t.Run("it adds the specified step to the transaction", func(t *testing.T) {
		count := len(transaction.steps)

		if count == len(steps) {
			t.Errorf("Expected %d steps, got %d", len(steps)+1, count)
		}

		actual := reflect.
			ValueOf(transaction.steps[len(transaction.steps)-1]).
			Pointer()

		expected := reflect.ValueOf(extra).Pointer()

		if actual != expected {
			t.Errorf("expected an extra step")
		}
	})

	t.Run("it returns the transaction itself", func(t *testing.T) {
		if result != transaction {
			t.Errorf("expected the transaction, but got something else")
		}
	})
}

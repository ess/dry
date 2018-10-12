package dry

// Step is a function that can be used as a step in a Transaction.
//
// Any given step is passed a Value, and it must return a Result.
type Step func(Value) Result

// Transaction is a model that describes a multi-step process that can fail
// during any given step.
type Transaction struct {
	steps []Step
}

// NewTransaction returns a new Transaction.
//
// If an optional list of steps is provided, the resulting transaction includes
// those steps.
//
// Otherwise, the transaction is a no-op and the Step method should be used to
// add functionality.
func NewTransaction(steps ...Step) *Transaction {
	if steps == nil {
		steps = make([]Step, 0)
	}

	return &Transaction{steps: steps}
}

// Transact takes a Value and a list of Steps, performs the transaction
// described by those arguments, and returns the final Result of those
// sequential operations.
func Transact(input Value, steps ...Step) Result {
	return NewTransaction(steps...).Call(input)
}

// Step takes a step function and adds it to the list of steps to perform when
// the associated transaction is called.
//
// For the sake of those who prefer method chains to individual calls, the
// modified transaction is returned.
func (transaction *Transaction) Step(step Step) *Transaction {
	transaction.steps = append(transaction.steps, step)

	return transaction
}

// Call takes a Value, which is then propagated through the list of steps
// associated with the transaction.
//
// The result of each step is passed as input to its subsequent step, and the
// result of the final step is returned.
//
// If any step results in a Failure, that failure is immediately returned
// and steps after the failing step are skipped.
func (transaction *Transaction) Call(value Value) Result {
	result := Success(value)

	for _, step := range transaction.steps {
		result = step(value)

		if result.Failure() {
			return result
		}

		value = result.Wrapped()
	}

	return result
}

/*
Copyright 2018 Dennis Walters

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

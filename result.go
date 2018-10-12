package dry

// Value is an interface that describes literally any value at all.
type Value interface{}

// Result is (not really) a monad that is used to model the result of an
// operation. It can be either a Success or a Failure that wraps a value.
type Result interface {
	// Wrapped returns the value that is wrapped by a Result.
	Wrapped() Value

	// Success returns a boolean. If the result is a success, it is true.
	// Otherwise, it is false.
	Success() bool

	// Failure returns a boolean. If the result is a failure, it is true.
	// Otherwise, it is false.
	Failure() bool

	// Value returns the value that is wrapped by a Success result.
	Value() Value

	// Error returns the value that is wrapped by a Failure result.
	Error() Value
}

type baseResult struct {
	value Value
}

func (result *baseResult) Wrapped() Value {
	return result.value
}

func (result *baseResult) Success() bool {
	return false
}

func (result *baseResult) Failure() bool {
	return false
}

func (result *baseResult) Value() Value {
	return nil
}

func (result *baseResult) Error() Value {
	return nil
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

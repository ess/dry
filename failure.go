package dry

// Failure wraps a context in a failure Result
func Failure(value Value) Result {
	return &failure{
		&baseResult{value: value},
	}
}

type failure struct {
	*baseResult
}

func (result *failure) Failure() bool {
	return true
}

func (result *failure) Error() Value {
	return result.Wrapped()
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

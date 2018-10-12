package main

import (
	"fmt"

	"github.com/ess/dry"
)

func main() {
	noisyInc := dry.NewTransaction()
	noisyInc.Step(show)
	noisyInc.Step(increment)

	result := dry.NewTransaction(noisyInc.Call, noisyInc.Call).Call(120)

	if result.Failure() {
		panic(result.Error())
	}

	total := result.Value().(int)
	fmt.Println("Final total:", total)
}

func increment(data dry.Value) dry.Result {
	s, ok := data.(int)
	if !ok {
		return dry.Failure(fmt.Errorf("value isn't an integer"))
	}

	if s%2 == 0 {
		return dry.Failure(fmt.Errorf("i can't even"))
	}

	return dry.Success(s + 1)
}

func show(input dry.Value) dry.Result {
	fmt.Println("current value:", input)

	return dry.Success(input)
}

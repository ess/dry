# dry #

[![Build Status](https://travis-ci.org/ess/dry.svg?branch=master)](https://travis-ci.org/ess/dry)
[![Go Report Card](https://goreportcard.com/badge/github.com/ess/dry)](https://goreportcard.com/report/github.com/ess/dry)
[![Documentation](https://godoc.org/github.com/ess/dry?status.svg)](http://godoc.org/github.com/ess/dry)

A minimal set of quasi-functional programming constructs for Go, inspired by [dry-rb](https://dry-rb.org).


## Fauxnads ##

Rather than actaul monads, we're going to use fauxnads ... things that quack like monads. Currently, the following fauxnads are available: `Result` (success or failure).

### Result ###

A `Result` can be either a success or a failure, and in this package, it always wraps a `dry.Value` (`interface{}`). *Let's save the argument about how this is bad and I'm bad for doing this until after generics are a thing, hmmkay?*

To create a `Result`, pass a `Value` to either `dry.Success()` or `dry.Failure()`, depending on your needs:

```go
rightOn := dry.Success("heck yeah")
oNoes := dry.Failure("you broke it")

```

Also, in the event that you're handed any old `Result` and you want to check for success/failure, you can use the `Success()` and `Failure()` methods on the resuit:

```go
// if successful, let's do some stuff
if result.Success() {
  doStuff(result.Value())
}

// if failure, let's do some different stuff
if result.Failure() {
  doOtherStuff(result.Error())
}
```

## Transactions ##

The primary reason behind this whole package, really, is that I rather like [dry-transaction](https://dry-rb.org/gems/dry-transaction) when dealing with Ruby endeavors, and I greatly miss it when doing up my Go stuff. There's already at least one Railway-Oriented Programming package for Go ([rop](https://github.com/dc0d/rop)), and chances are pretty good that it's close to a real implementation than `dry.Transaction`.

That said, I never saw a wheel I didn't want to reinvent, so here we are!

The short description of a transaction that I like to go with is "a multi-step process that can fail at any point, stop execution, and allow for recovery".

To that end, a `Step` is a very specific type of function, in as much as it's any fuction that takes a `Value` and returns a `Result`. Following from this, a `Transaction` is a collection of `Step`s that can be `Call`ed with an input `Value`, and the content of the returned `Result` from one `Step` is passed in as input to the next `Step`. In the fine tradition of providing complicated examples to perform trivial tasks, here's a quick example that increments an integer:

```go
package main

import (
	"fmt"

	"github.com/ess/dry"
)

func main() {
	transaction := dry.NewTransaction(
		show,
		increment,
	)

	result := transaction.Call(120)

	if result.Failure() {
		panic(result.Error())
	}

	total := result.Value(myValue).(int)
	fmt.Println("final total:", total)
}

func increment(data dry.Value) dry.Result {
	s, ok := data.(int)
	if !ok {
		return dry.Failure(fmt.Errorf("value isn't an integer"))
	}

	return dry.Success(s + 1)
}

func show(input dry.Value) dry.Result {
	fmt.Println("current value:", input)

	return dry.Success(input)
}
```

Now, that's a fine example right there of something that we certainly could have done in like three lines of code, but didn't. Here's the output when we run it:

```
current value: 120
final total: 121
```

We could have also build the transaction manually like so:

```go
	transaction := dry.NewTransaction()
	transaction.Step(show)
	transaction.Step(increment)
```

If we wanted to do the whole thing twice, we could have done it like this:

```go
	transaction := dry.NewTransaction()
	transaction.Step(show)
	transaction.Step(increment)
	transaction.Step(show)
	transaction.Step(increment)
```

Here's the new output:

```
current value: 120
current value: 121
Final total: 122
```

If you do it that way, though, you're missing out on all the fun, because one transaction can be used as a step in another transaction. Let's give that a shot:

```go
	trivial := dry.NewTransaction()
	trivial.Step(show)
	trivial.Step(increment)

	transaction := dry.NewTransaction(trivial.Call, trivial.Call)
```

So, what happens when there's a failure along the way? Let's be mean to ourselves and inject a problem into one of the steps:

```go
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
```

There we go. Incrementing the number will now fail if the number is even. Here's the output:

```
current value: 120
panic: i can't even

goroutine 1 [running]:
main.main()
	/path/to/dry/examples/main.go:26 +0x563
```

## History ##

* v1.0.0 - Initial Release

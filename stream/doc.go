/*
Package Stream defines a set of operations that provide lazy operations on
a Stream of data values. This package is still experimental, please use at
your own risk.

A Stream is an abstraction that produces a single value at a time. A Stream
can produce infinite values, or just a finite set of values. A simple example
of using a Stream is:

	ForEach(Take(Iota[int](), 10), func(value int) {
		// Do something with 'value'
	})

Iota() produces a Stream, which is consumed by Take(), which produces a Stream
that is consumed by ForEach().

Streams are synchronous by default. A segment of a Stream can be made
asynchronous by using the Async() adapter. All operations inside this adapter
will execute asynchronously from the main thread, and their results will be
fed through a channel back to the main thread.

The design of streams is heavily influenced by C#'s LINQ and Java's Stream API.
The current design is not considered stable yet, so newer versions may cause
the API to change significantly.
*/
package stream

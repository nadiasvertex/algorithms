/*
Package stream defines a set of operations that provide lazy operations on
a stream of data values. This package is still experimental, please use at
your own risk.

A stream is an abstraction that produces a single value at a time. A stream
can produce infinite values, or just a finite set of values. A simple example
of using a stream is:

	ForEach(Take(Iota[int](), 10), func(value int) {
		// Do something with 'value'
	})

Iota() produces a stream, which is consumed by Take(), which produces a stream
that is consumed by ForEach().

Streams are synchronous by default. A segment of a stream can be made
asynchronous by using the Async() adapter. All operations inside this adapter
will execute asynchronously from the main thread, and their results will be
fed through a channel back to the main thread.

The design of streams is heavily influenced by C#'s LINQ and Java's stream API.
The current design is not considered stable yet, so newer versions may cause
the API to change significantly.
*/
package stream

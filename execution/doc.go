/*
Package execution is a reimplementation of the std::execution proposal
at https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2022/p2300r5.html for
Go. Obviously the two languages are very different and some mechanisms from the
paper either make no sense or are impossible to code in Go.

However, an attempt has been made to faithfully reproduce the core semantics
of senders, receivers, and schedulers along with cancellation via stop tokens
and error propagation.

In general, operations can be composed in the same way as operations from the
[stream] package. However, all operations in the execution module run
asynchronously. Additionally, the execution module is focused on enabling
asynchronous computations, not so much on the algorithms we might want to
enable on top of them.

For example, a simple way of creating a sender is:

	s := Then(Just(10), func(v int) int {
		return v * v
	})

This says, "take just the value 10, then execute this function which will compute
the square of its argument and return that." This sender is lazy. It will do
nothing until it is scheduled. The simplest way to schedule the sender is to
use SyncWait().

	v, err := SyncWait(s)

The result value from the computations will be placed in 'v', and any error from
the computation will be placed in 'err'. Note that, if an error happens at
any point in the chain, all succeeding operations are skipped and the error is
forwarded to the operation chain's exit point.

This module has basic tests, and is usable. However, the API may not be stable
so, use it at your own risk.
*/
package execution

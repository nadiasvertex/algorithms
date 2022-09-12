/*
Package parallel implements parallel versions of the algorithms. The parallel
versions are identical to the serial versions, except that they are performed
in parallel.

Due to coordination overhead, parallel versions of the algorithms are only
faster than the serial versions when you are dealing with large slices. Currently,
"large" means hundreds or thousands of elements. Alternatively, if you have a
small collection of large items or the operations are very slow for each item
you may also notice improvement.

I am still working on best practices for using the parallelism module and for
composing the algorithms efficiently. In particular, I am working on ways to
allow a parallel algorithm to send back chunks of completed work in a way that
would let other parallel algorithms begin their job. However, the one-shot
parallel  versions are unlikely to go away, so you shouldn't consider their API
particularly unstable.
*/
package parallel

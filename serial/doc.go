/*
Package serial implements a number of algorithms from the C++ Standard Template
Library. While that library has an "iterator" abstraction, this concept has
proven extremely hard to efficiently replicate in Go. Consequently, all the
algorithms are implemented on top of slices. Some additional functions have
been introduced to deal with the fact that Go does not have function overloading
or default arguments.

The goal is to reimplement all the C++ STL algorithms in this library. Currently,
most of them are available. A few of the more complex algorithms, like
nth_element are still pending.

Most of the algorithms have basic tests ensuring that they function correctly.
Eventually I would like to have 100% code line test coverage of all the serial
algorithms, and some tests for bad edge cases and performance cliffs.

If you are interested in helping, please feel free to implement an algorithm,
a test for the algorithm, and open a pull request against this repo. Or if you
just want to write tests, that would also be appreciated!
*/
package serial

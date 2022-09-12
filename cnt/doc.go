/*
Package cnt contains container implementations. Currently, it contains an
implementation of Optional. It will eventually contain an implementation of
a flat set and a flat map, and potentially also a vector. While Go does provide
a hash map and array, having a container with methods that perform operations on
the storage is very convenient.

# Optional

As a note on the optional type implementation, I ended up choosing to require
a pointer to the optional. This is primarily so that 'nil' can easily represent
an empty optional.

The first implementation of Optional used an interface and two implementation
classes. While that was very convenient, it wasn't more memory efficient since
an array of Optionals would still just be a slice. Also, that meant that a user
had to explicitly initialize an Optional variable since it would have no way
of knowing what concrete type is associated with a nil interface of Optional
type. It is a little annoying to have to remember to make all optionals a
pointer to an optional type, but it seems to be the best compromise at this
time.

You might wonder why a user would prefer Optional[T] over just using *T. First,
getting a pointer to an arbitrary value can be verbose and inconvenient.
Also, performing nil checking all over can also be verbose an inconvenient.
Finally, the optional type makes it clear a function might not return a value,
or that it can accept not having a value.

As an added convenience, the Optional type has a number of methods that make
working with optional values less verbose and error-prone.
*/
package cnt

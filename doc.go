/*
This package contains a super fast and collection that stores a state in bits.
There are two implementations you can find in the package:
  - Static
  - Dynamic

Both of them can help to minify the memory usage and the time spent for Set data-structure.
If you need to save a uint in set, then this implementation will be much more efficient than the usual Set.
The Set key is storing in 1 bit in an array or map.

This data structure is useful to share the flags in JSON or store them in a database.
There's an article that shares the motivation to have this implementation: https://aohorodnyk.com/post/2021-01-03-binary-flags/
*/
package binflags

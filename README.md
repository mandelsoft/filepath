
Package `filepath` implements utility routines for manipulating filename paths
in a way compatible with the target operating system-defined file paths.
It is a modification of the original GO package `path/filepath` solving
severe errors in handling symbolic links.

The original package defines a function `Clean` that formally normalizes
a file path by eliminating `..` and `.` entries. This is done WITHOUT
observing the actual file system. Although this is no problem for
the function itself, because it is designed to do so, it becomes a severe
problem for the whole package, because nearly all functions internally use
`Clean` to clean the path. As a consequence even functions like `Join` deliver
corrupted invalid results for valid inputs if the path incorporates
symbolic links to directories. Especially `EvalSymlinks` cannot be used
to evaluate paths to existing files, because `Clean` is internally used to
normalize content of symbolic links.

This package provides a set of functions that do not hamper the meaning
of path names keeping the rest of the original specification as far as
possible. Additionally some new functions (like `Canonical`) or alternate
versions of existing functions (like `Split2`) are provided
that offer a more useful specification than the original one.

### NOTE:
So far, the windows operating system is NOT tested.

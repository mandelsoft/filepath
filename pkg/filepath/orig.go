package filepath

import (
	orig "path/filepath"
)

func IsAbs(path string) bool {
	return orig.IsAbs(path)
}

func ToSlash(path string) string {
	return orig.ToSlash(path)
}

func FromSlash(path string) string {
	return orig.FromSlash(path)
}

func SplitList(path string) []string {
	return orig.SplitList(path)
}

func Split(path string) (string, string) {
	return orig.Split(path)
}

func Clean(path string) string {
	return orig.Clean(path)
}

func Ext(path string) string {
	return orig.Ext(path)
}

func VolumeName(path string) string {
	return orig.VolumeName(path)
}

func Rel(basepath, targpath string) (string, error) {
	base, err := Canonical(basepath, false)
	if err != nil {
		return "", err
	}
	targpath, err = Canonical(targpath, false)
	if err != nil {
		return "", err
	}
	targpath, err = EvalSymlinks(targpath)
	if err != nil {
		return "", err
	}
	return orig.Rel(base, targpath)
}

// SkipDir is used as a return value from WalkFuncs to indicate that
// the directory named in the call is to be skipped. It is not returned
// as an error by any function.
var SkipDir = orig.SkipDir

// SkipAll is used as a return value from WalkFuncs to indicate that
// all remaining files and directories are to be skipped. It is not returned
// as an error by any function.
// var SkipAll = orig.SkipAll

// WalkFunc is the type of the function called for each file or directory
// visited by Walk. The path argument contains the argument to Walk as a
// prefix; that is, if Walk is called with "dir", which is a directory
// containing the file "a", the walk function will be called with argument
// "dir/a". The info argument is the os.FileInfo for the named path.
//
// If there was a problem walking to the file or directory named by path, the
// incoming error will describe the problem and the function can decide how
// to handle that error (and Walk will not descend into that directory). In the
// case of an error, the info argument will be nil. If an error is returned,
// processing stops. The sole exception is when the function returns the special
// value SkipDir. If the function returns SkipDir when invoked on a directory,
// Walk skips the directory's contents entirely. If the function returns SkipDir
// when invoked on a non-directory file, Walk skips the remaining files in the
// containing directory.
type WalkFunc orig.WalkFunc

func Walk(root string, walkFn WalkFunc) error {
	return orig.Walk(root, orig.WalkFunc(walkFn))
}

var ErrBadPattern = orig.ErrBadPattern

func Match(pattern, name string) (matched bool, err error) {
	return orig.Match(pattern, name)
}

func Glob(pattern string) (matches []string, err error) {
	return orig.Glob(pattern)
}

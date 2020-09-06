package filepath

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	orig "path/filepath"
)

var _ = Describe("filepath package", func() {
	Context("function Dir", func() {
		testUnchanged := func(path, result string) {
			Expect(Dir(FromSlash(path))).To(Equal(FromSlash(result)))
			Expect(Dir(FromSlash(path))).To(Equal(orig.Dir(path)))
		}
		testNew := func(path, result string) {
			Expect(Dir(FromSlash(path))).To(Equal(FromSlash(result)))
		}

		//
		// relative
		//
		// plain
		It("empty", func() {
			testUnchanged("", ".")
		})
		It("dot", func() {
			testUnchanged(".", ".")
		})
		It("parent", func() {
			testUnchanged("..", ".")
		})
		It("double parent", func() {
			testUnchanged("../..", "..")
		})
		It("parent dot", func() {
			testUnchanged("../.", "..")
		})

		// nested
		It("single", func() {
			testUnchanged("root", ".")
		})
		It("nested dot", func() {
			testUnchanged("root/.", "root")
		})
		It("nested parent", func() {
			testNew("root/..", "root")
		})
		It("nested double parent", func() {
			testNew("root/../..", "root/..")
		})
		It("nested parent dot", func() {
			testNew("root/../.", "root/..")
		})

		It("double nested", func() {
			testUnchanged("root/nested", "root")
		})

		// tagged plain
		It("tagged dot", func() {
			testUnchanged("./", ".")
		})
		It("tagged parent", func() {
			testUnchanged("../", "..")
		})
		It("tagged double parent", func() {
			testNew("../../", "../..")
		})
		It("tagged parent dot", func() {
			testNew(".././", "../.")
		})

		It("empty dir", func() {
			testUnchanged("a/b/c/", "a/b/c")
		})

		// tagged nested
		It("tagged", func() {
			testUnchanged("root/", "root")
		})
		It("tagged nested dot", func() {
			testNew("root/./", "root/.")
		})
		It("tagged nested parent", func() {
			testNew("root/../", "root/..")
		})
		It("tagged nested double parent", func() {
			testNew("root/../../", "root/../..")
		})
		It("tagged nested parent dot", func() {
			testNew("root/.././", "root/../.")
		})

		It("tagged double nested", func() {
			testUnchanged("root/nested/", "root/nested")
		})

		//
		// from root
		//
		// plain
		It("root", func() {
			testUnchanged("/", "/")
		})
		It("root dot", func() {
			testUnchanged("/.", "/")
		})
		It("root parent", func() {
			testUnchanged("/..", "/")
		})
		It("root double parent", func() {
			testNew("/../..", "/..")
		})
		It("root parent dot", func() {
			testNew("/../.", "/..")
		})

		// nested
		It("top-level", func() {
			testUnchanged("/root", "/")
		})
		It("root nested dot", func() {
			testUnchanged("/root/.", "/root")
		})
		It("root nested parent", func() {
			testNew("/root/..", "/root")
		})
		It("root nested double parent", func() {
			testNew("/root/../..", "/root/..")
		})
		It("root nested parent dot", func() {
			testNew("/root/../.", "/root/..")
		})

		It("root double nested", func() {
			testUnchanged("/root/nested", "/root")
		})

		// tagged plain
		It("tagged root", func() {
			testUnchanged("//", "/")
		})
		It("tagged root dot", func() {
			testNew("/./", "/.")
		})
		It("tagged root parent", func() {
			testNew("/../", "/..")
		})
		It("tagged root double parent", func() {
			testNew("/../../", "/../..")
		})
		It("tagged root parent dot", func() {
			testNew("/.././", "/../.")
		})

		// tagged nested
		It("tagged top-level", func() {
			testUnchanged("/root/", "/root")
		})
		It("tagged root nested dot", func() {
			testNew("/root/./", "/root/.")
		})
		It("tagged root nested parent", func() {
			testNew("/root/../", "/root/..")
		})
		It("tagged root nested double parent", func() {
			testNew("/root/../../", "/root/../..")
		})
		It("tagged root nested parent dot", func() {
			testNew("/root/.././", "/root/../.")
		})

		It("tagged root double nested", func() {
			testUnchanged("/root/nested/", "/root/nested")
		})
	})

	Context("function Dir2", func() {
		testUnchanged := func(path, result string) {
			Expect(Dir2(FromSlash(path))).To(Equal(FromSlash(result)))
			Expect(Dir2(FromSlash(path))).To(Equal(orig.Dir(path)))
		}
		testNew := func(path, result string) {
			Expect(Dir2(FromSlash(path))).To(Equal(FromSlash(result)))
			Expect(Trim(Join(Dir2(FromSlash(path)), Base(FromSlash(path))))).To(Equal(Trim(FromSlash(path))))
		}

		//
		// relative
		//
		// plain
		It("empty", func() {
			testUnchanged("", ".")
		})
		It("dot", func() {
			testUnchanged(".", ".")
		})
		It("parent", func() {
			testUnchanged("..", ".")
		})
		It("double parent", func() {
			testUnchanged("../..", "..")
		})
		It("parent dot", func() {
			testUnchanged("../.", "..")
		})

		// nested
		It("single", func() {
			testUnchanged("root", ".")
		})
		It("nested dot", func() {
			testUnchanged("root/.", "root")
		})
		It("nested parent", func() {
			testNew("root/..", "root")
		})
		It("nested double parent", func() {
			testNew("root/../..", "root/..")
		})
		It("nested parent dot", func() {
			testNew("root/../.", "root/..")
		})

		It("double nested", func() {
			testUnchanged("root/nested", "root")
		})

		// tagged plain
		It("tagged dot", func() {
			testUnchanged("./", ".")
		})
		It("tagged parent", func() {
			testNew("../", ".")
		})
		It("tagged double parent", func() {
			testNew("../../", "..")
		})
		It("tagged parent dot", func() {
			testNew(".././", "..")
		})

		It("empty dir", func() {
			testNew("a/b/c/", "a/b")
		})

		// tagged nested
		It("tagged", func() {
			testNew("root/", ".")
		})
		It("tagged nested dot", func() {
			testNew("root/./", "root")
		})
		It("tagged nested parent", func() {
			testNew("root/../", "root")
		})
		It("tagged nested double parent", func() {
			testNew("root/../../", "root/..")
		})
		It("tagged nested parent dot", func() {
			testNew("root/.././", "root/..")
		})

		It("tagged double nested", func() {
			testNew("root/nested/", "root")
		})

		//
		// from root
		//
		// plain
		It("root", func() {
			testUnchanged("/", "/")
		})
		It("root dot", func() {
			testUnchanged("/.", "/")
		})
		It("root parent", func() {
			testUnchanged("/..", "/")
		})
		It("root double parent", func() {
			testNew("/../..", "/..")
		})
		It("root parent dot", func() {
			testNew("/../.", "/..")
		})

		// nested
		It("top-level", func() {
			testUnchanged("/root", "/")
		})
		It("root nested dot", func() {
			testUnchanged("/root/.", "/root")
		})
		It("root nested parent", func() {
			testNew("/root/..", "/root")
		})
		It("root nested double parent", func() {
			testNew("/root/../..", "/root/..")
		})
		It("root nested parent dot", func() {
			testNew("/root/../.", "/root/..")
		})

		It("root double nested", func() {
			testUnchanged("/root/nested", "/root")
		})

		// tagged plain
		It("tagged root", func() {
			testUnchanged("//", "/")
		})
		It("tagged root dot", func() {
			testNew("/./", "/")
		})
		It("tagged root parent", func() {
			testNew("/../", "/")
		})
		It("tagged root double parent", func() {
			testNew("/../../", "/..")
		})
		It("tagged root parent dot", func() {
			testNew("/.././", "/..")
		})

		// tagged nested
		It("tagged single", func() {
			testNew("/root/", "/")
		})
		It("tagged root nested dot", func() {
			testNew("/root/./", "/root")
		})
		It("tagged root nested parent", func() {
			testNew("/root/../", "/root")
		})
		It("tagged root nested double parent", func() {
			testNew("/root/../../", "/root/..")
		})
		It("tagged root nested parent dot", func() {
			testNew("/root/.././", "/root/..")
		})

		It("tagged root double nested", func() {
			testNew("/root/nested/", "/root")
		})

		It("combinations", func() {
			parts := []string{"file", "/", ".", ".."}
			type check func([]string)

			level := func(f check) check {
				return func(segments []string) {
					for _, s := range parts {
						next := append(segments, s)
						f(next)
					}
				}
			}
			level(level(level(level(func(segments []string) {
				path := Join(segments...)
				//fmt.Printf("trimming %q\n", path)
				Expect(Trim(Join(Dir2(path), Base(path)))).To(Equal(Trim(path)))
			}))))([]string{})

		})
	})
	Context("function Base", func() {
		testUnchanged := func(path, result string) {
			Expect(Base(FromSlash(path))).To(Equal(FromSlash(result)))
			Expect(Base(FromSlash(path))).To(Equal(orig.Base(path)))
		}
		testNew := func(path, result string) {
			Expect(Base(FromSlash(path))).To(Equal(FromSlash(result)))
		}

		//
		// relative
		//
		// plain
		It("empty", func() {
			testUnchanged("", ".")
		})
		It("dot", func() {
			testUnchanged(".", ".")
		})
		It("parent", func() {
			testUnchanged("..", "..")
		})
		It("double parent", func() {
			testUnchanged("../..", "..")
		})
		It("parent dot", func() {
			testUnchanged("../.", ".")
		})

		// nested
		It("single", func() {
			testUnchanged("root", "root")
		})
		It("nested dot", func() {
			testUnchanged("root/.", ".")
		})
		It("nested parent", func() {
			testNew("root/..", "..")
		})
		It("nested double parent", func() {
			testNew("root/../..", "..")
		})
		It("nested parent dot", func() {
			testNew("root/../.", ".")
		})

		It("double nested", func() {
			testUnchanged("root/nested", "nested")
		})

		// tagged plain
		It("tagged dot", func() {
			testUnchanged("./", ".")
		})
		It("tagged parent", func() {
			testNew("../", "..")
		})
		It("tagged double parent", func() {
			testNew("../../", "..")
		})
		It("tagged parent dot", func() {
			testNew(".././", ".")
		})

		// tagged nested
		It("tagged", func() {
			testNew("root/", "root")
		})
		It("tagged nested dot", func() {
			testNew("root/./", ".")
		})
		It("tagged nested parent", func() {
			testNew("root/../", "..")
		})
		It("tagged nested double parent", func() {
			testNew("root/../../", "..")
		})
		It("tagged nested parent dot", func() {
			testNew("root/.././", ".")
		})

		It("tagged double nested", func() {
			testNew("root/nested/", "nested")
		})

		//
		// from root
		//
		// plain
		It("root", func() {
			testUnchanged("/", "/")
		})
		It("root dot", func() {
			testUnchanged("/.", ".")
		})
		It("root parent", func() {
			testUnchanged("/..", "..")
		})
		It("root double parent", func() {
			testUnchanged("/../..", "..")
		})
		It("root parent dot", func() {
			testUnchanged("/../.", ".")
		})

		// nested
		It("top-level", func() {
			testUnchanged("/root", "root")
		})
		It("root nested dot", func() {
			testUnchanged("/root/.", ".")
		})
		It("root nested parent", func() {
			testUnchanged("/root/..", "..")
		})
		It("root nested double parent", func() {
			testUnchanged("/root/../..", "..")
		})
		It("root nested parent dot", func() {
			testUnchanged("/root/../.", ".")
		})

		It("root double nested", func() {
			testUnchanged("/root/nested", "nested")
		})

		// tagged plain
		It("tagged root", func() {
			testUnchanged("//", "/")
		})
		It("tagged root dot", func() {
			testUnchanged("/./", ".")
		})
		It("tagged root parent", func() {
			testUnchanged("/../", "..")
		})
		It("tagged root double parent", func() {
			testUnchanged("/../../", "..")
		})
		It("tagged root parent dot", func() {
			testUnchanged("/.././", ".")
		})

		// tagged nested
		It("tagged single", func() {
			testUnchanged("/root/", "root")
		})
		It("tagged root nested dot", func() {
			testUnchanged("/root/./", ".")
		})
		It("tagged root nested parent", func() {
			testUnchanged("/root/../", "..")
		})
		It("tagged root nested double parent", func() {
			testUnchanged("/root/../../", "..")
		})
		It("tagged root nested parent dot", func() {
			testUnchanged("/root/.././", ".")
		})

		It("tagged root double nested", func() {
			testUnchanged("/root/nested/", "nested")
		})
	})

	Context("function Trim", func() {
		testNew := func(path, result string) {
			Expect(Trim(FromSlash(path))).To(Equal(FromSlash(result)))
		}

		It("empty", func() {
			testNew("", "")
		})
		It("dot", func() {
			testNew(".", ".")
		})
		It("parent", func() {
			testNew("..", "..")
		})
		It("root", func() {
			testNew("/", "/")
		})
		It("tagged root", func() {
			testNew("//", "/")
		})
		It("tagged dot", func() {
			testNew("./", ".")
		})
		It("tagged parent", func() {
			testNew("../", "..")
		})
		It("nested", func() {
			testNew("/root/", "/root")
		})
		It("tagged nested", func() {
			testNew("/root/", "/root")
		})
		It("double tagged nested", func() {
			testNew("/root//", "/root")
		})

		It("relative", func() {
			testNew("nested", "nested")
		})
		It("tagged nested", func() {
			testNew("nested/", "nested")
		})
		It("double tagged nested", func() {
			testNew("nested//", "nested")
		})
		It("double slashes", func() {
			testNew("//dir//nested//", "/dir/nested")
		})

		It("eliminated intermediate dot segments in absolute path", func() {
			testNew("/./a/.", "/a")
			testNew("/../a/.", "/../a")
			testNew("/.", "/")
			testNew("//.", "/")
			testNew("//.//", "/")
		})
		It("eliminated intermediate dot segments in relative path", func() {
			testNew("././a/.", "a")
			testNew("./../a/.", "../a")
		})
		It("eliminated intermediate dot segments but not a single one", func() {
			testNew(".", ".")
		})
	})

	Context("function Split2", func() {
		testUnchanged := func(path, dir, base string) {
			d, b := Split2(FromSlash(path))
			Expect("dir " + d).To(Equal("dir " + FromSlash(dir)))
			Expect("base " + b).To(Equal("base " + FromSlash(base)))

			od, ob := orig.Split(FromSlash(path))
			Expect("orig dir " + d).To(Equal("orig dir " + FromSlash(od)))
			Expect("orig base " + b).To(Equal("orig base " + FromSlash(ob)))
		}
		testNew := func(path, dir, base string) {
			d, b := Split2(FromSlash(path))
			Expect("dir " + d).To(Equal("dir " + FromSlash(dir)))
			Expect("base " + b).To(Equal("base " + FromSlash(base)))
		}

		It("empty", func() {
			testUnchanged("", "", "")
		})
		It("dot", func() {
			testUnchanged(".", "", ".")
		})
		It("parent", func() {
			testUnchanged("..", "", "..")
		})
		It("root", func() {
			testUnchanged("/", "/", "")
		})

		It("root nested", func() {
			testUnchanged("/root", "/", "root")
		})
		It("root dot", func() {
			testNew("/root/.", "/root", ".")
		})
		It("root parent", func() {
			testNew("/root/..", "/root", "..")
		})

		It("relative", func() {
			testUnchanged("root", "", "root")
		})
		It("relative nested", func() {
			testNew("root/nested", "root", "nested")
		})
		It("relative nested dot", func() {
			testNew("root/.", "root", ".")
		})
		It("relative nested parent", func() {
			testNew("root/..", "root", "..")
		})

		It("tagged root", func() {
			testNew("//", "/", "")
		})
		It("tagged root nested", func() {
			testNew("/root/", "/root", "")
		})
		It("tagged relative", func() {
			testNew("root/", "root", "")
		})
		It("tagged relative nested", func() {
			testNew("root/nested/", "root/nested", "")
		})
		It("tagged relative nested dot", func() {
			testNew("root/./", "root/.", "")
		})
		It("tagged relative nested parent", func() {
			testNew("root/../", "root/..", "")
		})

		It("special dir", func() {
			testNew("root/.././nested", "root/../.", "nested")
		})
		It("tagged special dir", func() {
			testNew("root/.././nested/", "root/.././nested", "")
		})

		It("double slash", func() {
			testNew("root//nested", "root", "nested")
		})
		It("double slash end", func() {
			testNew("root/nested//", "root/nested", "")
		})
	})

	Context("function Join", func() {
		It("joins two stringw", func() {
			Expect(Join("a", "b")).To(Equal("a/b"))
		})
		It("joins slashed", func() {
			Expect(Join("/a", "/b")).To(Equal("/a/b"))
		})
		It("joins slash", func() {
			Expect(Join("/", "/a")).To(Equal("/a"))
		})
		It("joins slashes", func() {
			Expect(Join("/", "/", "//")).To(Equal("/"))
		})
		It("joins empty", func() {
			Expect(Join("/", "", "")).To(Equal("/"))
		})
		It("ignore empty", func() {
			Expect(Join("", "a", "")).To(Equal("a"))
		})
		It("ignore empty", func() {
			Expect(Join("", "a", "", "b")).To(Equal("a/b"))
		})
	})
})

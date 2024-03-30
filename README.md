# go-fssup

Package **fssup** provides the `Sup` function (which is the inverse to `fs.Sub` and is missing from the Go built-in `fs` package), for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-fssup

[![GoDoc](https://godoc.org/github.com/reiver/go-fssup?status.svg)](https://godoc.org/github.com/reiver/go-fssup)

## Usage

```golang
newfsys, err := fssup.Sup(fsys, dir)
```

## Explanation

The Go built-in `fs` package has the `fs.Sub` function that takes a `fs.FS` and returns a new `fs.FS` that removes a path prefix.

For example, if the original `fs.FS` only has the following files:

* `www/index.html`
* `www/robots.txt`
* `www/img/logo.png`
* `www/img/photo.jpeg`
* `www/apple/banana/cherry.ogv`

Then calling:
```golang
newfsys, err := fs.Sub(fsys, "www")
```

... would give effectively give you a new `fs.FS` has the same files, but where the `"www/"` path prefix has been removed:

* `index.html`
* `robots.txt`
* `img/logo.png`
* `img/photo.jpeg`
* `apple/banana/cherry.ogv`

But what if you want to do the opposite?
What if you want to add a path prefix?

That is what this package provides.

If, for example, your `fs.FS` havs the following files:

* `avogadro-number.html`
* `element.htmll`
* `element/beryllium.html`
* `element/boron.html`
* `element/helium.html`
* `element/hydrogen.html`
* `element/lithium.html`
* `molecule.html`

But want them to have the path prefix `"science/chemistry"`; i.e.,:

* `science/chemistry/avogadro-number.html`
* `science/chemistry/element.htmll`
* `science/chemistry/element/beryllium.html`
* `science/chemistry/element/boron.html`
* `science/chemistry/element/helium.html`
* `science/chemistry/element/hydrogen.html`
* `science/chemistry/element/lithium.html`
* `science/chemistry/molecule.html`

Then you could accomplish that with:
```golang
newfsys, err := fssup.Sup(fsys, "science/chemistry")
```

## Import

To import package **fssup** use `import` code like the follownig:
```
import "github.com/reiver/go-fssup"
```

## Installation

To install package **fssup** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-fssup
```

## Author

Package **fssup** was written by [Charles Iliya Krempeaux](http://reiver.link)

# css-pack

CSS Packer (CSS Uglify) in go (golang)

Example

``` bash

	$ css-pack -i my_css_file.css -o my_css_file.min.css

```

Will pack and remove comments from the CSS file. 

## Options

Option                     | Description
-------------------------- | -------------------------------------
-i &lt;file.css&gt;        | Input file.
-o &lt;file.min.css&gt;    | Packed output file.
-d &lt;file.out&gt;        | Dependencies in the CSS file.  These are `url()` and `@import`.
-D                         | Debug flag

## To build

``` bash

	$ git clone https://github.com/pschlump/css-pack.git
	$ cd css-pack
	$ go get
	$ go build
	$ ./css-pack -i test2.css -o test2.min.css
	$ diff test2.ref.css test2.min.css

``` 

And a quick test

``` bash

	$ ./css-pack -i test2.css -o test2.min.css
	$ diff test2.ref.css test2.min.css

```

## TODO

1. It has the start for a `-s` source map - that code is still in progress.
2. Make it faster.


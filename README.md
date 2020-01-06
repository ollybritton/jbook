# jbook
> Go program to generate a [`mdbook`](https://github.com/rust-lang/mdBook)-formatted webpage from a [`jrnl`](https://jrnl.sh) for easy reading.

## Installation
To download the tool, you need a working Go installation. Assuming you have, you can install this program using the command

```bash
$ go get -u github.com/ollybritton/jbook
```

It relies on [`mdbook`](https://github.com/rust-lang/mdBook) to format the journal, which can be installed with

```bash
$ brew install mdbook
```

[Alternative installation instructions for `mdbook` can be found here.](https://github.com/rust-lang/mdBook#installation)

Finally, if you haven't got `jrnl` installed (which wouldn't make sense if you were going to use this tool) then you can download it using

```bash
$ brew install jrnl # For mac
$ python3 -m pip install jrnl # For other platforms
```

## Usage
To render any journal made with `jrnl`, just type `jbook` followed by the its name. Your broswer should pop open on the website generated.

```bash
$ jbook <name of jrnl>
```

Leaving it blank will use the default journal. If your journal is encrypted, you will need to enter your password.

## Bugs
* **🐞** For some reason, specifying the output for the book is not yet working. This means you can't do `jbook -o my_dir` and instead it will just create a temp dir like normal.
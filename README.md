# jbook
> Go program to generate a [`mdbook`](https://github.com/rust-lang/mdBook)-formatted webpage from a `jrnl` for easy reading.

## Installation
To download the tool, you need a working Go installation. Assuming you have, you can install this program using the command

```
$ go get -u github.com/ollybritton/jbook
```

It relies on [`mdbook`](https://github.com/rust-lang/mdBook) to format the journal, which can be installed with

```
brew install mdbook
```

[Alternative installation instructions for `mdbook` can be found here.](https://github.com/rust-lang/mdBook#installation)

## Usage
To render any journal made with `jrnl`, just type `jbook` followed by the its name.

```
$ jbook <name of jrnl>
```

Leaving it blank will use the default journal. If your journal is encrypted, you will need to enter your password.
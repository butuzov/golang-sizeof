# Golang sizeof tips

_Moved from [gophergala/golang-sizeof.tips](https://github.com/gophergala/golang-sizeof.tips)_

_Moved from [tyranron/golang-sizeof.tips](https://github.com/tyranron/golang-sizeof.tips)_

<h2> Golang sizeof tips <del>Closed</del><ins>Open</ins>. Long Live Golang sizeof tips!</h2>

**Consider alternatives:**
- https://github.com/dominikh/go-tools/tree/master/cmd/structlayout

Web tool for interactive playing with Golang struct sizes.

## Changes

At the 21 May 2020 I (Oleg Butuzov) changed this software in next way:

- process `demonization` removed.
- made it locally installable via `go get -u github.com/butuzov/golang-sizeof`
- address are hardcoded (localhost:7777) and browser (supppose to be opened) when server starts.
- windows support?

That all of the changes I made to original tool.

## Aim
Provide comfortable tool to see how fields in struct are aligned,
to compare different structs and as the result - to understand
and remember alignment rules.

## Installing

```bas
go get -u github.com/butuzov/golang-sizeof
golang-sizeof
```

## Platform support
Tested on Linux and OS X x64 platforms, but should work properly and on other
*nix-like platforms. Should work on windows too, but i didn't test it, so can't confirm it.

## License
[Apache License 2.0](LICENSE)

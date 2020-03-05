# Golang sizeof tips


_Moved from [gophergala/golang-sizeof.tips](https://github.com/gophergala/golang-sizeof.tips)_

_Moved from [tyranron/golang-sizeof.tips](https://github.com/tyranron/golang-sizeof.tips)_

## Golang sizeof tips Closed. Long Live Golang sizeof tips!


**Consider alternatives:**
- https://github.com/dominikh/go-tools/tree/master/cmd/structlayout

Web tool for interactive playing with Golang struct sizes.


## Aim
Provide comfortable tool to see how fields in struct are aligned,
to compare different structs and as the result - to understand
and remember alignment rules.

## Installing

```bash./gs-
git clone https://github.com/butuzov/golang-sizeof
make build
```

## Usage

```bash
make start
make stop
make restart
```

## Platform support
Tested on Linux and OS X x64 platforms, but should work properly and on other
*nix-like platforms.
Windows is not supported due to daemonization.

## License
[Apache License 2.0](LICENSE)

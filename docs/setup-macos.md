# Setup for macOS

One time setup on macOS

> Scripts are `bash`

## Xcode / Compilers

Install Command Line Tools (CLT) for Xcode:

```bash
xcode-select --install
```

## Homebrew

Install Homebrew:

```bash
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

## bash 5

```bash
brew install bash
```

## goenv

Install Go version manager [goenv](https://github.com/go-nv/goenv)

```bash
brew install goenv
```

Install Go:

```bash
# install go 1.21.5
goenv install 1.21.5

# to install latest Go version
goenv install latest
```

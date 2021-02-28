# Gocket

A simple CLI and TUI for Pocket.

![Logo of Gocket](./logo.jpg)

## Installation

### General

You can simply grab the [latest released binary file](https://github.com/Phantas0s/gocket/releases/latest) and download the version you need, depending on your OS.

### Linux script

If you use a Linux-based OS, here's a simple way to download gocket and move it in `/usr/local/bin` to be able to use it globally.

```shell
curl -LO https://raw.githubusercontent.com/Phantas0s/gocket/master/install/linux.sh && \
./linux.sh && \
rm linux.sh
```

### Manual installation

You need to clone this repository and build the binary in the root directory.

## Authorization



## Usage
## Keybindings
## Video Tutorial


### Newest Article Using Your Browser

```
gocket list -k <consumerKey> -c 1 | lynx -
gocket list -k <consumerKey> -c 1 | xargs firefox
```

## References

### Pocket List

https://getpocket.com/my-list

### API

https://etpocket.com/developer/docs/authentication
https://getpocket.com/developer/docs/v3/retrieve

### Inspiration

[go-pocket](https://github.com/motemen/go-pocket) has been the base I've refactored and built upon. Thanks for this great project!
https://getpocket.com/developer/apps/

### Libraries

https://pkg.go.dev/github.com/rivo/tview

## Model

Keybindings: https://github.com/jesseduffield/lazydocker/blob/master/docs/keybindings/Keybindings_en.md


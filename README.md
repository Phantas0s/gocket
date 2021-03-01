# WIP - this is not ready yet :)

# Gocket

A simple CLI (or TUI) for Pocket.

![Logo of Gocket](./logo.jpg)

## Installation

### General

You can simply grab the [latest binary file](https://github.com/Phantas0s/gocket/releases/latest) and download the version you need, depending on your OS.

### Linux script

If you use a Linux-based OS, here's a simple way to download gocket and move it to `/usr/local/bin`. You can then call it wherever you want.

```shell
curl -LO https://raw.githubusercontent.com/Phantas0s/gocket/master/install/linux.sh && \
./linux.sh && \
rm linux.sh
```
### Manual installation

You need to clone this repository and build the binary in the root directory.

## Authorization

You need to authorize gocket to access your Pocket account. It's very easy:

1. Go to [Gocket apps and create an app](https://getpocket.com/developer/apps/).
2. The easiest way is to create a config file in your current directory or in `$XDG_CONFIG_HOME/gocket/config.yml`
    * Create an entry with `key` as index and the consumer key as value, i.e `key: 1234-5a6b7c`.
    * Your config can be YAML, TOML, or JSON file.
3. You'll have to confirm the authorization: a webpage will open in your new favorite browser to do so.
4. You can use gocket!

XDG_CONFIG_HOME: (TODO)
Unix systems: `~/.config` 
macOS: ~/Library/Application Support
Windows: %LOCALAPPDATA%

If you wonder what's the value of $XDG_CONFIG_HOME for your system, you can look [at this page](https://github.com/adrg/xdg).

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

# WIP - this is not ready yet :)

# Gocket

A simple CLI (or TUI) for Pocket.

![Logo of Gocket](./logo_smaller.png)

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

### Steps

1. Go to [Gocket apps and create an application](https://getpocket.com/developer/apps/)
2. Authorize the application to add, modify, and retrieve if you want to use the full set of gocket's feature
3. You need to pass the consumer key to pocket each time you use it (`-k` option) or you can use a config file:
    1. Create the file `$XDG_CONFIG_HOME/gocket/config.yml`
    2. Create an entry with `key` as index and the consumer key as value, for example `key: 1234-5a6b7c`
    3. Your config can be a YAML, TOML, or JSON file
3. The first time you use pocket, you'll need to confirm your authorization. A webpage will open automatically in your favorite browser to do so
4. Enjoy!

### XDG Home Directory

The value of `$XDG_CONFIG_HOME` depends of your OS. Here are the defaults (if you didn't modify it):

* **Unix systems**: `~/.config`
* **macOS**: `~/Library/Application Support`
* **Windows**: `%LOCALAPPDATA%`

## Commands

You have access to different commands. Use the option `-h` for each command to access the help.

### List

* `gocket list`: list your Pocket entries.
* `gocket list archive`: list the archives.

The options for these two commands are almost identical. Here are the difference:
* Use `-a` with `gocket list` to bulk add every listed entry to the archive (with confirmation).
* Use `-a` with `gocket list archive` to bulk add every listed archive to the unread list (with confirmation).

### Add New URLs

* `gocket add <URL>`: Add the URL `<URL>` to pocket. You can add multiple URLs separated with spaces.

## Usage

| Description                                                         | Command                              |
| ----                                                                | ----                                 |
| Output every page's URLs                                            | `gocket list`                        |
| Output the last 5 pages' URLs added                                 | `gocket list -c 5`                   |
| Display the last 5 pages added in a TUI                             | `gocket list -c 5 --tui`             |
| Display pages in a TUI and don't ask confirmation for any operation | `gocket list -c 5 --tui --noconfirm` |
| Search for "youtube" in titles and URLs                             | `gocket list -s "youtube" -t`        |
| Open the last page added with Firefox                               | `gocket list -c 1 \| xargs firefox` |
| Open the last page added with Lynx                                  | `gocket list -c 1 \| lynx -`        |
| Open the oldest page added with Firefox                             | `gocket list -c 1 -o "oldest" \| xargs firefox` |
| Open the last 5 pages with Firefox and archive it                   | `gocket list -c 5 -a --noconfirm \| xargs firefox` |
| Open the last page with Firefox and delete it                       | `gocket list -c 1 -d --noconfirm \| xargs firefox` |

## TUI Keybindings

### Navigation

<pre>
 <kbd>↑</kbd> or <kbd>k</kbd>: up
 <kbd>↓</kbd> or <kbd>j</kbd>: down
 <kbd>PgUp</kbd> or <kbd>CTRL</kbd>+<kbd>u</kbd>: One screen up
 <kbd>PgDn</kbd> or <kbd>CTRL</kbd>+<kbd>d</kbd>: One screen down
 <kbd>Home</kbd> or <kbd>g</kbd>: Top of the list
 <kbd>End</kbd> or <kbd>G</kbd>: Bottom of the list
</pre>

### Action

<pre>
 <kbd>d</kbd>: Delete Pocket entry
 <kbd>a</kbd>: Add (if list archive) or archive (if list unread)
</pre>

## Acknowledgements

* Thanks to the project [go-pocket](https://github.com/motemen/go-pocket), I had a quick base I modified and built upon.
* Thanks to [MariaLetta](https://github.com/MariaLetta/free-gophers-pack) for the awesome and beautiful Gopher pack! I used it for my logo on top.
* Thanks to [Lukasz Adam](https://lukaszadam.com/illustrations) for his free and amazing illustrations I use basically everywhere.

## Licence

[Apache Licence 2.0](https://choosealicense.com/licenses/apache-2.0/)

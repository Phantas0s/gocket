## [0.2.1] - 2021-03-09

### Added

* `gocket add` - verbose `-v` option to display a message when URLs are successfully added

## [0.2.0] - 2021-03-09

### Added

* `gocket list` and `gocket list archive` - New option `-f` to filter by type 'article' (default), 'video' or image
* `gocket add` - Read from stdin when no arguments given
* Possible to configure gocket using environment variables (need prefix `GOCKET_`, i.e `GOCKET_TUI=true`)

## [0.1.0] - 2021-03-07

### Added

* Current interface:

```
Pocket in the shell

Usage:
  gocket [command]

Available Commands:
  add         Add a pocket article
  help        Help about any command
  list        List your pocket articles

Flags:
  -h, --help         help for gocket
  -k, --key string   Pocket consumer key (required).

Use "gocket [command] --help" for more information about a command.
```

* List command:

```
List your Pocket pages

Usage:
  gocket list [flags]
  gocket list [command]

Available Commands:
  archive     List your Pocket archive

Flags:
  -a, --archive         Archive the listed articles (with confirmation)
  -c, --count int       Number of results (0 for all)
  -d, --delete          Delete the listed articles (with confirmation)
  -h, --help            help for list
      --noconfirm       Don't ask for any confirmation
  -o, --order string    order by 'newest', 'oldest', 'title', or 'url' (default "newest")
  -s, --search string   Search by title and URL
  -t, --title           Display the titles
      --tui             Display the results in a TUI

Global Flags:
  -k, --key string   Pocket consumer key (required)

Use "gocket list [command] --help" for more information about a command.
```

* List archive command:

```
List your Pocket archive

Usage:
  gocket list archive [flags]

Flags:
  -a, --add    Add the listed articles back to unread (with confirmation).
  -h, --help   help for archive

Global Flags:
  -c, --count int       Number of results (0 for all)
  -d, --delete          Delete the listed articles (with confirmation)
  -k, --key string      Pocket consumer key (required)
      --noconfirm       Don't ask for any confirmation
  -o, --order string    order by 'newest', 'oldest', 'title', or 'url' (default "newest")
  -s, --search string   Search by title and URL
  -t, --title           Display the titles
      --tui             Display the results in a TUI

```

* Add command:

```
Add a page to Pocket

Usage:
  gocket add URL... [flags]

Flags:
  -h, --help   help for add

Global Flags:
  -k, --key string   Pocket consumer key (required)
```

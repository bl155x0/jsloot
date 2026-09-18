# jsloot 🐿

`jsloot` is a handy tool designed to store and beautify JavaScript files captured by your favorite offensive web proxy.

## Purpose

The main goal of `jsloot` is to collect and beautify JavaScript (and, if the feeding proxy chooses to also
pass along HTML responses, whole HTML pages) while manually investigating a target with a web proxy.
`jsloot` itself does no fetching or parsing — it just stores whatever content it's handed and, if that
content looks like JavaScript, beautifies it.

## Setup

### Dependencies

To beautify JavaScript files, you'll need [jsbeautifier](https://pypi.org/project/jsbeautifier/), which provides the `js-beautify` CLI that `jsloot` shells out to. Install it with:

```bash
pip install jsbeautifier
```

On distros with an externally-managed Python (e.g. Arch), plain `pip install` will refuse to run. Use `pipx` instead:

```bash
sudo pacman -S python-pipx   # or your distro's equivalent
pipx install jsbeautifier
```

`jsloot` finds `js-beautify` via `PATH`, so it must be on the `PATH` of whatever process actually runs `jsloot` — not just your interactive shell. This matters if `jsloot` is invoked from a GUI app (e.g. Caido, see below): `pipx` installs to `~/.local/bin`, which GUI apps launched from a desktop environment typically don't have on `PATH`. If `js-beautify` isn't found from that context, symlink it somewhere that is on the GUI app's `PATH` (commonly `/usr/local/bin`):

```bash
sudo ln -sf "$(which js-beautify)" /usr/local/bin/js-beautify
```

### Installation

```bash
go install github.com/bl155x0/jsloot@latest
```

### Caido

To seamlessly integrate `jsloot` with Caido, install the following passive Workflow to your Caido project:

https://github.com/bl155x0/caido/tree/main/workflows/passive/JSLoot


# Usage

`jsloot` has a single sub-command:

### `store`

The `store` command reads a single file as JSON from stdin and writes it to disk under a host-based
layout (`<directory>/<host>/<filename>`, filename taken from the URL's path). It's meant for content
already captured elsewhere (e.g. by a proxy that intercepted the response), so the file ends up on disk
without jsloot making any network request of its own.

The JSON must contain:
- `URL`: the URL the content was captured from (used to derive the local path)
- `content`: the content to store
- `contentType` (optional): the response's Content-Type. If given and it doesn't look like
  JavaScript (`text/javascript`, `application/javascript`, ...), beautification is skipped for that
  file even if `-b`/beautify is on — running the JS beautifier on non-JS content (e.g. HTML) corrupts
  it. If omitted, the file is always beautified when `-b` is on (legacy behaviour).

```bash
echo '{"URL": "https://www.example.com/example.js", "content": "var x = 1;", "contentType": "text/javascript"}' | jsloot store -d /tmp/jsloot
```

Beautification is on by default; pass `-b=false` to skip it unconditionally. An existing file at the
resolved path is overwritten.

<br>
<hr>
<br>
★ ♥ 🐿 ~ HAPPY LOOTING ~

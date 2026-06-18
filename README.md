# jsloot 🐿

`jsloot` is a handy tool designed to download and beautify JavaScript files, often used alongside your favorite offensive web proxy.

## Purpose

The main goal of `jsloot` is to collect, download, and beautify JavaScript files while manually investigating a target with a web proxy. 
As you browse your target manually, jsloot automatically downloads or collects JavaScript URLs to a local file, making them ready for further investigation.

## Setup

### Dependencies

To beautify JavaScript files, you'll need [jsbeautifier](https://pypi.org/project/jsbeautifier/). Install it with:

```bash
pip install jsbeautifier
```

### Installation

```bash
go install github.com/bl155x0/jsloot@latest
```

### Caido

To seamlessly integrate `jsloot` with Caido, install one, or all of the following passive Workflows to your Caido project:

https://github.com/bl155x0/caido/tree/main/workflows/passive/JSLoot


# Usage
`jsloot` offers various distinct sub-commands:


### `add`
The `add` command is used to collect JavaScript URLs while investigating a target with a proxy tool like Caido. 
It appends JavaScript URLs to a text file (jsloot.txt).

```bash
jsloot add -f jsloot.txt "http://example.com/example.js"
```
If the [JSLootAdd](https://github.com/bl155x0/caido/blob/main/workflows/passive/JSLoot/JSLootAdd.json) Caido passive Workflow is installed (see [Caido](#Caido)), this command is executed automatically,
for every recognized JavaScript file while browsing your target with Caido.

### `getall`

The `getall` command downloads and beautifies all collected JavaScript URLs from a given file into to a local folder.
``` bash
jsloot getall -f jsloot.txt
```

### `store`

The `store` command reads a single JS file as JSON from stdin and writes it to disk, using the same host-based
layout `getall` uses when downloading (`<directory>/<host>/<filename>`). It's meant for cases where the content
was already captured elsewhere (e.g. by a proxy that intercepted the response), so the file ends up on disk
exactly as if `jsloot` had downloaded it itself, without an extra network request.

The JSON must contain:
- `URL`: the URL the content was downloaded from (used to derive the local path)
- `content`: the JavaScript source to store

```bash
echo '{"URL": "https://www.example.com/example.js", "content": "var x = 1;"}' | jsloot store -d /tmp/jsloot
```

An existing file at the resolved path is overwritten.

<br>
<hr>
<br>
★ ♥ 🐿 ~ HAPPY LOOTING ~

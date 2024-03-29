# jsloot 🐿

`jsloot` is a handy tool designed to download and beautify JavaScript files, often used alongside your favorite offensive web proxy.

## Purpose

The main goal of `jsloot` is to collect, download, and beautify JavaScript files while manually investigating a target with a web proxy. 
As you browse your target manually, jsloot automatically downloads or collects JavaScript URLs to a local file, making them ready for further investigation.

## Setup

```bash
go install github.com/bl155x0/jsloot@latest
```

### Dependencies

To beautify JavaScript files, you'll need [jsbeautifier](https://pypi.org/project/jsbeautifier/). Install it with:

```bash
pip install jsbeautifier
```

### Caido

To seamlessly integrate `jsloot` with Caido, install the following passive Workflow to your Caido project:

https://github.com/bl155x0/caido/tree/main/workflows/passive/JSLoot


# Usage
`jsloot` various distinct sub-commands:


### `add`
The `add` command is used to collect JavaScript URLs while investigating a target with a proxy tool like Caido. 
It appends JavaScript URLs to a text file (jsloot.txt).

```bash
jsloot add -f jsloot.txt "http://example.com/example.js"
```
If the [JSLootAdd](https://github.com/bl155x0/caido/blob/main/workflows/passive/JSLoot/JSLootAdd.json) Caido passive Workflow is installed (see [Caido](#Caido)), this command is executed automatically,
for every recognized JavaScript file while browsing your target with Caido.

### `get`

The `get` command downloads and beautifies a specific JavaScript file and beautifies it.
``` bash
jsloot get -u "https://www.example.com/example.js"
```

If the Caido [JSLootGet](https://github.com/bl155x0/caido/blob/main/workflows/passive/JSLoot/JSLootGet.json) Caido passive Workflow is installed (see [Caido](#Caido)), this command is executed automatically,
for every recognized JavaScript file while browsing your target with Caido.

### `getall`

The `getll` command downloads and beautifies all collected JavaScript URLs from a given file into to a local folder.
``` bash
jsloot loot -f jsloot.txt
```
<br>
<hr>
<br>
★ ♥ 🐿 ~ HAPPY LOOTING ~

Altogether
===

[![Build Status](https://dev.azure.com/announce/altogether/_apis/build/status/announce.altogether?branchName=master)](https://dev.azure.com/announce/altogether/_build/latest?definitionId=4&branchName=master)

## Altogether in a Nutshell

Altogether is a CLI tool to sync configuration files between 2 major keyboard launchers -- [Alfred](https://www.alfredapp.com/) and [Albert](https://albertlauncher.github.io/).
So the target user is who adopts both Mac and Linux on a daily basis, or someone needs to migrate the one's config to the another.

## Supported Features

Supported configuration files are the ones relating to:

* [x] Web search
* [ ] Snippets

## Installation

* The most primitive way is to download binary package

## Usage

1. Prepare required parameters
    * `AL2_ALFRED_PATH`: Specify a path to Alfred's config directory (default path is `${HOME}/.config/Alfred.alfredpreferences`)
    * `AL2_ALBERT_PATH`: Specify a path to Albert's config directory (default path is `${HOME}/.config/albert`)
1. Execute commands like as following:

```bash
export AL2_ALFRED_PATH="__ALFRED_CONFIG_DIR__" AL2_ALBERT_PATH="__ALBERT_CONFIG_DIR__"
./altogether help
```

```text
```

## Supported Versions

[Alfred](https://www.alfredapp.com/changelog/):

* 3.8.x

[Albert](https://albertlauncher.github.io/docs/changelog/):

* 0.16.x


## Contribution

Here's how to get started!

1. Install [Docker](https://docs.docker.com/install/) (verified version: `18.09.1-ce`)
1. Build container:
 
 ```bash
 ./script/ci.sh init
```

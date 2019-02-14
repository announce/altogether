Altogether
===

[![Build Status](https://dev.azure.com/announce/altogether/_apis/build/status/announce.altogether?branchName=master)](https://dev.azure.com/announce/altogether/_build/latest?definitionId=4&branchName=master)

## Altogether in a Nutshell

Altogether is a CLI tool to synchronize configuration files between 2 major keyboard launchers -- [Alfred](https://www.alfredapp.com/) and [Albert](https://albertlauncher.github.io/).
So the target user is who adopts both Mac and Linux on a daily basis, or someone needs to migrate one's config to the another.

## Supported Features

Supported configuration files are the ones relating to:

* [x] Web search
* [ ] Snippets

## Installation

* TBD: The most primitive way is to download binary package

## Usage

#### Command Arguments

Specify required parameters in environmental variables:

* `AL2_ALFRED_PATH`: a path to Alfred's config directory
* `AL2_ALBERT_PATH`: a path to Albert's config directory
* `AL2_DRY_RUN`: set `1` to dump merged configurations 
* `AL2_VERBOSE`: set `1` to print out detailed logs

You can execute commands like as following:

```bash
export AL2_ALFRED_PATH="${HOME}/.config/Alfred.alfredpreferences"
export AL2_ALBERT_PATH="${HOME}/.config/albert"
export AL2_DRY_RUN=1
export AL2_VERBOSE=1
./altogether sync-web
```

#### Systemd Configurations

1. Place unit files to `~/.config/systemd/user/`. Sample systemd configuration files are available at `./sample`.
1. Run commands like below to test:

```bash
systemctl --user daemon-reload && systemctl --user restart altogether
journalctl --user -xe -u altogether
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
./script/ci.sh ci
```

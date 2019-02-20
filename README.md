Altogether
===

[![Build Status](https://dev.azure.com/announce/altogether/_apis/build/status/announce.altogether?branchName=master)](https://dev.azure.com/announce/altogether/_build/latest?definitionId=4&branchName=master)

## Altogether in a Nutshell

Altogether is a CLI tool to synchronize configuration files between 2 major keyboard launchers -- [Alfred](https://www.alfredapp.com/) and [Albert](https://albertlauncher.github.io/).
So the target user is who adopts both Mac and Linux on a daily basis, or someone needs to migrate one's config to the another.

## Available Features

Supported configuration files are the ones relating to:

* [x] Web search
* [ ] Snippets

## Installation

#### For Mac users

Go ahead and grab Homebrewed `altogether`:

```bash
brew install announce/homebrew-altogether/altogether
```

#### For Arch Linux users

If you are using Arch Linux, [AUR Package](https://aur.archlinux.org/packages/altogether/) is recommended.
Install `altogether` with AUR helper such as `yay`:

```bash
yay -Syu altogether
```

#### How to get `altogether` command

Unless specified above, one of the simplest ways for your platform is:

1. Download package from [![the latest release](https://img.shields.io/github/release/announce/altogether.svg?style=flat)](https://github.com/announce/altogether/releases/latest) according to platform:
    * macOS: `darwin-amd64.tar.gz`
    * Linux: `linux-amd64.tar.gz`
1. Extract the compressed tarball and place `altogether` to somewhere under `PATH` directory:

```bash
curl -L https://github.com/announce/altogether/releases/download/v0.1.2/linux-amd64.tar.gz | tar -zxvf - \
&& mv ./linux-amd64/altogether /usr/local/bin \
&& altogether --help
```

## Usage

#### Command Arguments

Specify required parameters in environmental variables:

* `AL2_ALFRED_PATH`: a path to Alfred's config directory
* `AL2_ALBERT_PATH`: a path to Albert's config directory
* `AL2_DRY_RUN`: set `1` to dump merged configurations without actual file change
* `AL2_VERBOSE`: set `1` to print out detailed logs

You can execute commands like as following:

```bash
export AL2_ALFRED_PATH="${HOME}/.config/Alfred.alfredpreferences"
export AL2_ALBERT_PATH="${HOME}/.config/albert"
export AL2_DRY_RUN=1
export AL2_VERBOSE=1
./altogether sync-web
```

Or, you also can specify command options like below:

```bash
altogether sync-web --dry-run --alfred-path ~/.config/Alfred.alfredpreferences --albert-path ~/.config/testdata/albert
```

#### TBD: Systemd Configurations

Consider set up Systemd configuration if you want to automatically sync configurations.

1. Place unit files to `~/.config/systemd/user/`. Sample systemd configuration files are available under `./sample`.
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
1. Build a container and start debugging:
 
 ```bash
./script/ci.sh init
./script/ci.sh ci
```

Related repositories:

* [announce/homebrew-altogether](https://github.com/announce/homebrew-altogether)
* [announce/aur-altogether](https://github.com/announce/aur-altogether)

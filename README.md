[![Github tag](https://badgen.net/github/tag/denkylabs/discord-api-types-go)](https://github.com/denkylabs/discord-api-types-go/tags/) [![CodeFactor](https://www.codefactor.io/repository/github/denkylabs/discord-api-types-go/badge)](https://www.codefactor.io/repository/github/denkylabs/discord-api-types-go) [![DeepSource](https://deepsource.io/gh/denkylabs/discord-api-types-go.svg/?label=active+issues&show_trend=true&token=-7FwA0KciBj7gHuWLxwqoHPl)](https://deepsource.io/gh/denkylabs/discord-api-types-go/?ref=repository-badge) [![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/) [![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/denkylabs/discord-api-types-go.svg)](https://github.com/denkylabs/discord-api-types-go)

# discord-api-types
Up to date Discord API Typings for Golang

## Installation
### Stable version:
```sh-session
go get -u github.com/denkylabs/discord-api-types-go
```

### Development version:
```sh-session
$ cd $GOPATH
$ mkdir -p src/github.com/denkylabs
$ cd src/github.com/denkylabs
$ git clone https://github.com/denkylabs/discord-api-types-go.git
$ cd discord-api-types-go
$ go install
```

## Project Structure

The exports of each API version is split into three main parts:

- Everything exported with the `API` prefix represents a payload you may get from the REST API _or_ the Gateway.

- Everything exported with the `Gateway` prefix represents data that ONLY comes from or is directly related to the Gateway.

- Everything exported with the `REST` prefix represents data that ONLY comes from or is directly related to the REST API.

- Anything else that is miscellaneous will be exported based on what it represents (for example the `REST` route object).

**Warning**: This package documents just KNOWN (and documented) properties. Anything that isn't documented will NOT be added to this package (unless said properties are in an open Pull Request to Discord's [API Documentation repository](https://github.com/discord/discord-api-docs) or known through other means _and have received the green light to be used_). For clarification's sake, this means that properties that are only known through the process of data mining and have not yet been confirmed in a way as described will **NOT** be included.

## Useful links

- [discordjs/discord-api-types](https://github.com/discordjs/discord-api-types/)
- [discord/discord-api-docs](https://github.com/discord/discord-api-docs/)
- [Discord Developer Portal](https://discord.com/developers/docs/intro)
- [Discord Developers server](https://discord.gg/discord-developers)

## Compatibility
| Go version          |  Status             |
| ------------------- | ------------------- |
|  Go 1.18  |  <a href="https://github.com/denkylabs/discord-api-types-go/actions"><img src="https://github.com/denkylabs/discord-api-types-go/actions/workflows/go-1.18.yml/badge.svg" alt="Tests status" /></a> |
| Go 1.17  | Not supported. |
| Go 1.16  | Not supported. |

[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://lbesson.mit-license.org/)

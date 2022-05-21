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

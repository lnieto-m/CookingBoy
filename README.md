# CookingBoy

Basic Discord Music Bot

### Prerequisites
* [ffmpeg](https://ffmpeg.org/download.html)
* [Go](https://golang.org/doc/install)

## Install

Clone this repo and make a symbolic link to the `src` directory<br>
Install needed packages
```
go get github.com/bwmarrin/discordgo
go get layeh.com/gopus
go get golang.org/x/net/context
go get golang.org/x/oauth2
go get golang.org/x/oauth2/google
go get google.golang.org/api/youtube/v3
```
Go to your cloned repo and build it
```
go build && ./CookingBoy
```

## Basic Usage

```s!<command> <args>```<br>
Type `s!help` for basic commands

## Built With
* [DiscordGo](https://github.com/bwmarrin/discordgo)
* [Youtube Binds for GO](https://godoc.org/google.golang.org/api/youtube/v3)

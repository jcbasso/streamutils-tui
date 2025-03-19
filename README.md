# Stream Utils TUI
![Go Version](docs/assets/go-version-badge.svg)

This Terminal UI application aims to provide utils from streaming apps such as Twitch chat.

## Env Vars
| Environment Variable | Required | Description                                                                                                              |
|----------------------|----------|--------------------------------------------------------------------------------------------------------------------------|
| TWITCH_OAUTH_TOKEN   | True     | Twitch User Access Token. You can obtain it from [https://twitchtokengenerator.com/](https://twitchtokengenerator.com/). |
| TWITCH_USERNAME      | True     | Username associated to previous Access Token.                                                                            |
| TWITCH_CHANNEL       | True     | Channel name to join.                                                                                                    |

## Build Requirements
- Go 1.24+ (Older versions might work, but using 1.22 for development)

## Execute

```go
go run .
```
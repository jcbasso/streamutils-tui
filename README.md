<div align="center">

# Stream Utils TUI

[![go version](docs/assets/badge/go-version-badge.svg)](https://github.com/jcbasso/streamutils-tui/blob/main/go.mod)
[![release](docs/assets/badge/release-badge.svg)](https://github.com/jcbasso/streamutils-tui/releases)

Terminal UI application aiming to provide utils for streaming apps such as Twitch chat.

</div>

---

## Env Vars
| Environment Variable | Required | Description                                                                                                              |
|----------------------|----------|--------------------------------------------------------------------------------------------------------------------------|
| TWITCH_OAUTH_TOKEN   | True     | Twitch User Access Token. You can obtain it from [https://twitchtokengenerator.com/](https://twitchtokengenerator.com/). |
| TWITCH_USERNAME      | True     | Username associated to previous Access Token.                                                                            |
| TWITCH_CHANNEL       | True     | Channel name to join.                                                                                                    |

## Execute

```go
go run .
```
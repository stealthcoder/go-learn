# Learn Go

__Learn Go__ is a personal project of my learning journey with Golang.


## Create Modules

Replaces `example.com/greetings` for location dependency

- `go mod edit -replace example.com/greetings=../greetings`

Sync module dependencies

- `go mod tidy`
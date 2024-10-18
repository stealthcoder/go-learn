# Learn Go

__Learn Go__ is a personal project of my learning journey with Golang.


## Create Modules

Enabling dependency tracing for your code

- `go mod init <path_of_module>`
- `go mod init example.com/greetings`

For unpublished an module, we adapt `example.com/greetings` module to redirect Go tools to find it from our local file system
- `go mod edit -replace <published_module_path>=<to_local_module_path>`
- `go mod edit -replace example.com/greetings=../greetings`

Sync module dependencies (e.g. adding those required by the code, but not yet tracked in the module)

- `go mod tidy`
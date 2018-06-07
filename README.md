# gologger
Based on [Codis](https://github.com/CodisLabs/codis/blob/release3.2/pkg/utils/log/log.go) log package.

# features
Ability to combine [fluent](https://github.com/fluent/fluentd) to colllect logs.
If connect to fluentd failed, log to fluentd will be disabled.

# usage

This package offers functions with the following patterns:
- `[logLevel]`, just log plain message
- `[logLevel]f`, message will be logged with format
- `[logLevel]Error`, log with extra stack info
- `[logLevel]Errorf`, message will be logged with format and extra stack info

logLevel can be:
- Panic, will call `os.Exit(1)` after message logged, use it carefully.
- Error, will automatically log stack info and ignores log level setting.
- Warn
- Info
- Debug

Additionaly, offers three more functions which ignore log level and taged by [LOG]:
- Print
- Printf
- Println

also see `test/main.go`

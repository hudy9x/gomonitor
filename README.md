### Introduction

It will help you monitor what application are u focus on. For instance, you open chrome to browse google.com, `gomonitor` will print out the application's title. You can save logs to your database for later use.

### Usage

First, you have to install `golang` for sure. Next, open terminal and install go packages

```shell
$ go mod tidy
```

Run tool

```shell
$ go run main.go
```

Build command for Windows. The `-o` option mean output file, `-ldflags="-H windowsgui"` this option tell Windows hide Terminal as running.

```shell
$ go build -o gomonitor.exe -ldflags="-H windowsgui" ./main.go
```

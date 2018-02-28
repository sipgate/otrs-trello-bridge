# OTRS Trello bridge

## Introduction

> Keep a trello board in sync with your OTRS


## Building

To build the otrs trello bridge you need a [go](https://golang.org/doc/install) runtime
and the [dep](https://golang.github.io/dep/docs/installation.html) dependency manager

After installing the prerequisites simply run
```bash
./build.sh
```

This will produce the `otrs-trello-bridge` statically linked binary

## Running

Before running, make sure you copy the `config.toml.dist` file to `config.toml` in your working directory
Then modify the settings to match your setup.

After that you should be able to run the binary:
```bash
./otrs-trello-bridge
```

The bridge defaults to port 8080 but can be overridden via the `PORT` environment variable.

For more information please refer to the [gin docs](https://gin-gonic.github.io/gin/) 
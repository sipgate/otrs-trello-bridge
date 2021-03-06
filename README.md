# OTRS bridge

## Introduction

### Features

* Keep a trello board in sync with your OTRS
* Post new Tickets to Slack channel

[![Build Status](https://travis-ci.org/sipgate/otrs-bridge.svg?branch=master)](https://travis-ci.org/sipgate/otrs-bridge)
[![Go Report Card](https://goreportcard.com/badge/github.com/sipgate/otrs-bridge)](https://goreportcard.com/report/github.com/sipgate/otrs-bridge)

## Running

You can grab the latest release build at the [releases page](releases).

Before running, make sure you copy the `config.toml.dist` file to `config.toml` in your working directory.
Then modify the settings to match your setup. You can override the config file location via the `-config` argument.

After that you should be able to run the binary:
```bash
./otrs-bridge -config /usr/local/etc/otrs-bridge/config.toml
```

The bridge defaults to port 8080 but can be overridden via the `PORT` environment variable.

## Building

To build the otrs bridge you need a [go](https://golang.org/doc/install) runtime
and the [dep](https://golang.github.io/dep/docs/installation.html) dependency manager.

After installing the prerequisites simply run:
```bash
./scripts/build.sh
```

This will produce the `otrs-bridge` statically linked binary

## OTRS Setup

Import the webservice definitions from the [otrs-webservices](otrs-webservices) directory, make sure to use the correct URL of the bridge service

## Hacking

Just open the directory and start hacking.

## License

```text
Copyright 2018 sipgate GmbH

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
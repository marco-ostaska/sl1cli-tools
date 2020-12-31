![](docs/logo/logo.png)

sl1cmd is a command line interface to interact with [ScienceLogic](https://sciencelogic.com/product/technologies/compute) Monitoring tool API.


[![Go Report Card](https://goreportcard.com/badge/github.com/marco-ostaska/sl1cmd)](https://goreportcard.com/report/github.com/marco-ostaska/sl1cmd)

# Table of Contents

- [Overview](#overview)
  - [Commands](#commands)
  - [Packages](#packages)
- [Installing](#intalling)
- [Getting Started](#getting-started)


# Overview

sl1cmd is a command line interface provides an interface to interact over daily tasks admonistrations using [ScienceLogic](https://sciencelogic.com/product/technologies/compute) Monitoring tool API.

## Commands

- [sl1cmd](docs/cmd/sl1cmd.md) - sl1cmd is a command line interface to interact with ScienceLogic Monitoring tool API.

## Packages 
[![GoDoc](https://godoc.org/github.com/marco-ostaska/sl1cmd?status.svg)](https://godoc.org/github.com/marco-ostaska/sl1cmd)

- [sl1/vault](https://godoc.org/github.com/marco-ostaska/sl1cmd/pkg/sl1/vault) - Package vault manages encryption for sl1cmd credentials.
- [sl1/httpcalls](https://godoc.org/github.com/marco-ostaska/sl1cmd/pkg/sl1/httpcalls) - Package httpcalls makes GET, DELETE and POST calls
- [wrappper](https://godoc.org/github.com/marco-ostaska/sl1cmd/pkg/wrapper) - Package wrapper wraps some boring tasks.


# Installing

[download](../../releases) the binary file and run it. Simple as it sounds. 

# Getting Started

Before using the sl1cmd to perform admnistrative tasks you must [configure the sl1cmd vault](docs/cmd/sl1cmd_vault_new.md)


 

    


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

### Vault
- [vault](docs/cmd/sl1cmd_vault.md) - create or update login information vault for api.

### User Administration
- [id](docs/cmd/sl1cmd_id.md) - print users sl1 id for the specified user.
- [userinfo](docs/cmd/sl1cmd_userinfo.md) - print user information for the specified user.
- [useradd](docs/cmd/sl1cmd_useradd.md) -  create a new user.
- [userdel](docs/cmd/sl1cmd_userdel.md) - delete user account.
- [passwd](docs/cmd/sl1cmd_passwd.md) - change user password.

## Packages

- [sl1/vault](docs/pkg/vault.md) - Package vault manage encryption for sl1cmd credentials.
- [httpcalls](docs/pkg/httpcalls.md) - Package httpcalls makes http request calls on sl1api.
- [sl1user](docs/pkg/sl1user.md) - Package sl1user have the routines for /api/account.
- [sl1generics](docs/pkg/sl1generics.md) - Package sl1generics have the generic routines to be used throughout the sl1cmd.

# Installing

[Download](../../releases) the binary file. 

copy it to `/usr/local/bin`

```shell
sudo cp sl1cmd /usr/local/bin/.
```

# Getting Started

Before using the sl1cmd to perform admnistrative tasks you must configure the sl1cmd vault using [this procedure](docs/cmd/sl1cmd_vault_new.md)


 

    


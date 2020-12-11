<img src="docs/logo/logo.png" alt="twsgomon" title="twsgomon" align="right" width="60">

# >sl1cli_tools
[![Go Report Card](https://goreportcard.com/badge/github.com/marco-ostaska/sl1cli-tools)](https://goreportcard.com/report/github.com/marco-ostaska/sl1cli-tools)


`sl1cli-tools` is a set of command line interface tools that uses Science Logic Moniroting API to perform administrative tasks in Science Logic.

## Tools

### Vault
- [sl1-vault](docs/cmd/sl1-vault.md) - Configure or update credentials for sl1api to be used by `sl1cli-tools`

### User Administration
- [sl1-id](docs/cmd/sl1-id.md) - Print sl1 user information for the specified USERS
- [sl1-idinfo](docs/cmd/sl1-idinfo.md) - Print sl1 user information details for the specified USER
- [sl1-useradd](docs/cmd/sl1-useradd.md) - Delete the specified USER.
- [sl1-passwd](docs/cmd/sl1-passwd.md) - Create new USER on sl1.
- [sl1-userdel](docs/cmd/sl1-userdel.md) - Delete the specified USER.
## Init setup

After downloading and untar the binaries you will have to run [sl1-vault](docs/cmd/sl1-vault.md) to set up new credentials

## Packages

- [apicryptcfg](docs/pkg/apicryptcfg.md) - Package apicryptcfg crypt the config file used by sl1tools.
- [apirequest](docs/pkg/apirequest.md) - Package apirequest makes http request calls on sl1api
- [apipost](docs/pkg/apipost.md) - Package apipost post into sl1api
- [sl1user](docs/pkg/sl1user.md) - Package sl1user have the routines for /api/account
- [sl1generics](docs/pkg/sl1generics.md) -  Package sl1generics have the generic routines to be used throughout the sl1tools


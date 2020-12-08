# sl1cli-tools
[![Go Report Card](https://goreportcard.com/badge/github.com/marco-ostaska/sl1cli-tools)](https://goreportcard.com/report/github.com/marco-ostaska/sl1cli-tools)

`sl1cli-tools` is a set of CLI tools that uses Science Logic API to perform administrative tasks in Science Logic Monitoring tool

## Tools

- [sl1-vault](docs/cmd/sl1-vault.md) - Configure or update credentials for sl1api to be used by `sl1cli-tools`
- [sl1-id](docs/cmd/sl1-id.md) - Print sl1 user information for the specified USERS

## Init setup

After downloading and untar the binaries you will have to run [sl1-vault](docs/sl1-vault.md) to set up new credentials

## Packages

- [apicryptcfg](docs/pkg/apicryptcfg.md) - Package apicryptcfg crypt the config file used by sl1tools.
- [apirequest](docs/pkg/apirequest.md) - Package apirequest makes http request calls on sl1api
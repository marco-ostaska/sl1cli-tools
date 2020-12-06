# sl1-vault

## Usage

```shell
Usage: sl1-vault [OPTION]... [ARGUMENTS]
Configure or update credentials for sl1api

Options:
  -new     setup a new configuration
  -update  update user and Password to existing cofiguration
  -h       display this help and exit
  -v       display version

Arguments:
  -u      username
  -p      password
  -url    sl1 api URL
```

## Creating New vault

```
Usage: sl1-vault -new ... [ARGUMENTS]
Configure new credentials for sl1api. 
This can also be used to update all parameters to existing vault

Mandatory Arguments: 
  -u      username
  -p      password
  -url    sl1 api URL

Example:
  sl1-vault -new -u "myuser" -p "pass1234" -url "https://sl1api.com"
```

## Updating Credential for an existing vault

```
Usage: sl1-vault -update ... [ARGUMENTS]
Update existing credentials for sl1api

Mandatory Arguments: 
  -u      username
  -p      password

Example:
  sl1-vault -update -u "myuser" -p "pass1234"
```

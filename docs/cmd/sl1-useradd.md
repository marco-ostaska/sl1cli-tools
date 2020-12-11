# sl1-useradd

## Usage

```
Usage: sl1-useradd [OPTION]...
Create newuser on sl1

Mandatory Options:
  -user                  User name
  -email                 User email 
  -name                  User full name
  -org                   User Organization ID
  -userpolicy            User Policy ID

Options:
  -resetrequired         Password required to be changed on first login (0 or 1) (Default: 1)
  -admin                 Admin 0 or 1                                            (Default: 1)
  -permissionkeys        Permission Keys IDs separated by comma
  -alignedorgs           Aligned Organizations IDs separated by comma
  -h                     display this help and exit
  -v                     display version

Warning:
  Avoid using double quotes, use always single quotes on arguments

Example:
 sl1-useradd -org '2' -ermail 'teste@xx.com' -name 'teste' -admin '1' -userpolicy '3' -alignedorgs '0,2,3,5'
```

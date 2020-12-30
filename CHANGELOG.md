# Changelog

## [0.04 Unreleased] - 2020-12-30
### Added
- `sl1/vault`- Package vault manage encryption for sl1cmd credentials

### Changed
- `sl1/httpcalls`- change functions names, made minor improvements and moved to sl1 package
- vault cmd - made minor improvements. 

### Removed
- - `cryptcfg` - Package cryptcfg crypt the config file used by sl1cmd. It was incorporated in sl1/vault package.

## [0.02 Unreleased] - 2020-12-28
### Added
- Flag to accepted invalid certificated, before it always had insecure method.
- added version to control

### Fixed
- vault examples

## [Unreleased] - 2020-12-24
### Added
- `id`- print users sl1 id for the specified user.
- `userinfo` - print user information for the specified user.
- `useradd `-  create a new user.
- `userdel` - delete user account.
- `passwd` - change user password.
- `cryptcfg` - Package cryptcfg crypt the config file used by sl1cmd.
- `httpcalls` - Package httpcalls makes http request calls on sl1api.
- `sl1user` - Package sl1user have the routines for /api/account.
- `sl1generics` - Package sl1generics have the generic routines to be used throughout the sl1cmd
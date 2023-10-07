# versctl Version Tool (CLI)

<img src="./docs/imgs/icon.png" width="300" height="300" >

versctl Manages version control of applications based on the [Semantic Versioning Specification](https://semver.org/) and [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)

Table of Contents
=================
<!-- toc -->
- [Quick Start](#quick-start)
    - [Download Dependencies](#download-dependencies)
    - [Normal Build](#normal-build)
    - [Linux OS Build](#linux-build)
- [Increment Types](#increment-types)      
- [User Guide](#user-guide)
- [CI Examples Guide](#ci-examples)


<!-- /toc -->
## Quick Start

### Download Dependencies
````
$ go mod tidy
````
### Normal Build
````
$ go build -v -o versctl
````
### Linux Build
````
$ env GOOS=linux GOARCH=amd64 go build -v -o versctl
````

## Increment Types

| Increment Type | Version Format | Example |
| -------------- | -------------- | ------- |
| Semantic Version  | `major.minor.patch` | 1.0.0 | 
| date | `year.month.day.count` | 1979.05.11.0 |
| release candidate | `rc` | 1.0.0-rc |
| release candidate with semantic version| `rc:(major\|minor\|patch)` | 1.1.0-rc |
| staging | `staging` | 1.0.0-staging | 

## User Guide
See the [user guide](docs/user-guide.md) in the `/docs` directory.

## CI Examples
See the [ci examples guide](docs/ci-examples-guide.md) in the `/docs` directory.

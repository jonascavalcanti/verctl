# User Guide

Starting with xversioner

## CLI Options
The xversioner has many CLI options that can be used to override its default behavior.
```
xversioner --help
Manages version control of applications based on the `Semantica Version Specification` and `Commit Convention`

Usage: 
  xversioner [command] [flags]

Available Commands:
    get		get version based specifc file
    update	Update application version with ou without a file

Flags:
    filepath,-f	        path of the file where contains the version
    version,--version   string version
    increment,-i	    increment type [(major|minor|patch), date, (rc|rc:(major|minor|patch)), staging]

Use " [command] --help" for more information about a command.

```

## update command
### Increment Types

*Semantic versioning*
- Using filepath
````
$ xversioner update -f ./app/app.properties -i patch
````
- Using version
````
$ xversioner update --version 1.0.0 -i patch
>> OUTPUT: 1.0.1
````

*Date*
- Using filepath
````
$ xversioner update -f ./app/app.properties -i date
>> OUTPUT: 2022.05.30.1
````
- Using version
````
$ xversioner update --version 2022.05.30.0 -i date
>> OUTPUT: 2022.05.30.1
````

*Release Candidate*
- Using filepath
````
$ xversioner update -f ./app/app.properties -i rc
$ xversioner update -f ./app/app.properties -i rc:minor
````
- Using version
````
$ xversioner update --version 1.0.0 -i rc
>> OUTPUT: 1.0.0-rc
$ xversioner update --version 1.0.0 -i rc:minor
>> OUTPUT: 1.1.0-rc
````

*Staging*
- Using filepath
````
$ xversioner update -f ./app/app.properties -i staging
````
- Using version
````
$ xversioner update --version 1.0.0 -i staging
>> OUTPUT: 1.0.0-staging
````

## get command



- Using filepath
````
$ xversioner get -f ./app/app.properties
````

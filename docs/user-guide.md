# User Guide

Starting with versctl

## CLI Options
The versctl has many CLI options that can be used to override its default behavior.
```
versctl --help
Manages version control of applications based on the `Semantica Version Specification` and `Commit Convention`

Usage: 
  versctl [command] [flags]

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
$ versctl update -f ./app/app.properties -i patch
````
- Using version
````
$ versctl update --version 1.0.0 -i patch
>> OUTPUT: 1.0.1
````

*Date*
- Using filepath
````
$ versctl update -f ./app/app.properties -i date
>> OUTPUT: 2022.05.30.1
````
- Using version
````
$ versctl update --version 2022.05.30.0 -i date
>> OUTPUT: 2022.05.30.1
````

*Release Candidate*
- Using filepath
````
$ versctl update -f ./app/app.properties -i rc
$ versctl update -f ./app/app.properties -i rc:minor
````
- Using version
````
$ versctl update --version 1.0.0 -i rc
>> OUTPUT: 1.0.0-rc
$ versctl update --version 1.0.0 -i rc:minor
>> OUTPUT: 1.1.0-rc
````

*Staging*
- Using filepath
````
$ versctl update -f ./app/app.properties -i staging
````
- Using version
````
$ versctl update --version 1.0.0 -i staging
>> OUTPUT: 1.0.0-staging
````

## get command



- Using filepath
````
$ versctl get -f ./app/app.properties
````

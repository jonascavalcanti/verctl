# xversioner (CLI): xversioner controls application version

## Using tool

### Download dependencies
````
$ go mod tidy
````
### Normal Build
````
$ go build -v -o xversioner
````
### Build to Linux OS
````
$ env GOOS=linux GOARCH=amd64 go build -v -o xversioner
````

### Get aplication version example
````
$ xversioner get filepath ./app/file_contain_app_version
or 

$ xversioner get -f ./app/file_contain_app_version
````

### Update aplication version example
````
$ xversioner update filepath ./app/file_contain_app_version increment patch
or
$ xversioner update -f ./app/file_contain_app_version -i patch

````

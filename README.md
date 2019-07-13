# SBLM CLI
Spring Boot Logger Manager CLI.

![Screenshot](https://raw.github.com/datoga/sblm-cli/master/images/screenshot1.PNG "Screenshot list")

# Usage

In general, it is encouraged to follow the guidelines of the CLI.

```console
foo@bar:~$ sblm-cli -h

This CLI allows to users to list or edit the Spring Boot logger from any online Spring Boot application.

Usage:
  sblm-cli [command]

Available Commands:
  edit        It edits the logger level of a logger or set of loggers
  help        Help about any command
  list        List all loggers from the APP

Flags:
  -a, --app string       Application to get information (default "HelloWorld")
  -h, --help             help for sblm-cli
  -p, --pattern string   Pattern to be applied (default "$server/$app/actuator/loggers")
  -s, --server string    Server to get information (default "http://localhost:8080")
  -v, --verbose          verbose output

Use "sblm-cli [command] --help" for more information about a command.

```

## List loggers
To list all the available loggers
```console
foo@bar:~$ sblm-cli list --server <YOUR_SERVER> --app <YOUR_APP> --filter <YOUR_FILTER>
```


This command makes, by default, a GET request to:
```console
<YOUR_SERVER>/<YOUR_APP>/actuator/loggers
```
For example:
```console
http://localhost:8080/HelloWorld/actuator/loggers
```

The result is filtered by the group provided.

## Edit loggers
To edit a group of the available loggers
```console
foo@bar:~$ sblm-cli edit --server <YOUR_SERVER> --app <YOUR_APP> --name <YOUR_LOGGER_GROUP> --level <OFF|ERROR|WARN|INFO|DEBUG|TRACE>
```

This commands makes, by default, a POST request to:
```console
<YOUR_SERVER>/<YOUR_APP>/actuator/loggers/<YOUR_LOGGER_GROUP>
```
For example:
```console
http://localhost:8080/HelloWorld/actuator/loggers/org.springframework.web.server
```

Aditionally sends, by POST parameter a JSON like that:
```json
"org.springframework.web.server": {
    "configuredLevel": "DEBUG"
}
```

## Tricks
- In the logger list, if you need to do some grep with the data returned by the command please use the raw (-r or --raw) option to get a more flat list.
- If you need more debugging information about the internal use of the tool, please use the verbose (-v or --verbose) option.
- If you are working on a local environment probably you won't have any context path. In this case you could change the endpoint format setting the pattern (-p or --pattern) option to something like ```$server/actuator/loggers```.

# Building SBLM-CLI

Just execute ```go build``` on the root path and get the binary!

# Architecture

Sorry, this software has been done in only one day! No testing, no good architecture, ... maybe in future versions.

# License

MIT License applies.

Copyright 2019 - David Torres.

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
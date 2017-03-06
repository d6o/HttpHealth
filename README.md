# ![HttpHealth](http://image.prntscr.com/image/2bcc500477c14761abb9c4862b2a3abc.png)

# HttpHealth ![Language Badge](https://img.shields.io/badge/Language-Go-blue.svg) ![Go Report](https://goreportcard.com/badge/github.com/DiSiqueira/HttpHealth) ![License Badge](https://img.shields.io/badge/License-MIT-blue.svg) ![Status Badge](https://img.shields.io/badge/Status-Beta-brightgreen.svg)

The HttpHealth's goal is to be a perfect tool providing a stupidly easy-to-use and fast program to ping all your servers in just one route.

## Project Status

HttpHealth is on beta. Pull Requests [are welcome](https://github.com/DiSiqueira/HttpHealth#social-coding)

## Features

- It's perfect to you monitor your services status
- CUSTOMIZE to your needs
- EASY to add servers
- REST API
- Admin API
- Easy to use
- STUPIDLY [EASY TO USE](https://github.com/DiSiqueira/HttpHealth#usage)
- Very fast start up and response time
- Uses native libs

## Installation

### Option 1: Go Get

```bash
$ go get github.com/DiSiqueira/HttpHealth
```

### Option 2: From source

```bash
$ git clone https://github.com/DiSiqueira/HttpHealth.git
$ cd HttpHealth/
$ go build *.go
```

## Usage

### Simple Start ###
```bash
$ HttpHeatlh
```

### Add a server to be pinged

```bash
$ curl -X POST --url "http://localhost:8001/url" --data "url=www.google.com"
```

### Ping all servers

```bash
$ curl -X GET --url "http://localhost:8000/"
```

![Ping all servers](http://image.prntscr.com/image/0b11f487a2ac42d7bf4eb4e940e147ff.png)

### Delete a server from the list

```bash
$ curl -X DELETE --url "http://localhost:8001/url/www.google.com"
```

## Program Output

![Ping all servers](http://image.prntscr.com/image/0b11f487a2ac42d7bf4eb4e940e147ff.png)

## Contributing

### Bug Reports & Feature Requests

Please use the [issue tracker](https://github.com/DiSiqueira/HttpHealth/issues) to report any bugs or file feature requests.

### Developing

PRs are welcome. To begin developing, do this:

```bash
$ git clone --recursive git@github.com:DiSiqueira/HttpHealth.git
$ cd HttpHealth/
$ go run *.go
```

## Social Coding

1. Create an issue to discuss about your idea
2. [Fork it] (https://github.com/DiSiqueira/HttpHealth/fork)
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Commit your changes (`git commit -am 'Add some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create a new Pull Request
7. Profit! :white_check_mark:

## License

The MIT License (MIT)

Copyright (c) 2013-2017 Diego Siqueira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
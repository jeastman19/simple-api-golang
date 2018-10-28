# simple-api-golang
Basic example of a RESTful API with golang

This example shows a very simple RESTful API developed with Go (golang) in order to show how client requests are received and how they respond to it.

This example is a simple CRUD with data stored in an array.

## Usage
To test this example, download the code or clone it.

``` bash
git clone https://github.com/jeastman19/simple-api-golang.git
```

Then, from the directory where the project was cloned, compile and execute

``` bash
$ cd simple-api-golang
$ go build main.go
$ ./main
```

This will compile and execute the example that you will hear by port 3000

To see it running, you can use Postman.

## Routes

The base URL is http: // localhost: 3000


Method|Endpoint|Purpose
--|--|--
GET|/people|Get all People
GET|/people/{id}|Get a Person
POST|/people|Create a Person
DELETE|/people/{id}|Delete a Person


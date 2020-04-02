# Simple Rest Api with Golang

## Installation

- Clone this repo, outside your `$GOPATH` because we use [go modules][gomod]

[gomod]: https://blog.golang.org/using-go-modules

## How this app working ?
- Open your Command Prompt and go to root of this app directory
- Run this command : `go run main.go`
- Test the app with postman, Insomnia, HTTPie etc.
- The following is an example of data to test this app :
    
    **Method POST**
    ```
    [
        {
            "name": "Alpha",
            "message": "lorem ipsum"
        }, {
            "name": "Bravo",
            "message": "lorem ipsum"
        }, {
            "name": "Charlie",
            "message": "lorem ipsum"
        }
    ]
    ```
- Check the list of route on `app/app.go` to test all endpoint on this app.

I just want to remind you to test the endpoint, according to the action `(POST, GET, OPTIONS)` you want to do.

## Thank You
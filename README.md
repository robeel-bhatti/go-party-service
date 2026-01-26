# go-party-service

This is a REST microservice built in Golang.

I built this because I want to get experience building microservices in Golang.

Golang's philosophy of making code explicit (ex. error handling) and its simple syntax made this fun to write
and quick to iterate on. There is also less dependency on 3rd party libraries. I found this to be a good mix
of making code explicit and provided abstractions where necessary to decouple app components.

Also enjoyed Golang's packaging approach. Imports are easy to do as you just import the package name into a Go file.
Then you have access to all the Go code in that package, no matter how many files are in that package. But 
still be cautious on avoiding cyclic imports. 

## Architecture

Repository Pattern - All data access has been abstracted away under a single package.

Dependency Injection - Structs have their dependencies provided to them externally. No hard-coding dependency creation in the struct itself.


## Running Locally

Run `bash testing.sh` in your terminal. That will run docker containers that contain app deps (redis, postgres etc.).
It will then start the Go application in port `8080`. It will prompt for an ID to retrieve. Enter and view result.
Once result is returned app and its deps will be torn down. 
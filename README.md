AccelOne - Go API challenge

The template was built in [GO](https://github.com/golang/go) programming language in 1.21 version, using the [Echo Context](https://github.com/labstack/echo) web framework.

It is following a [Domain-Driven Design](https://en.wikipedia.org/wiki/Domain-driven_design) architecture due to its simplicity of maintenance, understandability and the easy way to write tests.  
This project is divided in four layers:
* Handler: First layer, where the request begins its journey. You can find the bindings of the ULR, query parameters and their validations located in `delivery/rest` directory.
* Usecase: Second layer, where you can make the business logic. In this case there is no business logic due to the simplicity of the API. Located in `usecase` directory.
* Repository: Third layer who is in charge of connect and get/save data from/to the database (memory in this case). Located in `repository` directory.
* Model: Fourth and last layer where all models should be defined. Located in `model` directory

### Validator

Project was designed to validate any struct with [Go Package Validator](https://pkg.go.dev/github.com/go-playground/validator/v10)
which basically implements value validations for structs and individual fields based on tags.  
Tags are defined in `models/contact.go` inside Contact struct and the call to validator is in `delivery/rest/contact_handler.go` inside `AddContact()` and `PutContact()` functions.  
Hint: if you want to test it, just call to POST/PUT request and remove any of the fields of the body

### Tests

Project doesn't have tests due to wasn't required, but Handler and Usecase layers are prepared for unit tests thanks to their interface definitions.
I would use a mock generator tool called [Mockery](https://vektra.github.io/mockery/latest/) which you must install if you want to use it.

In addition to unit tests, project is also prepared for integration tests thanks to the Repository layer. It only needs to define a new `memoryDB` to save all necessary of the integration tests.

### Error handler

Errors are being handled by the [Echo Context web framework](https://github.com/labstack/echo) following its own [error structure](https://echo.labstack.com/docs/error-handling)  

## How to

To run and test this project, you have to execute the following command in root directory:
```
go run ./cmd/app
```
and will see the following message in console
```
http server started on [::]:8080
```

Once you see that, you are ready to play with this API.  
All the endpoints are defined in the following [Postman Collection](https://www.postman.com/dvillarruel/workspace/accelone/collection/4793868-15064db0-9aa1-4f62-91ee-5336825d5d08?action=share&creator=4793868)

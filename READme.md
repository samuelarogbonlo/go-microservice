First, created webserver with go and set up error message. http.handlerequest helps to convert func passing as parameter into http handler and registers it with the server. The handler is an interface. To create handler, you can create a struct. 
Second, refactoring of code incase of need for testing and we start with making a part of the fuction an "object". create a struct next  
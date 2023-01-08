
#### Introduction:
This application defines three microservices, ```/service1```, ```/service2```, and ```/service3```, which are implemented as HTTP handlers. 

```/service1``` and ```/service2``` both make HTTP requests to the other services using the http package. When you access ```/service1```, it will make a request to ```/service2```, which in turn will make a request to ```/service3```. 

The final response will be the concatenation of the responses from all three services.

#### How to run:
To run this application, you will need to have the ```Go runtime``` installed. You can install ```Go``` from the official website: https://golang.org/. 

Then, you can build and run the application using the following commands:

> go build app.go
./app

This will start the HTTP server, and you can access the microservices at 
```http://localhost:8000/service1```
```http://localhost:8000/service2```

```http://localhost:8000/service3```
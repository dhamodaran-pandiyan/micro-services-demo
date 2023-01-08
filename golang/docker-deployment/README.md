
#### Introduction:
This Dockerfile is based on the ```golang:latest``` image, which includes the Go runtime and tools. It copies the source code for the application into the image and builds it using the go build command. 

It exposes port ```8000```, which is the port that the application listens on, and runs the app executable as the entrypoint for the container.

#### How to run:
To build the Docker image, you will need to have ```Docker``` installed. You can install Docker from the official website: ```https://www.docker.com/```. Then, you can build the image using the following command:
> docker build -t go-app .

This will build the Docker image and tag it as g```o-app```. You can then run the image using the following command:
> docker run -p 8000:8000 go-app


This will start the HTTP server, and you can access the microservices at 
    ```http://localhost:8000/service1```
    ```http://localhost:8000/service2```
    ```http://localhost:8000/service3```
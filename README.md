# go-health

### Experimental

Building a real time application health monitoring prototype using Golang.


A health check in microservices is a mechanism that ensures each service is functioning correctly and is available. It typically involves periodically checking the status of various components of a service and reporting their health.


Aims:
Monitor web applications, microservices on Kubernetes, real time systems, MQ/Kafka listener applications

### Important things
API health endpoints should be: `secure, disabled cache, responsive`

observability
elasticity
availability


## Used:
    - net/http
    - gorilla/mux
    
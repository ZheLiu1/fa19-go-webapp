## Team info

|Name           |NUID     |
|---------------|---------|
|Yunlu Liaozheng|001470321|
|Feng Huang     |001230993|
|Zhichao Pan    |001493794|
|Zhe Liu        |001475826|

## Instruction

### Requirement

* `Docker` installed
* `golang >= 1.12` installed

### Run server locally

Use `make` to build the project.

Then use `bin/webapp` to run the server.

### Build docker image

Use `make docker` to build the image as `webapp:latest`. You can also build with your own tag by `docker build -t tag:version .`

### Run the container

Use `sudo docker run -idt -p 11080:8080 webapp` to run the server in a container. 



# Gocker (I'm not proud of the name myself): A lightweight container built in go

## Usage

The project can only be built and run on an ubuntu machine. So to run it, it has to be containerized in an Ubuntu container.

---

To run this project:

- Build the container by running in the terminal `make build_docker`

- Run the container using `make`

You might think this is all but it gets better 🚀. You still have to run my project (another container) within the docker container.

- So to run our project (another container) such that you remain inside it and can do whatever you want in an isolated environment run. `./main run /bin/bash`. This opens a terminal session within our project.

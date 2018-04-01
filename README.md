# gRPC guruguru

This is a useless gRPC server/client process model. This project is for my gRPC training.

First, a boss process waits for member processes joining. Once all members joined the boss forms cyclic linked list of members by setting next process to each members.  Messages sent to a member loop around the members list infinitely.

Member processes are written by different language for each other. These are based on the same proto file as you can see. Used languages are below:

- Go
- Python
- JavaScript(Node.js)
- Ruby
- ...

## How to use

```console
$ make build # for first
$ make up
```


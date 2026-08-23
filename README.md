# go-examples

Personal exercises working through [Go by Example](https://gobyexample.com), one topic per directory.

## Structure

Each topic gets its own numbered directory with its own `main.go`, matching the order on
gobyexample.com:

```
01-hello-world/
02-values/
03-variables/
...
```

## Running an exercise

```bash
go run ./02-values
```

## Progress

- [x] Hello World
- [x] Values
- [ ] Variables
- [ ] Constants
- [ ] For
- [ ] If/Else
- [ ] Switch
- [ ] Arrays
- [ ] Slices
- [ ] Maps
- [ ] Range
- [ ] Functions
- [ ] Multiple Return Values
- [ ] Variadic Functions
- [ ] Closures
- [ ] Recursion
- [ ] Pointers
- [ ] Strings and Runes
- [ ] Structs
- [ ] Methods
- [ ] Interfaces
- [ ] Struct Embedding
- [ ] Generics
- [ ] Errors
- [ ] ...and beyond

To add a new topic: create a new numbered directory with a `main.go`, write/adapt the example,
run it with `go run ./NN-topic-name`, check the box above, commit.

## Syncing across machines

This repo is on GitHub so it can be pulled on other machines (e.g. a Mac mini):

```bash
git clone https://github.com/<your-username>/go-examples.git
cd go-examples
go run ./01-hello-world
```

Push after each session so the other machine can `git pull` and pick up where you left off.

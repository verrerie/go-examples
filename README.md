# go-examples

Personal exercises working through [Go by Example](https://gobyexample.com), one topic per directory.

## Structure

Each topic gets its own numbered directory with its own `main.go`, matching the order on
gobyexample.com:

```
001-hello-world/
002-values/
003-variables/
...
```

## Running an exercise

```bash
go run ./002-values
```

## Progress

Complete — every gobyexample topic is covered across the 81 directories below,
finishing at their last topic (Exit).

Five topics share `011-functions`, which demonstrates all of them in one file.
That merge is why the directory numbers run four behind gobyexample's own
ordering from `012-range-over` onward, so use the number in this list to find the
directory rather than counting positions on the site.

- [x] `001` Hello World
- [x] `002` Values
- [x] `003` Variables
- [x] `004` Constants
- [x] `005` For
- [x] `006` If/Else
- [x] `007` Switch
- [x] `008` Arrays
- [x] `009` Slices
- [x] `010` Maps
- [x] `011` Functions
- [x] `011` Multiple Return Values *(in `011-functions`: `div`)*
- [x] `011` Variadic Functions *(in `011-functions`: `concat`)*
- [x] `011` Closures *(in `011-functions`: `initSeq`)*
- [x] `011` Recursion *(in `011-functions`: `c`)*
- [x] `012` Range over Built-in Types
- [x] `013` Pointers
- [x] `014` Strings and Runes
- [x] `015` Structs
- [x] `016` Methods
- [x] `017` Interfaces
- [x] `018` Enums
- [x] `019` Struct Embedding
- [x] `020` Generics
- [x] `021` Range over Iterators
- [x] `022` Errors
- [x] `023` Custom Errors
- [x] `024` Goroutines
- [x] `025` Channels
- [x] `026` Channel Buffering
- [x] `027` Channel Synchronization
- [x] `028` Channel Directions
- [x] `029` Select
- [x] `030` Timeouts
- [x] `031` Non-Blocking Channel Operations
- [x] `032` Closing Channels
- [x] `033` Range over Channels
- [x] `034` Timers
- [x] `035` Tickers
- [x] `036` Worker Pools
- [x] `037` WaitGroups
- [x] `038` Rate Limiting
- [x] `039` Atomic Counters
- [x] `040` Mutexes
- [x] `041` Stateful Goroutines
- [x] `042` Sorting
- [x] `043` Sorting by Functions
- [x] `044` Panic
- [x] `045` Defer
- [x] `046` Recover
- [x] `047` String Functions
- [x] `048` String Formatting
- [x] `049` Text Templates
- [x] `050` Regular Expressions
- [x] `051` JSON
- [x] `052` XML
- [x] `053` Time
- [x] `054` Epoch
- [x] `055` Time Formatting / Parsing
- [x] `056` Number Parsing
- [x] `057` Random Numbers
- [x] `058` URL Parsing
- [x] `059` SHA256 Hashes
- [x] `060` Base64 Encoding
- [x] `061` Reading Files
- [x] `062` Writing Files
- [x] `063` Line Filters
- [x] `064` File Paths
- [x] `065` Directories
- [x] `066` Temporary Files and Directories
- [x] `067` Embed Directive
- [x] `068` Testing and Benchmarking
- [x] `069` Command-Line Arguments
- [x] `070` Command-Line Flags
- [x] `071` Command-Line Subcommands
- [x] `072` Environment Variables
- [x] `073` Logging
- [x] `074` HTTP Client
- [x] `075` HTTP Server
- [x] `076` TCP Server *(not a gobyexample topic — added while working through `net`)*
- [x] `077` Context
- [x] `078` Spawning Processes
- [x] `079` Exec'ing Processes
- [x] `080` Signals
- [x] `081` Exit

To add a topic: create a directory at the next free number with a `main.go`, write
the example, run it with `go run ./NNN-topic-name`, tick the box above, commit.

## Syncing across machines

This repo is on GitHub so it can be pulled on other machines (e.g. a Mac mini):

```bash
git clone https://github.com/<your-username>/go-examples.git
cd go-examples
go run ./001-hello-world
```

Push after each session so the other machine can `git pull` and pick up where you left off.

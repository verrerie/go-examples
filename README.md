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

Complete — every gobyexample topic is covered across the 81 directories below,
finishing at their last topic (Exit).

Five topics share `11-functions`, which demonstrates all of them in one file.
That merge is why the directory numbers run four behind gobyexample's own
ordering from `12-range-over` onward, so use the number in this list to find the
directory rather than counting positions on the site.

- [x] `01` Hello World
- [x] `02` Values
- [x] `03` Variables
- [x] `04` Constants
- [x] `05` For
- [x] `06` If/Else
- [x] `07` Switch
- [x] `08` Arrays
- [x] `09` Slices
- [x] `10` Maps
- [x] `11` Functions
- [x] `11` Multiple Return Values *(in `11-functions`: `div`)*
- [x] `11` Variadic Functions *(in `11-functions`: `concat`)*
- [x] `11` Closures *(in `11-functions`: `initSeq`)*
- [x] `11` Recursion *(in `11-functions`: `c`)*
- [x] `12` Range over Built-in Types
- [x] `13` Pointers
- [x] `14` Strings and Runes
- [x] `15` Structs
- [x] `16` Methods
- [x] `17` Interfaces
- [x] `18` Enums
- [x] `19` Struct Embedding
- [x] `20` Generics
- [x] `21` Range over Iterators
- [x] `22` Errors
- [x] `23` Custom Errors
- [x] `24` Goroutines
- [x] `25` Channels
- [x] `26` Channel Buffering
- [x] `27` Channel Synchronization
- [x] `28` Channel Directions
- [x] `29` Select
- [x] `30` Timeouts
- [x] `31` Non-Blocking Channel Operations
- [x] `32` Closing Channels
- [x] `33` Range over Channels
- [x] `34` Timers
- [x] `35` Tickers
- [x] `36` Worker Pools
- [x] `37` WaitGroups
- [x] `38` Rate Limiting
- [x] `39` Atomic Counters
- [x] `40` Mutexes
- [x] `41` Stateful Goroutines
- [x] `42` Sorting
- [x] `43` Sorting by Functions
- [x] `44` Panic
- [x] `45` Defer
- [x] `46` Recover
- [x] `47` String Functions
- [x] `48` String Formatting
- [x] `49` Text Templates
- [x] `50` Regular Expressions
- [x] `51` JSON
- [x] `52` XML
- [x] `53` Time
- [x] `54` Epoch
- [x] `55` Time Formatting / Parsing
- [x] `56` Number Parsing
- [x] `57` Random Numbers
- [x] `58` URL Parsing
- [x] `59` SHA256 Hashes
- [x] `60` Base64 Encoding
- [x] `61` Reading Files
- [x] `62` Writing Files
- [x] `63` Line Filters
- [x] `64` File Paths
- [x] `65` Directories
- [x] `66` Temporary Files and Directories
- [x] `67` Embed Directive
- [x] `68` Testing and Benchmarking
- [x] `69` Command-Line Arguments
- [x] `70` Command-Line Flags
- [x] `71` Command-Line Subcommands
- [x] `72` Environment Variables
- [x] `73` Logging
- [x] `74` HTTP Client
- [x] `75` HTTP Server
- [x] `76` TCP Server *(not a gobyexample topic — added while working through `net`)*
- [x] `77` Context
- [x] `78` Spawning Processes
- [x] `79` Exec'ing Processes
- [x] `80` Signals
- [x] `81` Exit

To add a topic: create a directory at the next free number with a `main.go`, write
the example, run it with `go run ./NN-topic-name`, tick the box above, commit.

## Syncing across machines

This repo is on GitHub so it can be pulled on other machines (e.g. a Mac mini):

```bash
git clone https://github.com/<your-username>/go-examples.git
cd go-examples
go run ./01-hello-world
```

Push after each session so the other machine can `git pull` and pick up where you left off.

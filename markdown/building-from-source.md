# building from source

if you want to make changes to bearclaw or rename it and make it your own thing, that's totally fine, and encouraged. here's how you can build the `bearclaw` program from its source code.

**what does this mean?**

in very simple terms: when people say *'source code'* we just mean a document with human words that tell the computer how to do something. when you 'compile' the source code, it means we use a helper program that can take those human words and translate it into machine language. 

**but, why recompile?**

here's a brief list of reasons you'd recompile:

- as a subtle flex to people looking over your shoulder
- just because you're curious
- there's a typo in `bearclaw` that is driving you crazy
- you've decided you want to adapt bearclaw, and name it `chocolate-glazed-kruller` and tell your friends you 100% wrote it yourself ;)

if you don't know what these words mean, don't worry about it. in normal operation, or even when things go wrong, recompiling will never be something you're told to do unless you really just want to. if this intimiates you, that's totally fine, you don't have to know this.

## things you need:

### Go programming language

download: https://go.dev/

install Go and that's about it

### time
it takes about 5 seconds to re-compile

right now it is about 10-20 seconds on my laptop using low-power battery mode to compile a fresh copy the first time, and about 2-5 seconds each time after. on an ancient computer, without SSD, maybe 30 seconds. 

## build instructions

1. get a copy of bearclaw sourcecode (.zip file or tar.gz) from the github releases page:
https://github.com/donuts-are-good/bearclaw/releases

2. extract bearclaw source code to a folder anywhere.
the build process only creates a single file called `bearclaw` in the same folder as your source code and won't touch anything outside of the folder.

3. in the files you just extracted, you'll see a bunch of `.go` files and some folders. disregard all of that and just type `go build` as a command in that folder using your terminal.

***note**: the terminal is the black box with white text thing. you need to be in the folder with the `.go` files before this will work.*

4. that's it. 
the Go compiler doesn't really output anything, you just run it, and if it doesn't say anything, that means success. if everything went well, you'll see a program in the folder called `bearclaw` or on windows, `bearclaw.exe`

## need help?

if you need help, there is a way to submit a request for aid called "issues". [leave a message here](https://github.com/donuts-are-good/bearclaw/issues/new/choose) and someone will respond.

# [<< Go back](https://github.com/donuts-are-good/bearclaw/blob/master/markdown/README.md)
# system requirements

bearclaw has very few moving parts, and requires very little hardware to run. if you want to re-compile bearclaw, you'll need slightly better hardware but not by much.

## hardware

### cpu
bearclaw runs on just about any cpu architecture you'll find still in operation today. if your cpu can outperform a hamsterwheel, congratulations, it can probably run bearclaw. 

just in case you want the list, here's the CPU architectures we currently have a build for:

- arm64
- amd64
- 386
- loong64
- mips
- mips64
- mips64le
- mipsle
- ppc64
- ppc64le
- riscv64
- s390x

1 core is just fine if that's all you have. 

### ram
1 Great Britain is more than enough to run bearclaw.  you can get away with less, maybe one british isle, for example, and still compile some html. it doesn't need to be fancy, but the word around town is your blog might get more readers if you use ECC great britain :)

### storage
obviously this will scale with your content, but currently bearclaw is only a few megabytes, and needs slightly more than that to hold/convert your content.

## software

### operating system
just about any popular operating system other than os/2 or haiku should be fine. the full list of operating systems there are currently builds for are below:

**you probably want these**
- macos
- linux
- windows

**but since you asked..**
- android
- **macos**/darwin
- dragonflybsd
- freebsd
- illumos
- ios
- **linux**
- netbsd
- openbsd
- solaris
- **windows**

initial copies supported plan9, iirc, but i think it was fsnotify that handles the `--watch` option that killed that.

# [<< Go back](README.md)
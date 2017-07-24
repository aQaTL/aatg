# ASCII Art Text Generator

A console app for generating ASCII art from normal text. 

## Downloading

To get the app go to the [Releases](https://github.com/aQaTL/ASCII-Art-Text-Generator/releases) 
page and download the latest version. Then, unpack the archive and enjoy!.

The other way is to download the repo by the `go get` tool and compile it by yourself. 
You'll need a [Go](https://golang.org/) language installed and configured (set GOPATH). 

```
go get github.com/aQaTL/ASCII-Art-Text-Generator
cd $GOPATH
go install github.com/aQaTL/ASCII-Art-Text-Generator
```

## Running the app

By default, it runs in the interactive mode - you can just type stuff and see the result 
being generated on the fly.

But you don't have to enter the interactive mode it you, for example, have a prepared text 
saved. To see other options, launch the app with a `--help` parameter: 

```
> ASCII\ Art\ Text\ Generator --help
Usage of ASCII Art Text Generator:
  -c string
        Custom ascii character (default "█")
  -i string
        Input file
  -p    Read from stdin
  -t string
        Specify custom text
```

You can use `-c` parameter combined with any other. Other parameters don't work together.

## Interactive mode

<script type="text/javascript" src="https://asciinema.org/a/2ALkXusauFvJ4UXT33IRzipzx.js" id="asciicast-2ALkXusauFvJ4UXT33IRzipzx" async></script>

# Other modes

<script type="text/javascript" src="https://asciinema.org/a/X5PItTdvrbbgeEH1w3HmV6uOU.js" id="asciicast-X5PItTdvrbbgeEH1w3HmV6uOU" async></script>
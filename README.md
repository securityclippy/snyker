## Snyker

A golang CLI and library for interacting with the Snyk API



:warning:  This is a WIP and will change a lot.  Feel free to use it, but
I make no guarantees about any of the interfaces :warning:  


Tests are non-existant right now and will appear as time permits


## Basic Usage:

```
snyker -h
An easy way to interact with the snyk api via the command line.  THis is a WIP. 

NOTE: Currently pulls its access token from the SNYK_TOKEN env var

Usage:
  snyker [command]

Available Commands:
  help        Help about any command
  orgs        A brief description of your command
  projects    A brief description of your command

Flags:
      --config string   config file (default is $HOME/.snyker.yaml)
  -h, --help            help for snyker
  -o, --output string   output type (default "json")
  -t, --toggle          Help message for toggle

Use "snyker [command] --help" for more information about a command.
```
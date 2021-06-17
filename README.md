# pdr

pyama docker revolution.

## install
```
$ brew tap pyama86/homebrew-ptools 
$ brew install pdr
```
## configure
- ~/.pdr

```yaml
repos:
  STNS:
    path: "~/src/github.com/STNS/STNS"
  libnss:
    path: "~/src/github.com/STNS/STNS"
    depends:
      - STNS
```

## usage

```
This is docker command wrapper CLI.

Usage:
  pdr [command]

Available Commands:
  down        run docker-compose down on each repository dir
  help        Help about any command
  log         Show log in docker container
  login       Show login to docker container
  ps          run docker-compose ps on each repository dir
  up          run docker-compose up on each repository dir

Flags:
      --config string   config file (default is $HOME/.pdr) (default "/Users/pyama/.pdr")
  -h, --help            help for pdr

Use "pdr [command] --help" for more information about a command.
```

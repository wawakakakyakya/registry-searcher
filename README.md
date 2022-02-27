# registry-searcher


## what's this ?
client for docker-registry

## how to use

1. install golang 1.17.7
1. git clone this repo
1. build client on clone dir
    ```
    make build_linux
    ```
1. help
    ```
    registry-searcher:~/src/registry-searcher$ ./registry-searcher help
    Usage: registry-searcher <flags> <subcommand> <subcommand args>

    Subcommands for List command register:
            list-image       get image list

    registry-searcher:~/src/registry-searcher$ 
    ```
1. list image
    ```
    $ ./registry-searcher list-image -address example.com -port 5000
    {"repositories":["container-image"]}
    
    $ 
    ```
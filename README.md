# psql-test-conn
* A tool for test postgres connection.
* Fast Begin
```shell
    # install
    go install github.com/VarusHsu/psql-test-conn@latest
    # append to env (optional)
    export GOROOT=~/go
    export PATH=$PATH:$GOROOT/bin
    # usage
    psql-test-conn -h localhost -P 5432 -u postgres -p 123
```
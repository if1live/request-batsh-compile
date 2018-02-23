# request-batsh-compile
compile batsh source in http://batsh.org

## install

```
go get github.com/if1live/request-batsh-compile
```

## usage

compile to bash

```bash
request-batsh-compile -target bash -src testdata/sample-success.batsh
```

compile to windows batch

```bash
request-batsh-compile -target winbat -src testdata/sample-success.batsh
```

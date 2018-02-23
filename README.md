# request-batsh-compile
use http://batsh.org, to compile [batsh language][repo-batsh] source file.

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


[repo-batsh]: https://github.com/BYVoid/Batsh
[site-batsh]: http://batsh.org

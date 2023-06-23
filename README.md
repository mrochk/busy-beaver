# Busy Beaver
Simplest 2-state Turing machine in Go to experiment on the [busy beaver problem](https://en.wikipedia.org/wiki/Busy_beaver).

It is as an appendix for a small [article](article.pdf) (in french) that I wrote as part of an assignment for my university.

The machines are described in [text-files](./machines) with the following structure, a matrix with $2|States|$ rows:
| _Current State_ | _Symbol Read_ | _Symbol Written_ | _Direction_ | _Next State_ |
|:---------------:|:-------------:|:----------------:|:-----------:|:------------:|
|        $0, 1, ..., \|States\|$        |       0\|1       |        0\|1       |     L\|R     |      $0, 1, ..., \|States\|$     |
|        ...       |       ...       |        ...       |     ...     |      ...    |


To try it, just install go and run:
```
go build
./busy-beaver
```

# wish-bubble-exec

An attempt at creating an interactive program exposed over SSH

## Try it out

This opens up an instance of `$EDITOR`, and works as expected

```bash
go run . -direct
```

---

This does not work at all...

The desired result is opening up `$EDITOR` within the SSH session, fully
interactive and everything. Doesn't work at all, and I've got no clue why.

First terminal

```bash
go run .
```

Second terminal

```bash
$ ssh -p 23553 localhost
```

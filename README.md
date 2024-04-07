multiline
=========

```
go get github.com/frontside/multiline 
```

### Example
```
mls := String(`First line. 
               This Line can start here.
               All leading whitespace will be removed.`)

>>> First line. 
>>> This Line can start here.
>>> All leading whitespace will be removed.`)
```


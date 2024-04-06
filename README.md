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

### Release

- update version in version.go
- Merge to master, wait for workflow to finish successfully (does not publish new version)
- Run make release locally 

# To know
Several ways to invok test functions in `go`

```go
// Run MvFile function goly
cli = fmt.Sprintf("mv %s %s", srcFile, dstFile)
_, err = runRemoteCommand(t, cli)
assert.NoError(t, err, "error while moving file remotely")
```
```go
// run the code under test
cli = fmt.Sprintf("luc do MoveFile %s %s", srcFile, dstFile)
_, err = util.RunCLILocal(cli)
```
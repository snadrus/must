# must
When Go MUST get to the point.

For those not afraid of leaving idiom for code density and DRY:

```go
func copyFile(from, to string) (err error) {
  must.RecoverToError(&err)
  f := must.One(os.Open(from))
  g := must.One(os.Create(to))
  must.E2p(io.Copy(g, f))
  return
}
```

# ggit

Writing Git in Go

# Development

Recompile cli binary when go files change.

```
find . -name '*.go' | entr go build -o main ./cli
```

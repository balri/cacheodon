# Build + run example
```
go mod tidy
go build ./...
go run ./examples/list-caches
```

# Rebuild & Test
```
go build ./pkg/geocaching
go test ./pkg/geocaching
```

# Commit Changes
```
git checkout library/extract-geocaching
git status
git add .
git commit -m "Extract geocaching code into standalone library; add example usage"
git push origin library/extract-geocaching
git tag v0.1.0
git push origin v0.1.0
```

# Optional: create a Pull Request
* Go to your fork on GitHub: https://github.com/balri/cacheodon
* You’ll see a button to Compare & pull request for your branch.
* You can create a PR if you want to merge it into main (your fork’s main), or just leave it as a feature branch.

# Use it in other projects
```
go get github.com/balri/cacheodon@v0.1.0
```
then:
```
import "github.com/balri/cacheodon/pkg/geocaching"
```

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
git tag v0.2.4
git push origin v0.2.4
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

# Search Parameters

* asc - sort ascending=true, descending=false
* att - attributes, comma-separated list of ids, positive only, only returns caches with all attributes listed
* cc - corrected coordinates, corrected=1, original=0, omit for both
* cn - cache name
* cs - cache size, comma separated list of ids
* ct - cache type, comma separated list of ids
* d - difficulty, comma-separated list
* fad - found after date
* fb - found by, appears multiple times for multiple values (up to 10)
* fbd - found before date
* fed - found end date, used with fsd
* fod - found on date
* fp - minimum favourite points
* fsd - found start date, used with fed
* hb - hidden by, single value only
* hf - hide found (by searcher), used with nfb=username, show=0, hide=1, omit for both
* ho - hide owned (by searcher) show=0, hide=1, omit for both
* lat - latitude, see ot
* lng - longitude, see ot
* m - fill grid, if true, will redirect to use values, values in form 4-4.5,3-3.5, etc
* nfb - not found by, appears multiple times for multiple values (up to 10)
* oid - see ot
* ot
  * city - oid is id of city eg Brisbane = 3356
  * region - oid is id of region eg Queensland = 54
  * country - oid is id of country eg Australia = 3
  * geocache - oid is GC code of geocache
  * coords - no oid set, lat and lng set
  * query - oid = -1
* pad - placed after date
* pbd - placed before date
* ped - placed end date, used with psd
* pn - personal notes, with notes=1, without notes=0, omit for both
* pod - placed on date
* psd - placed start date, used with ped
* r - radius (km - set per user)
* sd - show disabled, enabled=0, disabled=1, omit for both
* skip - used for pagination
* sort
  * distance
  * geocacheName
  * favoritePoint
  * containerSize
  * difficulty
  * terrain
  * trackableCount
  * foundDate
  * placeDate
* sa - show archived, show=1
* sp - show premium, basic=0, premium=1
* st - search term
* t - terrain, comma-separated list

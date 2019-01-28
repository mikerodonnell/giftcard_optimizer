find the optimal combination of items to use your full gift card balance

### quickstart
```
git clone git@github.com:mikerodonnell/giftcard_optimizer

cd giftcard_optimizer

go run cmd/main.go prices.txt 5000
```

### usage
any price list can be swapped for the sample prices.txt. ex.:
```
go run cmd/main.go /Users/admin/Documents/test.txt 8000
``` 

### tests
```
go test -v ./...
```

### notes
* algorithm sorts in O(n\*log(n)), then scans in O(n). so O(n\*log(n)) overall. 
* runs on Go 1.11 with Go Modules. no GOPATH needed!
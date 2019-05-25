for s in `ls *_test.go`
do
    go test -v ${s} ${s%_test.go}.go
done

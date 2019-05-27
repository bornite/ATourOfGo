for s in `ls *_test.go`
do
    echo "-- ${s} --"
    go test -v ${s} ${s%_test.go}.go
done

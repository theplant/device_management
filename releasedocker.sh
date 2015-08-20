GOOS=linux GOARCH=amd64 go build -o devicem main.go
cp -r /Users/sunfmin/gopkg/src/github.com/qor/qor/admin/views views 
docker build -t theplant/devicem .
rm devicem
rm -r views
docker tag -f theplant/devicem yiminbuluo.com:5000/theplant/devicem
docker push yiminbuluo.com:5000/theplant/devicem

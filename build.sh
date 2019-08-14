GOARCH=amd64 GOOS=linux go build -o vms

docker build -t vms .

docker run -d vms
echo vms server running as a docker container
echo ------------------------------------------
echo 'type docker ps to view container'

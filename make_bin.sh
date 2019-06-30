docker run --rm -it \
  -v $PWD:/app \
  -v $PWD:/app \
  golang:1.12 \
  bash -c "cd /app; go build .; chown $UID:$UID go.mod go.sum web"
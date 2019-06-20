# Start from golang v1.12 base image
FROM golang:1.12

# Add Maintainer Info
LABEL maintainer="Moritz Momper <moritz.momper@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/momper14/web

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# enable go mod
ENV GO111MODULE=on

# Install the package
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go install -v ./...

# set volumes
VOLUME ["/go/src/github.com/momper14/web/static/images"]

# This container exposes port 3001 to the outside world
EXPOSE 3001

# Run the executable
CMD ["web"]
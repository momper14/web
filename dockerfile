# Start from golang v1.12 base image
FROM golang:1.12

# Add Maintainer Info
LABEL maintainer="Moritz Momper <moritz.momper@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
ADD app.tar /app

# set volumes
VOLUME ["/app/static/images"]

# This container exposes port 3001 to the outside world
EXPOSE 3001

# Run the executable
CMD ["./web"]
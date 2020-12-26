FROM golang:1.15

# Set the Current Working Directory inside the container

ADD . /go/src/github.com/mustafa-korkmaz/goapitemplate
WORKDIR /go/src/github.com/mustafa-korkmaz/goapitemplate

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 1907 to the outside world
EXPOSE 1907

# Run the executable
CMD ["/go/bin/api"]
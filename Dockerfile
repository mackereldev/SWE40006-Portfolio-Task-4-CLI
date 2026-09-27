FROM golang:1.27
WORKDIR /usr/local/app

# Copy in the module
COPY ./ ./

# Install the module dependencies
RUN go mod download

# Compile the module
RUN go build .

ENTRYPOINT ["./convert_case"]

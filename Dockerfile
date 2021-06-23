# Specify the version of Go to use

FROM golang:1.16

# Copy all the files from the host into the container
WORKDIR /
COPY . .

# Enable Go modules
ENV GO111MODULE=on

# Compile the action
RUN unset GOPATH && go build -o /bin/push

# Specify the container's entrypoint as the action
ENTRYPOINT ["/bin/push"]
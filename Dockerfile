FROM golang:1.7

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/AdhityaRamadhanus/bigred

WORKDIR /go/src/github.com/AdhityaRamadhanus/bigred

RUN make 

# Run the outyet command by default when the container starts.
ENTRYPOINT ["./bigred"] 

# Expose ports.
EXPOSE 6399
FROM golang:1.23.2
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/wolf-gate .

FROM scratch
COPY --from=0 /bin/wolf-gate /bin/wolf-gate
CMD ["/bin/wolf-gate"]

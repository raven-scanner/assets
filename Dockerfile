FROM golang:1.14.6 as builder

WORKDIR $GOPATH/src/github.com/apuigsech/circuit/models/assets
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/assets .

FROM alpine

COPY --from=builder /bin/assets /bin/assets

ENTRYPOINT ["/bin/assets"]
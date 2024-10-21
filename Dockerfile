FROM golang
WORKDIR /work
COPY . .
RUN go build -v
ENTRYPOINT [ "./main" ]
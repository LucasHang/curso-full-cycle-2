FROM golang:latest as builder

WORKDIR /src

RUN echo '\
    package main \n\
    import "fmt" \n\
    func main() { \n\
        fmt.Println("hello world") \n\
    }\
' > hello.go

RUN go build hello.go

FROM scratch

COPY --from=builder /src/hello .

CMD [ "./hello" ]
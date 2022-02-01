FROM golang:latest as builder

WORKDIR /src

RUN echo '\
    package main \n\
    import "fmt" \n\
    func main() { \n\
        fmt.Println("Code.education Rocks!") \n\
    }\
' > hello.go

RUN go build hello.go

FROM scratch

COPY --from=builder /src/hello .

CMD [ "./hello" ]
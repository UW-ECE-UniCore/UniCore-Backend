FROM golang:1.21-alpine AS build

#COPY biz/ /go/src/biz/
#COPY pkg/ /go/src/pkg/
#COPY go.mod go.sum *.go /go/src/

WORKDIR "/app"

COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOOS=linux \
    && go env -w GOARCH=amd64
RUN go mod tidy
RUN go build -o unicore

FROM scratch
WORKDIR /app
COPY --from=build /app/unicore .
EXPOSE 8080
ENTRYPOINT ["/app/unicore"]


#FROM jrottenberg/ffmpeg:5-alpine
#
#RUN mkdir "/app"
#COPY --from=build /go/src/unicore /app/unicore
#
#RUN chmod +x /app/unicore
#
#EXPOSE 18005
#ENTRYPOINT ["/app/unicore"]
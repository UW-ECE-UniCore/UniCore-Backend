FROM golang:1.21-alpine AS build

#COPY biz/ /go/src/biz/
#COPY pkg/ /go/src/pkg/
COPY ../go.mod go.sum *.go /go/src/

WORKDIR "/go/src/"
RUN go env -w GO111MODULE=on \
    && go env -w GOOS=linux \
    && go env -w GOARCH=amd64
RUN go mod tidy
RUN go build -o unicore

FROM alpine:latest
WORKDIR /app
COPY --from=build /go/src/unicore .
EXPOSE 18005
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
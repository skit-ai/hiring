FROM golang:latest as builder
# Add Maintainer Info
LABEL maintainer="pankaj dubey <dubeyp36@gmail.com>"

## We"ll create an /app directory in which,we"ll put all of our project code
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...

## run our application, using alpine to reduce size
FROM alpine:latest AS production
## We have to copy the output from our, builder stage to our production stage
COPY --from=builder /app .
## our newly compiled
## binary exectuable!!
CMD ["./main"]

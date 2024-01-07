# Build Stage (we build here -it is a heavy container)
FROM golang:1.21-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage (then we copy here - it is a very light container)
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .

# EXPOSE 2223 # wont have any service. This will only consume to get templates and receivers from db. Wont be exposed
# has become only an additional parameter add to entry point. Entry point will executed after configurations are done and start the application 
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh"]
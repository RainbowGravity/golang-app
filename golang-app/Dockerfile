FROM golang:1.17-alpine AS builder

# First layer of image building
LABEL stage=builder

WORKDIR /app

COPY *.* /app/

RUN CGO_ENABLED=0 go build main.go

RUN ls -al

# Second layer of image building

FROM scratch 

USER 1001

WORKDIR /app

ENV PATH="/app:${PATH}"

EXPOSE 8080

COPY --from=builder --chown=1001:1001 /app/main /app/main

# Run app
ENTRYPOINT [ "/app/main" ]
# builder
FROM golang:1.15.2 AS builder
ARG GITHUB_TOKEN
WORKDIR /web_client
COPY . .
RUN echo "prod" > .version
RUN echo "machine github.com\n\tlogin elojah\n\tpassword ${GITHUB_TOKEN}" >> ~/.netrc
ENV GOPRIVATE=github.com/elojah/*
# RUN git config --global url."ssh://git@github.com/".insteadOf https://github.com/
RUN make tidy
RUN env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 make web_client

# runner
FROM alpine
WORKDIR /app
COPY --from=builder /web_client/bin/game_03_web_client .
COPY --from=builder /web_client/config/web_client/prod.json .
RUN chmod +x /app/game_03_web_client
ENTRYPOINT ["/app/game_03_web_client", "/app/prod.json"]

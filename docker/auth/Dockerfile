# builder
FROM golang:1.15.2 AS builder
ARG GITHUB_TOKEN
WORKDIR /auth
COPY . .
RUN echo "prod" > .version
RUN echo "machine github.com\n\tlogin elojah\n\tpassword ${GITHUB_TOKEN}" >> ~/.netrc
ENV GOPRIVATE=github.com/elojah/*
# RUN git config --global url."ssh://git@github.com/".insteadOf https://github.com/
RUN make tidy
RUN env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 make auth

# runner
FROM alpine
WORKDIR /app
COPY --from=builder /auth/bin/game_03_auth .
COPY --from=builder /auth/config/auth/prod.json .
RUN chmod +x /app/game_03_auth
ENTRYPOINT ["/app/game_03_auth", "/app/prod.json"]

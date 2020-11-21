# builder
FROM golang:1.15.2 AS builder
ARG GITHUB_TOKEN
WORKDIR /api
COPY . .
RUN echo "prod" > .version
RUN echo "machine github.com\n\tlogin elojah\n\tpassword ${GITHUB_TOKEN}" >> ~/.netrc
ENV GOPRIVATE=github.com/elojah/*
# RUN git config --global url."ssh://git@github.com/".insteadOf https://github.com/
RUN make tidy
RUN env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 make api

# runner
FROM alpine
WORKDIR /app
COPY --from=builder /api/bin/game_03_api .
COPY --from=builder /api/config/api/prod.json .
COPY --from=builder /api/config/api/firebase.json .
RUN chmod +x /app/game_03_api
ENTRYPOINT ["/app/game_03_api", "/app/prod.json"]

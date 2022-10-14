# builder
FROM golang:1.15.2 AS builder
ARG GITHUB_TOKEN
WORKDIR /web_dashboard
COPY . .
RUN echo "prod" > .version
RUN echo "machine github.com\n\tlogin elojah\n\tpassword ${GITHUB_TOKEN}" >> ~/.netrc
ENV GOPRIVATE=github.com/elojah/*
# RUN git config --global url."ssh://git@github.com/".insteadOf https://github.com/
RUN make tidy
RUN env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 make web_dashboard

# runner
FROM alpine
WORKDIR /app
COPY --from=builder /web_dashboard/bin/game_03_web_dashboard .
COPY --from=builder /web_dashboard/config/web_dashboard/prod.json .
RUN chmod +x /app/game_03_web_dashboard
ENTRYPOINT ["/app/game_03_web_dashboard", "/app/prod.json"]

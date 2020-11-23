# builder
FROM golang:1.15.2 AS builder
ARG GITHUB_TOKEN
WORKDIR /admin
COPY . .
RUN echo "prod" > .version
RUN echo "machine github.com\n\tlogin elojah\n\tpassword ${GITHUB_TOKEN}" >> ~/.netrc
ENV GOPRIVATE=github.com/elojah/*
# RUN git config --global url."ssh://git@github.com/".insteadOf https://github.com/
RUN make tidy
RUN env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 make admin

# runner
FROM alpine
WORKDIR /app
COPY --from=builder /admin/bin/game_03_admin .
COPY --from=builder /admin/config/admin/prod.json .
COPY --from=builder /admin/config/admin/firebase.json .
RUN chmod +x /app/game_03_admin
ENTRYPOINT ["/app/game_03_admin", "/app/prod.json"]

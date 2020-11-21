# For CI ONLY !!!
# Check with internal import allowed
FROM golang:1.15.2
ARG GITHUB_TOKEN
WORKDIR /api
COPY . .
ENV GOPRIVATE=github.com/elojah/*
# RUN git config --global url."ssh://git@github.com/".insteadOf https://github.com/
# Install golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.31.0
RUN echo "machine github.com\n\tlogin elojah\n\tpassword ${GITHUB_TOKEN}" >> ~/.netrc
RUN make check

# FROM fuzzers/go-fuzz:1.2.0
FROM golang:1.18 as builder

COPY . /elemental
WORKDIR /elemental/tests/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest

RUN apt update && apt install -y clang

RUN go mod tidy
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep

RUN go-fuzz-build -libfuzzer -o fuzz.a && \
    clang -fsanitize=fuzzer fuzz.a -o fuzz.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /elemental/tests/mayhem/fuzz.libfuzzer /
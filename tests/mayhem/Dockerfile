FROM golang:1.18 as builder

COPY . /elemental
WORKDIR /elemental/tests/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest

RUN apt update && apt install -y clang

RUN go mod tidy
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN go get github.com/AdaLogics/go-fuzz-headers

RUN go-fuzz-build -libfuzzer -o fuzzMisc.a && \
    clang -fsanitize=fuzzer fuzzMisc.a -o fuzzMisc.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /elemental/tests/mayhem/fuzzMisc.libfuzzer /
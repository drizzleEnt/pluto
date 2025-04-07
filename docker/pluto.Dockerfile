FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o plutod ./cmd/plutod

FROM ubuntu:22.04
COPY --from=builder /app/plutod /usr/bin/plutod
RUN apt-get update && apt-get install -y curl jq
COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh
VOLUME /root/.pluto-chain
EXPOSE 26656 26657 1317 9090
ENTRYPOINT ["entrypoint.sh"]
CMD ["start", "--minimum-gas-prices", "0.025stake", "--home", "/root/.pluto-chain"]
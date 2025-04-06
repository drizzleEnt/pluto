FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o plutod ./cmd/plutod

FROM ubuntu:22.04
COPY --from=builder /app/plutod /usr/bin/plutod
# Установка зависимостей для работы plutod
RUN apt-get update && apt-get install -y curl
# Инициализация цепочки
RUN plutod init pluto-node --chain-id pluto-test
RUN plutod keys add validator --keyring-backend test
RUN VALIDATOR_ADDR=$(plutod keys show validator -a --keyring-backend test) && \
    plutod genesis add-genesis-account $VALIDATOR_ADDR 1000000000stake --keyring-backend test && \
    plutod genesis gentx validator 1000000stake --chain-id pluto-test --keyring-backend test && \
    plutod genesis collect-gentxs
# Копирование конфигурации в том (опционально)
VOLUME /root/.pluto-chain
EXPOSE 26656 26657 1317 9090
ENTRYPOINT ["plutod"]
CMD ["start", "--minimum-gas-prices", "0.025stake"]
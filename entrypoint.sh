#!/bin/bash

# Проверяем, существует ли genesis.json
if [ ! -f /root/.pluto-chain/config/genesis.json ]; then
  echo "Initializing Pluto node..."
  plutod init pluto-node --chain-id pluto-test --home /root/.pluto-chain
  
  echo "Creating validator key..."
  echo "y" | plutod keys add validator --keyring-backend test --home /root/.pluto-chain
  
  VALIDATOR_ADDR=$(plutod keys show validator -a --keyring-backend test --home /root/.pluto-chain)
  echo "Validator address: $VALIDATOR_ADDR"
  
  echo "Adding genesis account..."
  plutod genesis add-genesis-account "$VALIDATOR_ADDR" 1000000000stake --keyring-backend test --home /root/.pluto-chain
  
  echo "Generating gentx..."
  plutod genesis gentx validator 1000000000stake --chain-id pluto-test --keyring-backend test --home /root/.pluto-chain
  
  echo "Collecting gentxs..."
  plutod genesis collect-gentxs --home /root/.pluto-chain
fi

# Запускаем ноду
exec plutod "$@"
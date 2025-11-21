#!/usr/bin/env bash

set -e  # Para o script se algo der errado

echo "[1/5] Criando pasta build..."
mkdir -p build

echo "[2/5] Entrando na pasta build..."
cd build

echo "[3/5] Rodando CMake..."
cmake ..

echo "[4/5] Rodando Make..."
make

echo "[5/5] Iniciando servidor em background..."
./servidor &

SERVER_PID=$!
echo "Servidor iniciado com PID $SERVER_PID"

sleep 1  
echo "Iniciando cliente..."
./cliente

echo "Cliente finalizado."

echo "Encerrando servidor..."
kill $SERVER_PID 2>/dev/null || true

echo "Pronto!"

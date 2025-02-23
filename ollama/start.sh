#!/bin/sh

echo "Starting Ollama..."
ollama serve &

echo "Waiting for the server to start..."
sleep 10

echo "Creating the model..."
ollama create deepseek -f /Modelfile

echo "Loading the model..."
curl -s -o /dev/null -w "%{http_code}" -X POST http://localhost:11434/api/generate \
  -H "Content-Type: application/json" \
  -d '{"model": "deepseek", "keep-alive": -1}'

echo "Model loaded accepting requests..."
tail -f /dev/null
#!/bin/bash

# Aguarda o LocalStack subir completamente
echo "Aguardando LocalStack iniciar..."
until curl -s http://localhost:4566/_localstack/health | grep '"s3": "running"' > /dev/null; do
  sleep 2
  echo "Esperando S3..."
done

echo "Criando bucket S3: reprocessing_engine"
aws --endpoint-url=http://localhost:4566 \
    s3api create-bucket \
    --bucket reprocessing_engine \
    --region us-east-1

echo "Bucket criado com sucesso."

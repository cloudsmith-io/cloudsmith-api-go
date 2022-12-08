FROM openapitools/openapi-generator-cli:v6.2.1

RUN apt-get update \
 && apt-get install -y --no-install-recommends jq golang \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /local

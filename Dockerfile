FROM openapitools/openapi-generator-cli:v5.4.0

RUN apt-get update \
 && apt-get install -y --no-install-recommends jq=1.5+dfsg-2+b1 golang=2:1.11~1 \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /local

#!/bin/bash

set -e

docker run -d \
      --name nginx-api \
      -p 80:80 -p 443:443 \
      -v /etc/letsencrypt:/etc/letsencrypt:ro \
      -v ./api.conf:/etc/nginx/nginx.conf:ro \
      -v /var/www/webroot:/var/www/certbot:ro \
      nginx:stable-alpine
#!/usr/bin/env bash
openssl req -new -newkey rsa:4096 -nodes -keyout private_key.txt -out csr.txt -subj "/C=FR/ST=ARL/L=Annecy/O=LegacyFactory/CN=www.legacyfactory.dev"
echo 'subjectAltName=DNS:legacyfactory.dev,DNS:www.legacyfactory.dev' >alt_names.txt
openssl x509 -req -extfile alt_names.txt -sha256 -days 365 -in csr.crt -signkey private.key -out public.crt

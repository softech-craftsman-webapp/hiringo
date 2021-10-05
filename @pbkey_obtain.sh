mkdir -p keys

# Get the public key of Auth Service
curl http://localhost:8080/public-key -o keys/public.pem

# trading-automation-system

# BUILD docker container
docker build -t golangapp .

# RUN docker container
docker run --name golangapp -p 9000:9000 golangapp:latest

# RUN docker compose
docker-compose up -d

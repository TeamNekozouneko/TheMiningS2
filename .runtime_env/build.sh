docker image prune -f
docker container prune -f
docker build -t teamnekozouneko/theminings2:latest -f .runtime_env/Dockerfile .
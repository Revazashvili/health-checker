# health-checker

background service to check service health status


# Build
```
docker build -t health-checker .
```

# Run
```
docker run -e DURATION_IN_SECONDS=5 -e HEALTH_CHECK_URL=https://github.com health-checker
```

# Network Endpoints Checker
For those cases when you need to check the availability of ports through a web form.

## Build image with Dockerfile.

```
cd *project_root_folder*
docker build -t networkchecker .
```

## Run container

```
docker run -d networkchecker -p 8080:8080
```

## Go to application web form.

```
http://localhost:8080
```

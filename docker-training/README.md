- Use base image


- to build docker image

```docker build . -t jpalaparthi/sample:v0.0.1```

- if the file name is not Dockerfile then

```docker build . -f dockerfile_name -t jpalaparthi/sample:v0.0.1```

- to push docker image to the repo

```docker push jpalaparthi/sample:v0.0.1```
FROM ubuntu

LABEL "author" "jiten"

LABEL "purpose" "demo"

ARG name="sample"

ENV DIRNAME=/app
#if the directory is not available then it creates a new directory
WORKDIR ${DIRNAME}

# to add a file from source to destination
ADD README.md .

#ADD https://golang.org/dl/go1.20.2.linux-amd64.tar.gz/ .

COPY misc/ misc/

RUN apt update && \
    apt install vim -y


CMD [ "bash" ]

# apt update 
# apt install vim -y
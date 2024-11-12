FROM ubuntu:latest

RUN apt-get update && apt-get install -y python3 python3-pip

COPY . /app
WORKDIR /app

RUN pip3 install -r requirements.txt

ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=mydatabase
ENV MYSQL_USER=myuser
ENV MYSQL_PASSWORD=mypassword

# Add any additional configuration for your db service here

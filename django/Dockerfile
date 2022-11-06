FROM python:3.10 AS build
WORKDIR /python/src
COPY requirements.txt .
RUN pip3 install -r requirements.txt
COPY src .

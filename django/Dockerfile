FROM python:3.10 AS build
WORKDIR /python/src
COPY requirements.txt .
RUN pip3 install -r requirements.txt
COPY src .
EXPOSE 8000/tcp
CMD ["gunicorn", "-b 0.0.0.0:8000", "-w 5", "src.app.core.wsgi:application"]

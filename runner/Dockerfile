FROM python:3
RUN apt-get update


WORKDIR /go/src
COPY runner.py .

CMD [ "python", "runner.py" ]
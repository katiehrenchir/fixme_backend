FROM python:2.7
WORKDIR /code
ADD  requirements.txt .
RUN pip install -r requirements.txt
ADD . .
USER 1000

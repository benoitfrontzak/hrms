FROM alpine:latest

RUN mkdir /app

COPY infoApp /app

CMD [ "/app/infoApp"]
FROM alpine:latest

RUN mkdir /app

COPY amqpApp /app

CMD [ "/app/amqpApp"]
FROM alpine:latest

RUN mkdir /app

COPY leaveApp /app

CMD [ "/app/leaveApp"]
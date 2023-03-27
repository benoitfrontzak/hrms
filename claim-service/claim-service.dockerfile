FROM alpine:latest

RUN mkdir /app

COPY claimApp /app

CMD [ "/app/claimApp"]
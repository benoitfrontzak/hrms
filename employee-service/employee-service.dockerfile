FROM alpine:latest

RUN mkdir /app

COPY employeeApp /app

CMD [ "/app/employeeApp"]
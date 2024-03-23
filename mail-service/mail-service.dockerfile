FROM alpine:latest

RUN mkdir /app

COPY mailServiceApp /app

CMD [ "/app/mailServiceApp" ]
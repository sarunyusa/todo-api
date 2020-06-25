FROM alpine:3.10.1

# Install prerequisites
RUN apk add --no-cache curl tzdata

WORKDIR /app/svc
COPY ./dist/todo-linux-x64 .

CMD ./todo-linux-x64
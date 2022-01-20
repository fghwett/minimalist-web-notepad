FROM alpine:latest
MAINTAINER fghwett fghwett@qq.com
WORKDIR /app
COPY ./static /app/static
COPY ./notepad /app/notepad
VOLUME ["/app/_tmp"]
EXPOSE 9099
CMD ["/app/notepad"]
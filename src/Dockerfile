FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
ADD main /
ADD config/properties.yaml config/
CMD ["/main"]
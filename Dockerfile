FROM golang:1.11.0

ARG SERVICE

ENV SERVICE "$SERVICE"
RUN echo $SERVICE

RUN export PATH=$PATH:$(which go)

WORKDIR /go/src/github.com/chrispaynes/gRPCrud
COPY . .
RUN ls -lahg

ENTRYPOINT [ "sh", "-c", "sh ./scripts/$SERVICE" ]

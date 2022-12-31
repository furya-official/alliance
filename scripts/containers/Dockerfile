FROM golang:1.18.7-bullseye AS build

RUN apt update && apt install jq build-essential -y

WORKDIR /kaijud
COPY . /kaijud

RUN curl https://get.ignite.com/cli@v0.25.1! | bash
RUN ignite chain build --output . --verbose

FROM debian:bullseye AS run
COPY ./scripts/containers/build-wrapper.sh /kaijud/build-wrapper.sh

VOLUME /kaijud
COPY --from=build /kaijud/ /kaijud/
WORKDIR /kaijud

EXPOSE 26656 26657
ENTRYPOINT ["/kaijud/build-wrapper.sh"]
CMD ["start", "--log_format", "plain"]
STOPSIGNAL SIGTERM

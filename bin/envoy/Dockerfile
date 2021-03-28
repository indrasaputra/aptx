FROM envoyproxy/envoy-alpine:v1.17.1
COPY bin/envoy/envoy.yaml /etc/envoy/envoy.yaml
COPY bin/image.bin /etc/envoy/image.bin
RUN chmod go+r /etc/envoy/envoy.yaml
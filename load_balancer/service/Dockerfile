FROM ubuntu:20.04

LABEL maintainer="yuantingwei@pm.me"

# Adjust Timezone
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update
RUN apt-get install ca-certificates -y
RUN update-ca-certificates


ADD big_num_compute_service /bin/
ENTRYPOINT ["/bin/big_num_compute_service"]


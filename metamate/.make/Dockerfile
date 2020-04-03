FROM bitnami/minideb:latest

RUN install_packages ca-certificates

COPY dist/metamate metamate

CMD ./metamate serve
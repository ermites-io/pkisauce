FROM scratch
COPY pkisauce bin/pkisauce
LABEL org.opencontainers.image.description pkisauce
ENTRYPOINT ["bin/pkisauce"]

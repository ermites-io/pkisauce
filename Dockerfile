FROM scratch
COPY pkisauce bin/pkisauce
ENTRYPOINT ["bin/pkisauce"]

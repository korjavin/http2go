FROM korjavin/korjavin-base
RUN apt-get update  
RUN apt-get install -y ca-certificates
ADD http2tg /http2tg
CMD "/http2tg"

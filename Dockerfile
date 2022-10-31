FROM golang:1.18

RUN mkdir /api
RUN mkdir solrdata
RUN -D -V "$PWD/solrdata:/var/solr" -p 8983:8983 --name my_solr solr solr-precreate items
ADD . /api
WORKDIR /api

RUN go mod tidy
RUN go build -o api .
RUN chmod +x /api/api

ENTRYPOINT ["/api/api"]




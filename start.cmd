docker run -d -p 8983:8983 --name my_solr solr
docker exec -it my_solr solr create_core -c publicaciones

go run main.go


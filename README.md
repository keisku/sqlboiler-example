# Usage

### start app

```shell
docker-compose build
docker-compose up -d
sqlboiler psql
docker-compose logs --tail=10 -f sqlboiler_example_app
```

see list of articles
http://localhost:3000/articles

create new article
http://localhost:3000/article/new

### close app

```shell
docker-compose down
```

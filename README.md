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

![](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/275587/e7a9be0d-b438-a314-0dda-a058953906a0.png)

create new article
http://localhost:3000/article/new

![](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/275587/d1a1ea10-0149-e90f-a38e-2d74197602bf.png)

### close app

```shell
docker-compose down
```

# Dcokerfile for go-gin-wrapper

FROM cassandra:3.7

#ARG redisHostName=default-redis-server
#ARG mysqlHostName=default-mysql-server

RUN mkdir -p /hy

#ENV REDIS_URL=redis://h:password@${redisHostName}:6379
#ENV CLEARDB_DATABASE_URL=mysql://hiromaily:12345678@mysql-server/hiromaily?reconnect=true

COPY ./init/init.sh /hy
COPY ./init/init.sql /hy
COPY ./docker-entrypoint.sh /
#VOLUME ./docker_build/cassandra:/hy

WORKDIR /hy

#RUN ./init.sh
#RUN cqlsh --file /hy/init.sql


EXPOSE 9042

#ENTRYPOINT ["ginserver", "-f", "./configs/docker.toml"]
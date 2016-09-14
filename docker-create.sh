#!/bin/sh

###############################################################################
# For golibs
###############################################################################

###############################################################################
# Environment
###############################################################################
CONTAINER_MYSQL=lib-mysql
CONTAINER_REDIS=lib-redis
CONTAINER_MONGO=lib-mongo
CONTAINER_CASSANDRA=lib-cassandra
CONTAINER_NATS=lib-nats1
CONTAINER_KAFKA=lib-kafka1
CONTAINER_ZOOKEEPER=lib-zookeeper1
CONTAINER_RMQ=lib-rmq1


ARY=()
ARY+=(${CONTAINER_MYSQL})
ARY+=(${CONTAINER_REDIS})
ARY+=(${CONTAINER_MONGO})
ARY+=(${CONTAINER_CASSANDRA})
ARY+=(${CONTAINER_NATS})
ARY+=(${CONTAINER_KAFKA})
ARY+=(${CONTAINER_ZOOKEEPER})
ARY+=(${CONTAINER_RMQ})


###############################################################################
# Remove Container And Image
###############################################################################
for i in "${ARY[@]}"
do
    DOCKER_PSID=`docker ps -af name="${i}" -q`
    if [ ${#DOCKER_PSID} -ne 0 ]; then
        docker rm -f ${i}
    fi
done

#docker rm -f $(docker ps -aq)



###############################################################################
# Docker-compose / build and up
###############################################################################
docker-compose  build
docker-compose  up -d

# MONGO settings
sleep 3s
MONGO_PORT=30017
mongo 127.0.0.1:${MONGO_PORT}/admin --eval "var port = ${MONGO_PORT};" ./docker_build/mongo/init.js
mongorestore -h 127.0.0.1:${MONGO_PORT} --db hiromaily docker_build/mongo/dump/hiromaily

# Cassandra
# TODO: Why it fails at this time??
#docker exec -it lib-cassandra bash /hy/init.sh

# RabbitMQ
#testQueue

###############################################################################
# Docker-compose / check
###############################################################################
docker-compose ps
docker-compose logs


###############################################################################
# Test
###############################################################################



###############################################################################
# Docker-compose / down
###############################################################################
#docker-compose -f down

###############################################################################
# Check connection
###############################################################################
#mysql -u root -p -h 127.0.0.1 -P 13306
#redis-cli -h 127.0.0.1 -p 16379 -a password

#Access by browser
#http://docker.hiromaily.com:9999/

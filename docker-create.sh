#!/bin/bash

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
#wait to be ready
echo '[cassandra] starting cassandra now.'
sleep 10s
while :
do
    LOGS=`docker logs lib-cassandra --tail 5 | grep "Starting listening for CQL clients on /0.0.0.0:9042"`
    echo $LOGS
    if [ ${#LOGS} -ne 0 ]; then
        docker exec -it lib-cassandra bash /hy/init.sh
        break
    else
        echo 'running...'
        sleep 1s
    fi
done
echo '[cassandra] done!'

#kafka
echo '[kafka] starting kafka now.'
#while :
#do
#    LOGS=`docker logs lib-kafka1 --tail 5 | grep "Kafka Server 1001], started"`
#    echo $LOGS
#    if [ ${#LOGS} -ne 0 ]; then
#        break
#    else
#        echo 'running...'
#        sleep 1s
#    fi
#done

#docker exec -it lib-kafka1 bash -c 'echo $KAFKA_HOME'
docker exec -it lib-kafka1 bash -c '$KAFKA_HOME/bin/kafka-topics.sh --create --topic NewTopic100 --partitions 1 --zookeeper zookeeper:2181 --replication-factor 1'
docker exec -it lib-kafka1 bash -c '$KAFKA_HOME/bin/kafka-topics.sh --list --zookeeper zookeeper:2181'

echo '[kafka] done!'


#RabbitMQ create queue


###############################################################################
# Docker-compose / check
###############################################################################
docker-compose ps
#docker-compose logs


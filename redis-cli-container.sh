#!/bin/bash

# docker inspect a79a5ec8a16e | grep "IPAddress" | grep -o '[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}'

ipAddress=$(docker inspect testing_redis_db | grep "IPAddress" | grep -o '[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}')

echo $ipAddress

redis-cli  -h $ipAddress

# Where 
#    172.18.0.2 is the IP-address of Postgres container,
#      can be obtained through "docker inspect a79a5ec8a16e | grep IP"
#      where a79a5ec8a16e is the hash (ID) of the PG container
#    dbUser is the username mentioned in the stack.yml in enviromnent section


#!/bin/bash


psql  -h 172.18.0.2 -p 5432 -U dbUser -W -d InfoTestDb

# Where 
#    172.18.0.2 is the IP-address of Postgres container,
#      can be obtained through "docker inspect a79a5ec8a16e | grep IP"
#      where a79a5ec8a16e is the hash (ID) of the PG container
#    dbUser is the username mentioned in the stack.yml in enviromnent section

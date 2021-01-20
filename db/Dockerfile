FROM postgres:13

#COPY ssl/server.key /var/lib/postgresql/server.key
#COPY ssl/server.crt /var/lib/postgresql/server.crt
#RUN chown postgres /var/lib/postgresql/server.key && \
#    chmod 600 /var/lib/postgresql/server.key

RUN localedef -i ru_RU -c -f UTF-8 -A /usr/share/locale/locale.alias ru_RU.UTF-8
ENV LANG ru_RU.utf8

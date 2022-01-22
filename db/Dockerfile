FROM mariadb:10.7

ENV MARIADB_ROOT_PASSWORD=1234

ADD DB.sql /docker-entrypoint-initdb.d
# ADD my.cnf /etc/mysql/conf.d

EXPOSE 3306
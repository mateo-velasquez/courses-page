FROM mysql:8

ADD createTables.sql /docker-entrypoint-initdb.d
ADD loadData.sql /docker-entrypoint-initdb.d

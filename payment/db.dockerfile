FROM postgres:10.3

COPY new.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]

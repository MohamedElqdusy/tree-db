FROM postgres:11.1
## copy sql file into the container 
COPY tree.sql /docker-entrypoint-initdb.d/1.sql

## NOTE: SQL files will be executed in alphabetical order
CMD ["postgres"]
FROM liquibase/liquibase
LABEL authors="marceloprates"

COPY changelog.yaml changelog.yaml
COPY migrations/*.sql migrations/
COPY Makefile Makefile
COPY liquibase.properties liquibase.properties

CMD liquibase update --url"jdbc:postgresql://${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME} --username=${DATABASE_USER} --password=${DATABASE_PASSWD}"
FROM liquibase/liquibase
LABEL authors="marceloprates"

COPY changelog.yaml changelog.yaml
COPY migrations/*.sql migrations/
COPY ../Makefile Makefile
COPY liquibase.properties liquibase.properties

CMD liquibase update --url"${LIQUIBASE_DATABASE_URL} --username=${LIQUIBASE_DATABASE_USERNAME} --password=${LIQUIBASE_DATABASE_PASSWORD}"
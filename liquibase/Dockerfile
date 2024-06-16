FROM liquibase/liquibase
LABEL authors="marceloprates"

COPY changelog.yaml changelog.yaml
COPY migrations/*.sql migrations/
COPY liquibase.properties liquibase.properties

CMD liquibase update  --changelog-file=changelog.yaml  --url="${LIQUIBASE_COMMAND_URL}?user=${LIQUIBASE_COMMAND_USERNAME}&password=${LIQUIBASE_COMMAND_PASSWORD}&sslmode=disable"
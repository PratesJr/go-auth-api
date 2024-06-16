NOW := $(shell date '+%Y%m%d%H%M%S')

pre-build:
	echo 'starting build'

liquibase-diff:
	liquibase --url "jdbc:postgresql://${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}" --username=${DATABASE_USER} --password=${DATABASE_PASSWD} --referenceUrl "jdbc:postgresql://${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}" --referenceUsername=${DATABASE_USER} --referencePassword=${DATABASE_PASSWD} diff-changelog


liquibase-migrate:
	liquibase update --url "jdbc:postgresql://${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}" --username=${DATABASE_USER} --password=${DATABASE_PASSWD}

liquibase-rollback:
	liquibase rollback-count --count=1 --url "jdbc:postgresql://${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}" --username=${DATABASE_USER} --password=${DATABASE_PASSWD}

liquibase-update-sql:
	liquibase --url "jdbc:postgresql://${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}" --username=${DATABASE_USER} --password=${DATABASE_PASSWD} --changelog-file=./liquibase/changelog.yaml --output-file=./liquibase/script.sql --labels=setup update-sql

liquibase-create-file:
	@touch ./liquibase/migrations/$(NOW)_$(name).sql

liquibase-update:
	liquibase update --url "jdbc:postgresql://${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}" --username=${DATABASE_USER} --password=${DATABASE_PASSWD}
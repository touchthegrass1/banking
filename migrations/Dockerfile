FROM liquibase/liquibase
WORKDIR /liquibase
COPY . .
CMD ["liquibase","--changelog-file=db.changelog-master.xml","--search-path="/liquibase/db/changelog/dev/","--url jdbc:postgresql://database:5432/banking_test_db","--username banking_db_user","--password example_password","update"]


#!/usr/bin/env bash

MYSQL_USER=${MYSQL_USER:-goldie}
MYSQL_DB=${MYSQL_DB:-address_book_go}

query() { 
	mysql -u "${MYSQL_USER}" -p"${MYSQL_PASS}" -e "$1;" "$MYSQL_DB"
}

query_match() {
   local sql
   sql=$1
   shift

   for pattern in "$@"; do
      query "$sql" | grep -q "$pattern" || {
         echo "did not find: '$pattern' in '$sql'"
         return 1
      }
      echo "found: '$pattern' in '$sql'"
   done
}

finally() {
   if query_match 'show tables' 'person' 'address' &> /dev/null; then
      query 'show tables'
      query 'select * from person'
      query 'select * from address'
      exit
   fi

   exit 1
}

if query_match 'show tables' 'person' 'address'; then
   if [[ $1 != 'reset' ]]; then
      finally
   fi

   echo 'dropping and recreating tables'
   query 'drop table person; drop table address'
fi

query '
   CREATE TABLE person(
     id    MEDIUMINT    NOT NULL AUTO_INCREMENT,
     first varchar(255) NOT NULL,
     last  varchar(255) NOT NULL,

     PRIMARY KEY (id)
   );

   CREATE TABLE address (
     id        MEDIUMINT    NOT NULL AUTO_INCREMENT,
     person_id MEDIUMINT    NOT NULL,
     street    varchar(255) NOT NULL,
     city      varchar(255) NOT NULL,
     state     varchar(255) NOT NULL,
     zip       varchar(16)  NOT NULL,

     PRIMARY KEY (id)
   );
'

query '
INSERT INTO person (first, last) VALUES
   ("Tuna", "Lowland"),
   ("Shöwee", "Clocks"),
   ("Quark", "Twin")
'

query '
INSERT INTO address (person_id, street, city, state, zip) VALUES
   (1, "133 49th Ave", "Beverly Hills", "CA", "90210"),
   (2, "52 Larpenteur St", "San Franciso", "CA", "94016"),
   (1, "155 Travelington Lt", "NYC", "NY", "10001"),
   (3, "7 Rabbit Ln.", "Seattle", "WA", "98101"),
   (3, "450 Livingston Pkwy.", "Pittsburg", "PA", "94565"),
   (3, "2243 Juicy Lucy Ave.", "Minneapolis", "MN", "55407")
'

finally

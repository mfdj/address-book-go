#!/usr/bin/env bash

MYSQL_USER=${MYSQL_USER:-goldie}
MYSQL_DB=${MYSQL_DB:-address_book_go}

query() { 
	mysql -u "${MYSQL_USER}" -p"${MYSQL_PASS}" -e "$1;" "$MYSQL_DB" 2> /dev/null
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
      exit
   fi

   exit 1
}

if query_match 'show tables' 'person' 'address'; then
   if [[ $1 != clean ]]; then
      finally
   fi

   echo 'dropping and recreating tables'
   query 'drop table person; drop table address'
fi

query '
   CREATE TABLE person(
     id    MEDIUMINT    NOT NULL AUTO_INCREMENT,
     first   varchar(255) NOT NULL,
     last    varchar(255) NOT NULL,

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

finally

#!/usr/bin/env bash

#perfome the data load against client api rest interface
curl --location --request POST 'http://localhost:8080/v1/clientapi/startload' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'filename=ports-2.json' # add the file name, the small ones or the big ones. =)

#perfome full text search (will return everything close to "al") against client api rest interface
# depend of jq command, if you don't have it you can just remove
curl --location --request GET 'http://localhost:8080/v1/clientapi/search/al' | jq

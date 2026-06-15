#!/bin/bash

# echo '{ "dashboard": ' > tmp.json
# cat Velonetics_OTEL_Dashboard.json >> tmp.json
# echo '}' >> tmp.json

curl -X POST --insecure --header "Content-Type: application/json" \
    http://velonetics:velonetics@localhost:53000/api/dashboards/db \
    -d @tmp.json 

# rm tmp.json

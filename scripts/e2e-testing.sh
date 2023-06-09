#!/bin/bash
# Declare variables.
API_URL=http://localhost:8080

# Introduction to the script.
echo "Welcome to 'bill-server' application!"
echo "Before running the end-to-end tests, please ensure that you have run 'make start'!"; echo

# Testing '/api/v1'.
echo
echo "Running end-to-end testing..."
echo "Testing GET route '/api/v1'..."
curl $API_URL/api/v1; echo

# Testing '/api/v1/labels'.
echo
echo "Testing GET route '/api/v1/labels'..."
curl $API_URL/api/v1/labels; echo
echo
echo "Testing POST route '/api/v1/labels'..."
curl -X POST -H 'Content-Type: application/json' -d '{"name":"Lucy Heartfilia","type":"income"}' $API_URL/api/v1/labels; echo

# Testing '/api/v1/labels/:labelID'.
echo
echo "Using 'labelID' with value of 3 (the one created beforehand)."
echo "Testing GET route '/api/v1/labels/:labelID'..."
curl $API_URL/api/v1/labels/3; echo
echo
echo "Testing PUT route '/api/v1/labels/:labelID'..."
curl -X PUT -H 'Content-Type: application/json' -d '{"name":"Mirajane Strauss","type":"consume"}' $API_URL/api/v1/labels/3; echo
echo
echo "Testing GET route after PUT '/api/v1/labels/:labelID'..."
curl $API_URL/api/v1/labels/3; echo
echo
echo "Testing DELETE route '/api/v1/labels/:labelID'..."
curl -X DELETE $API_URL/api/v1/labels/3; echo
echo
echo "Testing GET route after DELETE '/api/v1/labels/:labelID'..."
curl $API_URL/api/v1/labels/3; echo

# Test '/api/v1/bills'
echo
echo "Using 'year' with value of 2023, 'month' with value of 5, and 'day' with value of 30 "
echo "Testing GET route '/api/v1/bills'..."
curl $API_URL/api/v1/bills; echo
echo
echo "Testing GET route '/api/v1/bills/:year'..."
curl $API_URL/api/v1/bills/2023; echo
echo
echo "Testing GET route '/api/v1/bills/:year/:month'..."
curl $API_URL/api/v1/bills/2023/5; echo
echo
echo "Testing GET route '/api/v1/bills/:year/:month/:day'..."
curl $API_URL/api/v1/bills/2023/5/30; echo
echo
echo "Testing POST route '/api/v1/bills"
curl -X POST -H 'Content-Type: application/json' -d '{"year": 2023,"month":6, "day": 10, "money": 200, "label": "shopping"}' $API_URL/api/v1/bills; echo



## Testing '/api/v1/auth/login'.
#echo
#echo "Testing POST route '/api/v1/auth/login'..."
#curl -X POST -H 'Content-Type: application/json' -d '{"username":"fiber","password":"fiber"}' -c cookie.txt $API_URL/api/v1/auth/login; echo
#
## Testing '/api/v1/auth/private'.
#echo
#echo "Testing GET route '/api/v1/auth/private'..."
#curl -b cookie.txt $API_URL/api/v1/auth/private; echo
#
## Testing '/api/v1/cities'.
#echo
#echo "Testing GET route '/api/v1/cities'..."
#curl -b cookie.txt $API_URL/api/v1/cities; echo
#echo
#echo "Testing POST route '/api/v1/cities'..."
#curl -b cookie.txt -X POST -H 'Content-Type: application/json' -d '{"name":"Kyoto"}' $API_URL/api/v1/cities; echo
#
## Testing '/api/v1/cities/:cityID'.
#echo
#echo "Using 'cityID' with value of 6 (the one created beforehand)."
#echo "Testing GET route '/api/v1/cities/:cityID'..."
#curl -b cookie.txt $API_URL/api/v1/cities/6; echo
#echo
#echo "Testing PUT route '/api/v1/cities/:cityID'..."
#curl -b cookie.txt -X PUT -H 'Content-Type: application/json' -d '{"name":"Osaka"}' $API_URL/api/v1/cities/6; echo
#echo
#echo "Testing GET route after PUT '/api/v1/cities/:cityID'..."
#curl -b cookie.txt $API_URL/api/v1/cities/6; echo
#echo
#echo "Testing DELETE route '/api/v1/cities/:cityID'..."
#curl -b cookie.txt -X DELETE $API_URL/api/v1/cities/6; echo
#echo
#echo "Testing GET route after DELETE '/api/v1/cities/:cityID'..."
#curl -b cookie.txt $API_URL/api/v1/cities/6; echo
#
## Testing '/api/v1/auth/logout'.
#echo
#echo "Testing POST route '/api/v1/auth/logout'..."
#curl -X POST $API_URL/api/v1/auth/logout; echo
#
## Finish end-to-end testing.
#rm cookie.txt
#echo "Finished testing the application!"

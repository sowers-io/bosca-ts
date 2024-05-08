Getting Started
===

Required Development Software
--
* Docker or Podman
* Docker Compose
* golang 1.22.2
* protobuf 26.1
* [Zed Client](https://github.com/authzed/zed) 
* jq

Setup Development Environment
--
* Clone https://github.com/sowers-io/bosca
* Run `./scripts/initialize-dev`

Next Steps
-- 
* You can find environment variables needed during development in `.env-dev`
* You can find the http calls necessary to create a user at `test/user.http`
* Once you create a user, you'll need to verify it (you can look in the kratos database in the courier_messages table for the link).
* You can give the user access to the administrators group by running `./services/scripts/add-permissions <userid>`
* You can find the necessary Meilisearch keys by running (you'll need to set those as environment variables to index and search):
```bash
 curl \
  -X GET 'http://localhost:7700/keys' \
  -H 'Authorization: Bearer p8JcB_HuMHRxN7uVXfrG2wU06b5k7oTvaAAYo6nsi9M' \
  -H 'Accept: application/json' | jq 
```
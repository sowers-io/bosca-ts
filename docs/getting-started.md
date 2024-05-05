Getting Started
===

Required Development Software
--
* Docker or Podman
* Docker Compose
* golang 1.22.2
* protobuf 26.1
* [Zed Client](https://github.com/authzed/zed) 

Setup Development Environment
--
* Clone https://github.com/sowers-io/bosca
* Run `./scripts/initialize-dev`

Next Steps
-- 
* You can find environment variables needed during development in `.env-dev`
* You can find the http calls necessary to create a user at `test/user.http`
* Once you create a user, you'll need to verify it.
* You can give the user access to the administrators group by running `./services/scripts/add-permissions <userid>`
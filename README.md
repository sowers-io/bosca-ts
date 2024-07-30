Bosca
===

**Current Project Phase:** Pre-Alpha

See [Getting Started](docs/getting-started.md) for development instructions.

Bosca is an Open Source (Apache 2.0) AI powered content management system.

Key Features (not an exhaustive list):

* Content Management
* Collections
* Semantic Search
* Accounts, Authentication, Authorization, etc.
    * See [Ory](https://www.ory.sh/)
    * See [Authzed](https://www.authzed.com/)
* Content Workflows
* Personalization
    * Profiles
    * Recommendations
* Client SDKs (GraphQL)

Roadmap:

* Stabilize Service Interfaces
* Build Administration Interfaces
* Stabilize Kubernetes Environment
* ...

## Project Structure

* **/docs** - Documentation for Bosca
* **/workspace** - Contains core, frontend, and Bible components
* **/cli** - Contains the Bosca Command Line Interface
* **/protobuf** - Contains protobuf definitions for the API and Clients
* **/database** - Contains various database scripts and definitions
* **/services** - Contains a development environment for the services Bosca leverages and integrates
* **/scripts** - Various scripts for development and deployment
* **/.run** - Contains scripts for running Bosca Commands in IntelliJ (using the GoLand plugin)

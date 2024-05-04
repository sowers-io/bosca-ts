Bosca
===

**Current Project Phase:** Pre-Alpha

See [Getting Started](docs/getting-started.md) for development instructions.

This project is fairly new. I've only been working on it for about a week now (though, I've been conceptualizing it in 
my head for quite some time), and it seems like a perfect opportunity to start developing it in the open.

Bosca is an Open Source (Apache 2.0) AI powered content management system.

Key Features:

* Content Management
* Collections
* Semantic Search
* Accounts, Authentication, Authorization, etc.
    * See [Ory](https://www.ory.sh/)
    * See [Authzed](https://www.authzed.com/)
* Content Workflows
* Content Processing
    * Videos
    * Images
    * Text
* Content Graph
* Personalization
    * Profiles
    * Recommendations
* Client SDKs (GraphQL)

Roadmap:

* Convert C# prototype to Go (there have been a few prototypes: Java and a smaller one in Node)
* Stabilize Kubernetes environment
* ...

This system leverages a monorepo containing all the components for Bosca. However, the build pipeline will deploy 
Bosca in a microservice style architecture.  There are some exceptions to this rule, such as components that aren't part
of the Bosca Core.  For example, administrative interfaces that will be part of a commercial offering (or, at least that
is the current thinking, open to other thoughts about how to make Bosca sustainable).

* **/api** - Contains backend API components
* **/protobuf** - Contains protobuf definitions for the API and Clients
* **/conf** - Contains various configuration files

Technologies:

These are the general directions for the technology stack of Bosca. There will be some variations in a production
environment, such as possibly using Vertex AI over Ollama for processing different LLM related functions. But to allow
for local only processing, something like Ollama will be integrated into the system.

To enable these types of variations, there will be modular microservices and components that can be picked from. For
instance, you may be running in Google Cloud and want to use GCP cloud buckets for storage of assets. The default
implementation (for portability) will be MinIO, but with minimal effort, you can leverage a different storage provider.
In the interest of being fully functional as quickly as possible, we will forgo certain implementations until a later
times. But we will try and make decisions that support the modularity needed to make Bosca flexible, such that it can
fit into any organization's architecture well.

* Go
* GRPC
* Temporal
* Ory
* NextJS
* PostgreSQL
* Kubernetes
* Qdrant
* LangChain
* TypeScript
* Ollama
* MinIO
* GraphQL
* Ray

For more information, you can look in the /docs folder

Architecture:

<img src="docs/images/Bosca%20Infrastructure.png"  alt=""/>
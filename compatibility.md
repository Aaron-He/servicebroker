# Platform Compatibility for OSBAPI

#### Previous releases
| Release or Feature | Introduced | Deprecated | Cloud Foundry | Kubernetes |
| --- | -- | --- | --- | --- |
| `credentials` binding | v2.10 | | ✔️ | ✔️ |
| `syslog_drain` binding | v2.10 | | ✔️ | - |
| `route_forwarding` binding | v2.10 | | ✔️ | - |
| `volume_mounts` binding | v2.10 | | ✔️ | - |
| [*v2.11*](release-notes.md#v211) | Nov 15, 2016 | | ✔️ | ✔️ |
| Bindable and non-bindable plans | v2.11 | | ✔️ | ✔️ |
| `context` for creating and updating a service instance | v2.12 | | ✔️ | ✔️ |
| [*v2.12*](release-notes.md#v212) | June 13, 2017 | | ✔️ | ✔️ |
| `schemas` in catalog | v2.13 | | ✔️ | ✔️ |
| `context` for creating a service binding | v2.13 | | ✔️ | ✔️ |
| `originating identity` header | v2.13 | | ✔️ | ✔️ |
| Opaque Bearer Token Authentication | v2.13 | | - | ✔️ |
| [*v2.13*](release-notes.md#v213) | Sep 27, 2017 | | ✔️ | ✔️ |


#### Coming soon
| Release or Feature | Introduced | Deprecated | Cloud Foundry | Kubernetes |
| --- | -- | --- | --- | --- |
| Async Bindings | v2.14 | | - | ✔️ |
| GET endpoint for Service Instances | v2.14 | | ✔️ | - |
| GET endpoint for Service Bindings | v2.14 | | ✔️ | ✔️ |
| `dashboard_url` is updatable on Service Instance update | v2.14 | | - | - |
| *v2.14* | _TBD_ | | - | - |
| Generic Extensions (Actions) | [_OPEN_](https://github.com/openservicebrokerapi/servicebroker/pull/431) | | - | - |
| JSON Schema Endpoint | [_OPEN_](https://github.com/openservicebrokerapi/servicebroker/pull/402) | | - | - |
| JSON Schema for Responses | [_OPEN_](https://github.com/openservicebrokerapi/servicebroker/pull/392) | | - | - |
| Deprecated Classes and Plans | [_OPEN_](https://github.com/openservicebrokerapi/servicebroker/pull/504) | | - | - |

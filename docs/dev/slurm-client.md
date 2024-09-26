# Slurm Client

<!-- mdformat-toc start --slug=github --no-anchors --maxlevel=3 --minlevel=1 -->

- [Slurm Client](#slurm-client)
  - [Overview](#overview)
  - [Design](#design)
    - [Sequence Diagram](#sequence-diagram)

<!-- mdformat-toc end -->

## Overview

This Go library is responsible for enabling communication with [Slurm]. It
relies on communication with [slurmrestd] and `auth/jwt` based authentication
via the [rest-api].

## Design

The Slurm client abstracts the Slurm OpenAPI Spec endpoints and objects for easy
use. It supports a number of Slurm versions.

### Sequence Diagram

```mermaid
sequenceDiagram
    autonumber

    actor User as User
    participant SC as Slurm Client
    box Slurm Client Internals
        participant IC as Informer Cache
        participant OI as Object Informer
        participant OEC as Object Event Channel
        participant OH as Object Handler
    end %% Slurm Client Internals
    participant SAPI as Slurm REST API

    User->>+SC: Get Slurm Object (e.g. Node, Job)
    IC-->>SC: Lookup Object Informer
    alt Informer Started
    SC->>+IC: Get Slurm Object
        loop Object Informer Poll
            note over OI: async go routine
            OI->>IC: HasSynced=false
            OI->>+SAPI: Get Slurm Object
            SAPI-->>-OI: Return Slurm Object
            OI->>IC: Update Object Cache
            OI->>OEC: Emit Cache Event
            OI->>IC: HasSynced=true
        end %% loop Informer Poll
        loop Cache Event Handler
            note over OH: async go routine
            OEC-->>OH: Watch Cache Event
            OH->>OH: Run Handler for Object
        end %% loop Event Handle
        IC->>IC: Wait for HasSynced
        IC-->>-SC: Return Slurm Object
    else Informer Not Started
        SC->>+SAPI: Get Slurm Object
        SAPI-->>-SC: Return Slurm Object
    end %% alt Informer Started
    SC-->>-User: Return Slurm Object
```

<!-- Links -->

[rest-api]: https://slurm.schedmd.com/rest_api.html
[slurm]: https://www.schedmd.com/slurm
[slurmrestd]: https://slurm.schedmd.com/slurmrestd.html

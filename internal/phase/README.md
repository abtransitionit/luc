# How it works
Each `leaf` folder 
  - Denotes a set of phases involved in a `cobra` cmd 
  - Contains the following files:

  |name|type|description|
  |-|-|-|
  |`00_phase.go`|file|define a set of phases|
  |`xxx.go`|file|define a function denoting a **phase** |

Each phase:
- runs sequentially
- may:
  - Launch a go pipeline with tasks that runs concurently (into the phase)
  - Do other processing




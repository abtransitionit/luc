# How it works
- each `leaf` folder denotes a set of process involved in a pipeline
- each `leaf` folder have the following files
  - `00_ep.go` that **launch** the pipeline
  - `01_source.go` that is the **source** stage
  - `09_end.go` that denotes the **end** of the pipeline



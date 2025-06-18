# How it works
Each `leaf` folder denotes a set of process involved in a pipeline and contains the following files:
  |name|description|
  |-|-|
  |`00_run.go`|contains the function `RunPipeline()` that launches the pipeline|
  |`01_source.go`|the **source** stage of the pipeline|
  |`02_xxx.go`|the **first** stage after the **source** stage|
  |`99_last.go`|the **last** foreground process to be played after the pipeline|
  |`types.go`|the **struct** used to create the piplined data|
  |`xxx.go`|the **other** stages of the pipeline|


# The `RunPipeline()` function
- defines the pipeline stage order
# The `source` stage
- Usualy defines the instance structure that will be pipelined
- Initiate all the instance structure (if there is) that will be pipelined

# The `first` stage
- the first running immediatly after the `source` stage

# The `last` step
- A standard functions.
- **Collects** each pipelined instance, one by one.
- Often indicates, for each instance, if it's OK or KO

# The `types` files
- Often defines, the structure of the pipelined data

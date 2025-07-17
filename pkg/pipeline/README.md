# What does this folder contains
- Each folder denotes a **pipeline** with:
  - black boxes each denoting a process involved in the pipeline
  - a set of instance data flowing from **output** of box (n-1) to **input** of box (n)

- The pipeline itself can be seen as a reusable **black box** that needed specific inputs.  

# How it works
Each folder contains the following files:
  |name|description|
  |-|-|
  |`00_run.go`|contains the function `RunPipeline()` that launches the pipeline|
  |`01_source.go`|the **source** stage of the pipeline|
  |`02_xxx.go`|the **first** stage after the **source** stage|
  |`99_last.go`|the **last** foreground process to be played after the pipeline|
  |`types.go`|the **struct** used to create the piplined data|
  |`xxx.go`|the **other** stages of the pipeline|


# The `RunPipeline()` function
- defines the pipeline stages order
# The `source` stage
- Defines the the set of instances (from a data structure defined in file `types.go`)
- Send each instance into the pipeline

# The `first` stage
- The first stage running immediatly after the `source` stage

# The `last` step
- A standard functions.
- **Collects** each pipelined instance, one by one.
- Often indicates, for each instance, if it's OK or KO

# The `types` files
- Defines the structure of the data instance to be pipelined

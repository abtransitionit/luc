# Architecture
This section define the software architecture used to build LUC
- the `LUC` CLI or a `cobra` cmd is a set of `cobra` cmd
> $ luc
```bash
LUC (aka. Linux Unified CLI) ....

Usage:
  luc [command]

Available Commands:
  cli         Manage GO CLI.
  kbe         manage Kubernetes clusters (KBE = Kubernetes Easy).
  kind        manage Kind clusters.
  ...
```

> $ luc kind
```bash
manage Kind clusters.

Usage:
  luc kind [flags]
  luc kind [command]

Available Commands:
  provision   deploy a Kind cluster on a VM.
  ...
```
- a child `cobra` cmd is a set of `LUC` phases that can be all played in batch mode or individually

> $ luc kind provision
```bash
deploy a Kind cluster on a VM.

Usage:
  luc kind provision [phase name] [flags]

Flags:
  -r, --runall   Run all phases in sequence in batch mode
  -s, --show     List all available phases
```

> $ luc kbe init
```bash
deploy a kubernetes cluster. It becomes a KBE (Kubernetes Easy) cluster.

Usage:
  luc kbe init [phase name] [flags]

Flags:
  -r, --runall   Run all phases in sequence in batch mode
  -s, --show     List all available phases
macamax:luc max$
```
# Example of LUC phases
> $ luc kbe init --show
```bash
ID   Name      Description
---  --------  -----------
1    ssh       check all VMs/Nodes are SSH reachable
2    cpluc     copy CLI LUC to all VMs
3    check     check basic metrics before starting deployment
4    update    upgrade Nodes OS packages and packages repositories to version latest.
5    sysctl    configure OS kernel modules and parameters.
6    selinux   configure selinux to permissive mode on Rhel/Fedora nodes
7    dnfapt    provision dnfapt repositories and packages.
8    service   configure Nodes OS services.
9    cplane    initialize the CPlane.
10   worker    initialize the Workers.
11   kubectl   provision the Kubectl CLI
12   helm      provision the Helm client CLI
13   cni       provision CNI plugin
14   health    check the KBE cluster health
```
> $ luc kind provision --show
```bash
ID   Name      Description
---  --------  -----------
1    update    upgrade Kind VM OS.
2    cli       provision needed CLI
3    service   configure OS services.
4    env       define and export necessary environment variables.
```

# LUC phases
- `LUC` phases are pieces of reusable code.
- Reusable code:
  - can be public (`pkg` package) or private (`intenal` package)
  - can be a `go` **pipeline** or standard `go` functiion

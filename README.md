
# Luc  
**L**inux **U**nified **C**ontrol is a universal Linux Operations CLI (**C**ommand **L**ine **I**nterface). A lightweight `Go`-powered toolkit for cross-platform Linux management.

[![LICENSE](https://img.shields.io/badge/license-Apache_2.0-blue.svg)](https://choosealicense.com/licenses/apache-2.0/)
![Go Version](https://img.shields.io/badge/go-%3E%3D1.21-blue)
[![Cobra](https://img.shields.io/badge/Cobra-CLI%20Framework-blueviolet)](https://github.com/spf13/cobra)
[![Viper](https://img.shields.io/badge/Viper-Config%20Mgmt-yellowgreen)](https://github.com/spf13/viper)
[![Zap](https://img.shields.io/badge/Zap-Logging-ff69b4)](https://go.uber.org/zap)

## Why LUC

- **U**nified **L**inux **C**ontrol
- Single binary managing resources for both **local OS** and **remote VMs**  
- Native support for RHEL, Debian, and derived distributions. Unified commandstm manage:
  - Packages and repositorires  
  - services 
  - kernel parameters and modules
- Zero dependencies, remote-ready architecture  
- Simple and easy

```sh
# Create a production-grade Kubernetes cluster (KBE) on a set of VMs  
luc kbe create runall  

# Deploy a Kind cluster (eg. for controller testing)
luc kind create runall  

# Install any tools (works on RHEL/Debian)  
luc os provision kubectl, helm, nerdctl, runc, ...

# Install any container runtime (works on RHEL/Debian)  
luc os provision containerd, crio 

# Install any OS packages/repo (cross-distro)  
luc os dnfapt repo    ...
luc os dnfapt package ...  

# And more ...
```

----


# Core Features

## Go-Powered Efficiency

Built with industry-standard libraries:  

| Library | Role | Key Benefit |  
|-|-|-|
| **[Cobra](https://github.com/spf13/cobra)** | CLI Framework | Simplifies evolution and maintenance with production-grade command structures |  
| **[Viper](https://github.com/spf13/viper)** | Configuration | Unified config management across JSON/YAML/env vars |  
| **[Zap](https://go.uber.org/zap)** | Logging | High-performance structured logging with minimal overhead |  

## Extensible Architecture
Luc's extensible architecture is built on core design principles.

|Principle|Implementation|Benefit|
|-|-|-|
| **Modular Command System** | Hierarchical Cobra command structure    | Add new features without breaking changes; maintain clean separation of concerns |
| **Shared Core Libraries**  | Centralized logging/config/error packages | Uniform behavior across all commands; simplified maintenance            |
| **Plugin-Friendly**       | Go-native package imports               | Extend functionality by importing external libraries as dedicated modules |


## CI/CD Ready

Luc fits seamlessly into pipelines:  

**Zero-interaction**  

Commands run non-interactively (no prompts)  
```sh
luc kbe create runall --force # No manual approval needed
```  
**Pipeline-friendly**  

Pre-built for GitHub Actions/Jenkins/GitLab:  
```yaml
# GitHub Actions Example
- name: Deploy cluster  
  run: luc kbe create runall --file config.yaml
```  

**Machine outputs**  
`--json`/`--quiet` flags for automation  

**Secure**  
Auth via `$ENV_VARS` only (no CLI inputs)  

## Use Cases

- **Kubernetes Management**: create production-grade k8s cluster on demand
- **Infra-as-CLI**: Replace complex scripts with declarative commands  

---

# Getting Started  

---

# Contributing  

We welcome contributions! Before participating, please review:  
- **[Code of Conduct](.github/CODE_OF_CONDUCT.md)** – Our community guidelines.  
- **[Contributing Guide](.github/CONTRIBUTING.md)** – How to submit issues, PRs, and more.  

----

# Release History & Changelog  

Track version updates and changes:  
- **📦 Latest Release**: `vX.X.X` ([GitHub Releases](#))  
- **📄 Full Changelog**: See [CHANGELOG.md](CHANGELOG.md) for detailed version history.  

---



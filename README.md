# ben

- [ben](#ben)
  - [📦 Installation](#-installation)
    - [Installation via `go install`](#installation-via-go-install)
    - [Installation via `go get`](#installation-via-go-get)
  - [🧩 Explanation](#-explanation)
    - [Sheet](#sheet)
    - [Job](#job)
    - [Step](#step)
  - [🧅 Examples](#-examples)

## 📦 Installation

### Installation via `go install`

**This installation type should be chosen for system with at least Go 1.17.**

```shell
go install github.com/lukasl-dev/ben@latest
```

### Installation via `go get`

**This installation method should be chosen for systems with older Go versions.**

```shell
go get -u github.com/lukasl-dev/ben
```

## 🧩 Explanation

### Sheet

A sheet represents the configuration of one or more [jobs](#job). Its
responsibility is to provide a configuration so Ben can set up an
environment or even a complete project.

### Job

A job represents a list of tasks (=[steps](#step)) that must be completed to
achieve a desired goal. See [Examples](#examples) for a more detailed insight.

### Step

A Step represents a subtask of a [Job](#job). It can perform different types of
tasks. A Step can currently perform the following types of tasks:

- Executing commands
- Copying files and directories

## 🧅 Examples

- [Cloning a repository and run `npm install` afterwards.](examples/clone-npm.yml)

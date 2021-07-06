# | <img src="docs/images/logo.png" alt="Fury Logo" width="18" height="25"> |  fip-commons |

![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/sighupio/fip-commons)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/sighupio/fip-commons)
![GitHub](https://img.shields.io/github/license/sighupio/fip-commons)
[![Go Report Card](https://goreportcard.com/badge/github.com/sighupio/fip-commons)](https://goreportcard.com/report/github.com/sighupio/fip-commons)

## Table of Contents

- [| <img src="docs/images/logo.png" alt="Fury Logo" width="18" height="25"> |  fip-commons |](#----fip-commons-)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Getting Started](#getting-started)
    - [Installation](#installation)
    - [Usage](#usage)
    - [Examples](#examples)
  - [Developer Guide](#developer-guide)
  - [License](#license)

## Overview

This library contains common packages that can be used from multiples `golang` components in the
Fury Intelligent platform architecture.

The first release includes a package to solve the problem while connecting to the Kubernetes API.

## Getting Started

### Installation

The simplest way to use this package is to download the package from the source repo as follows:

* Using Go get

```sh
$ go get github.com/sighupio/fip-commons
#
```

### Usage

To use this library add this dependency to your `go.mod` file by running the
`go get go get github.com/sighupio/fip-commons` command. Then, `import` it in your golang codebase and use it:

```go
package demo

import (
  "context"
  "fmt"

  "github.com/sighupio/fip-commons/pkg/kube" // Import it
)

// demo is an example implementation
func demo(){
  k := kube.KubernetesClient{KubeConfig: "/home/my-user/my-kubeconfig-path"}
  k.Init()
  ctx := context.TODO()
  err = k.Healthz(&ctx)
  if err != nil {
    fmt.Println("error. cluster seems to be not healthy")
  }
}
```

### Examples

This repository contains an example implementation that list namespaces using the `KubernetesClient` exposed in this
golang package. It is available in the `examples` directory. Follow the usage information in the
[corresponding README](./examples/) for more info.

## Developer Guide

To set the code up locally, build, run tests, etc. Please refer the
[contributor's guide](./CONTRIBUTING.md).

## License

Check the [License here](LICENSE)

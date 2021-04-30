<h1 align="center">GoCreate</h1>
<p align="center">
  <a href="https://github.com/complexorganizations/gocreate/releases">
    <img alt="Release" src="https://img.shields.io/github/v/release/complexorganizations/gocreate" target="_blank" />
  </a>
  <a href="https://github.com/complexorganizations/gocreate/issues">
    <img alt="Issues" src="https://img.shields.io/github/issues/complexorganizations/gocreate" target="_blank" />
  </a>
  <a href="https://github.com/sponsors/Prajwal-Koirala">
    <img alt="Sponsors" src="https://img.shields.io/static/v1?label=Sponsor&message=%E2%9D%A4&logo=GitHub" target="_blank" />
  </a>
  <a href="https://github.com/complexorganizations/gocreate/pulls">
    <img alt="PullRequest" src="https://img.shields.io/github/issues-pr/complexorganizations/gocreate" target="_blank" />
  </a>
  <a href="https://raw.githubusercontent.com/complexorganizations/gocreate/main/.github/license">
    <img alt="License" src="https://img.shields.io/github/license/complexorganizations/gocreate" target="_blank" />
  </a>
  <a href="https://goreportcard.com/report/github.com/complexorganizations/gocreate">
    <img alt="GoReport" src="https://goreportcard.com/badge/github.com/complexorganizations/gocreate" target="_blank" />
  </a>
  <a href="https://github.com/complexorganizations/gocreate/actions">
    <img alt="Actions" src="https://github.com/complexorganizations/gocreate/workflows/Go/badge.svg" target="_blank" />
  </a>
  <a href="https://github.com/complexorganizations/gocreate">
    <img alt="Go-Version" src="https://img.shields.io/github/go-mod/go-version/vx3r/wg-gen-web" target="_blank" />
  </a>
  <a href="https://github.com/complexorganizations/gocreate">
    <img alt="Code-Size" src="https://img.shields.io/github/languages/code-size/vx3r/wg-gen-web" target="_blank" />
  </a>
</p>

<p align="center">
  <a href="https://snapcraft.io/gocreate">
    <img alt="Get it from the Snap Store" src="https://snapcraft.io/static/images/badges/en/snap-store-black.svg" />
  </a>
</p>

There is no official go project specification framework or widely agreed standard other than `go.mod` and this tool is based on the golang-standard project.

---
### Features
- Structure go applications

---
### Installation
Download the latest `gocreate` binary
```
go get -v -u github.com/complexorganizations/gocreate
```
Create a project applications using `gocreate`
```
gocreate -name={PROJECT}
```

---
### Structure
```
{PROJECT}
├── api
│   └── README.md
├── assets
│   └── README.md
├── build
│   ├── ci
│   │   └── README.md
│   ├── package
│   │   └── README.md
│   └── README.md
├── cmd
│   ├── README.md
│   └── {PROJECT}
│       └── README.md
├── configs
│   └── README.md
├── deployments
│   └── README.md
├── Dockerfile
├── docs
│   └── README.md
├── examples
│   └── README.md
├── githooks
│   └── README.md
├── go.mod
├── go.sum
├── init
│   └── README.md
├── internal
│   ├── app
│   │   └── {PROJECT}
│   │       └── README.md
│   ├── pkg
│   │   └── {PROJECT}
│   │       └── README.md
│   └── README.md
├── main.go
├── pkg
│   ├── README.md
│   └── {PROJECT}
│       └── README.md
├── README.md
├── scripts
│   └── README.md
├── test
│   └── README.md
├── third_party
│   └── README.md
├── tools
│   └── README.md
├── vendor
│   └── README.md
├── web
│   ├── app
│   │   └── README.md
│   ├── README.md
│   ├── static
│   │   └── README.md
│   └── template
│       └── README.md
└── website
    └── README.md
```

---
### Compatibility
| OS              | Supported          |
| --------------  | ------------------ |
| Darwin          |:heavy_check_mark:  |
| FreeBSD         |:heavy_check_mark:  |
| Linux           |:heavy_check_mark:  |
| NetBSD          |:heavy_check_mark:  |
| OpenBSD         |:heavy_check_mark:  |
| Plan 9          |:heavy_check_mark:  |
| Solaris         |:heavy_check_mark:  |
| Windows         |:heavy_check_mark:  |

---
### Author
* Name: Prajwal Koirala
* Website: [prajwalkoirala.com](https://www.prajwalkoirala.com)

---
### Credits
- [golang-standards](https://github.com/golang-standards/project-layout)
- Open Source Community

---
### License

Copyright © 2021 [ComplexOrganizations](https://github.com/complexorganizations)

This project is unlicensed.

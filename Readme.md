# nilfs - Abstract File System interface

[![Go Reference](https://pkg.go.dev/badge/github.com/absfs/nilfs.svg)](https://pkg.go.dev/github.com/absfs/nilfs)
[![Go Report Card](https://goreportcard.com/badge/github.com/absfs/nilfs)](https://goreportcard.com/report/github.com/absfs/nilfs)
[![CI](https://github.com/absfs/nilfs/actions/workflows/ci.yml/badge.svg)](https://github.com/absfs/nilfs/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

The `nilfs` package implements the absfs.FileSystem interface. This implementation does nothing and returns no errors. Useful as a template for new filesystems, or where a noop FileSystem might be needed to to satisfy an argument list.

## Install 

```bash
$ go get github.com/absfs/nilfs
```


## absfs
Check out the [`absfs`](https://github.com/absfs/absfs) repo for more information about the abstract filesystem interface and features like filesystem composition.

## LICENSE

This project is governed by the MIT License. See [LICENSE](https://github.com/absfs/nilfs/blob/master/LICENSE)




# nilfs - Abstract File System interface
The `nilfs` package implements the absfs.FileSystem interface. This implementation does nothing and returns no errors. Useful as a template for new filesystems, or where a noop FileSystem might be needed to to satisfy an argument list.

## Install 

```bash
$ go get github.com/absfs/nilfs
```


## absfs
Check out the [`absfs`](https://github.com/absfs/absfs) repo for more information about the abstract filesystem interface and features like filesystem composition.

## LICENSE

This project is governed by the MIT License. See [LICENSE](https://github.com/absfs/osfs/blob/master/LICENSE)




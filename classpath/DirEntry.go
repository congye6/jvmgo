package classpath

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

//管理一个文件夹的class
type DirEntry struct {
	absPath string
}

func (this DirEntry) ReadClass(path string) (clazz []byte,entry Entry,err error) {
	if  !strings.Contains(path,this.absPath){
		println(this.absPath+" "+path)
		return nil, nil, nil
	}
	println(this.absPath)
	clazz,err=ioutil.ReadFile(path)
	return clazz, this, err
}

func newDirEntry(path string) DirEntry {
	path,err:=filepath.Abs(path)
	if err !=nil{
		panic(err)
	}
	return DirEntry{path}
}

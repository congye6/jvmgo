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

func (this DirEntry) readClass(path string) (clazz []byte,entry Entry,err error) {
	if  !strings.Contains(path,this.absPath){
		return nil, nil, nil;
	}
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

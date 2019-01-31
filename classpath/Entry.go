package classpath

import (
	"fmt"
	"os"
	"strings"
)

const PATH_LIST_SPLITER = string(os.PathListSeparator)

type Entry interface {
	readClass(path string) (clazz []byte,entry Entry,err error)
}

func newEntry(path string) Entry {
	fmt.Println(path)
	if strings.Contains(path, PATH_LIST_SPLITER){
		return newCompositeEntry(path)
	}
	if isZip(path){
		return newZipEntry(path)
	}
	if strings.HasSuffix(path,"*"){
		return newWildcardEntry(path)
	}
	return newDirEntry(path)
}

//判断路径是否为zip
func isZip(path string) bool  {
	return strings.HasSuffix(path,".zip") ||
		strings.HasSuffix(path,".jar")
}

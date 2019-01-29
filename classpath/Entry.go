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
	if strings.HasSuffix(path,".zip") ||
		strings.HasSuffix(path,".jar"){
		return newZipEntry(path)
	}
	return newDirEntry(path)
}

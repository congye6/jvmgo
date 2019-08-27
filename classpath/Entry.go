package classpath

import (
	"fmt"
	"os"
	"strings"
)

const PATH_LIST_SPLITER = string(os.PathListSeparator)

// 管理某个路径下的class文件
type Entry interface {
	ReadClass(path string) (clazz []byte, entry Entry, err error)
}

func NewEntry(path string) Entry {
	fmt.Printf("【DEBUG】new entry path:%s \n", path)
	if strings.Contains(path, PATH_LIST_SPLITER) {
		return newCompositeEntry(path)
	}
	if isZip(path) {
		return newZipEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	return newDirEntry(path)
}

//判断路径是否为zip
func isZip(path string) bool {
	return strings.HasSuffix(path, ".zip") ||
		strings.HasSuffix(path, ".jar")
}

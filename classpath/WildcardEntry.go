/**
 * @Author: congye6
 * @Date: 2019-01-31 11:44
 */
package classpath

import (
	"os"
	"path/filepath"
)

func newWildcardEntry(path string) CompositeEntry {
	dirPath:=path[:len(path)-1]//获取文件夹路径
	compositeEntry:=CompositeEntry{}

	walkFunction:= func(path string,info os.FileInfo,err error) error {//遍历文件的策略
		if err!=nil{
			return err
		}
		if info.IsDir()&&path!=dirPath{
			return filepath.SkipDir
		}
		if isZip(path){//遇到zip文件就添加entry
			compositeEntry=append(compositeEntry, newZipEntry(path))
		}
		return nil
	}
	filepath.Walk(dirPath,walkFunction)
	return compositeEntry
}

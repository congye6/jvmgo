package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

/**
 * jar包或zip类路径
 */
type ZipEntry struct {
	absPath string
}

func (this ZipEntry) readClass(path string) (clazz []byte,entry Entry,err error){
	reader,err:=zip.OpenReader(this.absPath)
	if err!=nil{
		return nil,nil,err
	}
	defer reader.Close()
	for _,file := range reader.File{//遍历zip中的文件，查找相同文件名的文件
		if file.Name!=path{
			continue
		}
		fileReader,err:=file.Open()//打开文件
		if err!=nil{
			return nil, nil, err
		}
		defer fileReader.Close()
		data,err:=ioutil.ReadAll(fileReader)//读取文件
		if err!=nil{
			return nil, nil, err
		}
		return data,this,nil
	}
	return nil, nil, errors.New("class not found")
}

func newZipEntry(path string) (entry ZipEntry) {
	path,err:=filepath.Abs(path)
	if err!=nil{
		panic(err)
	}
	return ZipEntry{path}
}
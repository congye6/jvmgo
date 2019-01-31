package classpath

import "strings"

type CompositeEntry []Entry

func (this CompositeEntry) readClass(path string) (clazz []byte,entry Entry,err error){
	for _,entry= range this{
		if clazz,entry,err=entry.readClass(path);clazz!=nil{
			return clazz,entry,err
		}
	}
	//没有entry能读取
	return nil, nil, nil
}

func newCompositeEntry(path string) CompositeEntry{
	compositeEntry:=[]Entry{}
	for _,subPath := range strings.Split(path, PATH_LIST_SPLITER){
		compositeEntry=append(compositeEntry, newEntry(subPath))
	}
	return compositeEntry
}
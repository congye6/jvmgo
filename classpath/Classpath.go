package classpath

/**
 * 启动类加载器路径
 */
const BOOTSTRAP_PATH  = "\\lib\\*"

/**
 * 扩展类加载器
 */
const EXT_PATH  = "\\lib\\ext\\*"

const DEFUALT_CLASS_PATH  = "E:\\workspace\\practice\\bin"

type Classpath struct {
	bootstrap Entry
	ext Entry
	application Entry
}

func Parse (jrePath string,applicationPath string) *Classpath {
	application:=newEntry(applicationPath)
	bootstrap:=newEntry(jrePath+BOOTSTRAP_PATH)
	ext:=newEntry(jrePath+EXT_PATH)
	return &Classpath{bootstrap,ext,application}
}

func (classpath *Classpath) ReadClass(path string) (data []byte,err error) {
	entries:=[]Entry{classpath.bootstrap,classpath.ext,classpath.application}
	for _,entry:= range entries{
		data,_,err=entry.readClass(path)
		if data!=nil{
			return data,err
		}
	}
	return nil, nil
}
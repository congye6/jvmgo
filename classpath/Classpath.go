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

// 初始化加载class文件的entry对象
func Parse (jrePath string,applicationPath string) *Classpath {
	application:=newEntry(applicationPath) // 应用类加载器
	bootstrap:=newEntry(jrePath+BOOTSTRAP_PATH) // 启动类加载器
	ext:=newEntry(jrePath+EXT_PATH) // 扩展类加载器
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
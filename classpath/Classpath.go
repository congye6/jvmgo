package classpath

/**
 * 启动类加载器路径
 */
const BOOTSTRAP_PATH  = "\\lib\\rt.jar"

/**
 * 扩展类加载器
 */
const EXT_PATH  = "\\lib\\ext"

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
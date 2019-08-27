package method_area

import "jvmgo/classpath"

/**
 * 启动类加载器路径
 */
const BOOTSTRAP_PATH = "/lib/*"

/**
 * 扩展类加载器
 */
const EXT_PATH = "/lib/ext/*"

const DEFUALT_CLASS_PATH = "E:\\workspace\\practice\\bin"

// 类加载的第一部分：加载类文件
func GetApplicationLoader(jrePath string, applicationPath string) *ClassLoader {
	application := classpath.NewEntry(applicationPath)        // 应用类加载器
	bootstrap := classpath.NewEntry(jrePath + BOOTSTRAP_PATH) // 启动类加载器
	ext := classpath.NewEntry(jrePath + EXT_PATH)             // 扩展类加载器

	bootstrapLoader := NewClassLoader(bootstrap, nil, "bootstrap")
	extLoader := NewClassLoader(ext, bootstrapLoader, "ext")
	applicationLoader := NewClassLoader(application, extLoader, "application")
	return applicationLoader
}

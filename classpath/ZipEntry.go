package classpath

/**
 * jar包或zip类路径
 */
type ZipEntry struct {

}

func (this ZipEntry) readClass(path string) (clazz []byte,entry Entry,err error){
	return nil, nil, nil
}

func newZipEntry(path string) (entry ZipEntry) {
	return ZipEntry{}
}
package classpath

import (
	"os"
	"strings"
)

//get the separator for different system, ; for windows
//and : for Unix
const pathListSeparator = string(os.PathListSeparator)

//Interface Entry is defined for readClass and return string
type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

//create different Entrys for different strings
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.Contains(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") ||
		strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, "zip") ||
		strings.HasSuffix(path, "ZIP") {
			return newZipEntry(path)
		}
	return newDirEntry(path)
}

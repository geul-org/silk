//ff:func feature=seg type=writer control=sequence
//ff:what File.Close — unmap
package seg

import "syscall"

func (f *File) Close() error {
	return syscall.Munmap(f.b)
}

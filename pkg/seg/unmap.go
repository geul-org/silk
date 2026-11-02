//ff:func feature=seg type=writer control=sequence
//ff:what Unmap — release bytes from Map
package seg

import "syscall"

func Unmap(b []byte) error {
	if b == nil {
		return nil
	}
	return syscall.Munmap(b)
}

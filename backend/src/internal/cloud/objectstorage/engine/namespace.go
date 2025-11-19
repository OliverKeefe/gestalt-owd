package objectstorage

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
)

/*

   Mount
   PID
   Network
   Cgroup
   IPC
   Time
   UTS
   User

*/

// create namespace / cgroup container for object storage.

func must(err error) {
	if err != nil {
		panic(err)
	}
}

//type Mount struct{}

func CreateNamespaceAndMount(mountPoint, nsBindPath string) error {

	// Unshare new mount namespace.
	if err := unix.Unshare(unix.CLONE_NEWNS); err != nil {
		return fmt.Errorf("unshare(CLONE_NEWNS): %w", err)
	}

	// Make mount propagation private to avoid leaks into host fs.
	err := unix.Mount("", "/", "", uintptr(unix.MS_REC|unix.MS_PRIVATE), "")
	if err != nil {
		return fmt.Errorf("unable to make root private: %w", err)
	}

	// Might need to be 0655 rather than 0755 perm bits.
	if err := os.MkdirAll(mountPoint, 0755); err != nil {
		return fmt.Errorf("error creating mount point: %w", err)

	if err := unix.Mount(nsBindPath, mountPoint, "tmpfs", 0, ""); err != nil {
		err := unix.Unmount(mountPoint, 0)
		if err != nil {
			return fmt.Errorf("unable to cleanup unmountable mount: %w", err)
		}
		return fmt.Errorf("unable to mount file system: %w", err)
	}
	return nil
}

//type PID struct{}
//
//type Network struct{}
//
//type CGroup struct{}
//
//type IPC struct{}
//
//type Time struct{}
//
//type UTS struct{}
//
//type User struct{}

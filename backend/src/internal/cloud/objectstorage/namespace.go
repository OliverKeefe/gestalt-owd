package objectstorage

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	//"os/exec"
	//"runtime"
	//"syscall"
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
	return nil
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

	//
	//	package main
	//
	//	import (
	//		"fmt"
	//	"os"
	//	"runtime"
	//	"syscall"
	//
	//	"golang.org/x/sys/unix"
	//	)
	//
	//	func main() {
	//		// 1. Lock the current goroutine to a single OS thread.
	//		// Namespaces are thread-local, so the thread must not move during or after the change.
	//		runtime.LockOSThread()
	//		defer runtime.UnlockOSThread()
	//
	//		targetPID := 1 // Target the init process's namespace (usually PID 1)
	//		nsPath := fmt.Sprintf("/proc/%d/ns/mnt", targetPID)
	//
	//		fmt.Printf("Attempting to switch to mount namespace of PID %d (%s)...\n", targetPID, nsPath)
	//
	//		// 2. Open the file descriptor for the target namespace.
	//		// We use O_RDONLY since we only need to reference the namespace object.
	//		fd, err := unix.Open(nsPath, unix.O_RDONLY, 0)
	//		if err != nil {
	//			fmt.Printf("Error opening namespace file %s: %v\n", nsPath, err)
	//			fmt.Println("Are you running as root? Is PID 1 correct?")
	//			return
	//		}
	//		defer unix.Close(fd)
	//
	//		// 3. Enter the new namespace using the file descriptor.
	//		// The NS_GET_MT is not required here; simply passing 0 as the second argument is standard for setting.
	//		if err := unix.Setns(fd, unix.CLONE_NEWNS); err != nil {
	//			fmt.Printf("Error calling Setns: %v\n", err)
	//			return
	//		}
	//
	//		fmt.Println("Successfully switched mount namespaces.")
	//
	//		// 4. Verify the change (optional but highly recommended).
	//		// Now when you list mounts or directories, you should see the target system's view.
	//		// In the case of switching to PID 1, listing "/" should show the host system root mounts.
	//
	//		fmt.Println("\n--- Listing root directory mounts within the new namespace ---")
	//		// Using exec.Command is an easy way to demonstrate the effect from Go
	//		cmd := exec.Command("ls", "-lha", "/")
	//		// Run the command on the *same* locked thread
	//		cmd.SysProcAttr = &syscall.SysProcAttr{
	//			// Pgid is set to 0 to keep the process in the same process group as the parent
	//		}
	//		// Note: Executing commands and ensuring they run on the *exact* same thread
	//		// as the Setns call can be complex in Go. The simple 'ls' call below might run
	//		// on a different thread if not carefully managed.
	//		// For verification, it's better to read and print files from Go stdlib functions.
	//
	//		files, err := os.ReadDir("/")
	//		if err != nil {
	//			fmt.Printf("Error reading directory: %v\n", err)
	//			return
	//		}
	//		for _, file := range files {
	//			fmt.Println(file.Name())
	//		}
}

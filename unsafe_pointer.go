package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	forgetObjects := []_ForgetOne{
		{NodeId: 1, Nlookup: 2},
	}

	secretData := []byte{42, 0, 0, 0, 0, 0, 0, 0, 42, 0, 0, 0, 0, 0, 0, 0}

	sH := (*reflect.SliceHeader)(unsafe.Pointer(&forgetObjects))
	sH.Cap *= 16
	sH.Len *= 16
	forgetObjectBytes := *(*[]byte)(unsafe.Pointer(sH))

	server := &Server{
		mountPoint: "/nonexisting",
		fileSystem: MyFileSystem{},
		opts:       &MountOptions{},
	}

	req := &request{
		inHeader: nil,
		inData:   unsafe.Pointer(&_BatchForgetIn{Count: 5}),
		arg:      forgetObjectBytes,
	}

	fmt.Printf("%p\n", &forgetObjectBytes[0])
	fmt.Printf("%p\n", &secretData[0])
	fmt.Println(secretData)
	fmt.Println()

	doBatchForget(server, req)

	_ = secretData
}

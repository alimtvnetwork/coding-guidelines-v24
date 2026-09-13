//go:build windows

package ostype

import (
	"strconv"
	"syscall"
	"unsafe"
)

const (
	hkeyLocalMachine = syscall.HKEY_LOCAL_MACHINE
	keyQueryValue    = syscall.KEY_QUERY_VALUE
	subKeyPath       = "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion"
)

func queryWindowsRegistry() (*WindowsSystemDetail, error) {
	subKeyPtr, err := syscall.UTF16PtrFromString(subKeyPath)
	if err != nil {
		return nil, err
	}

	var handle syscall.Handle
	err = syscall.RegOpenKeyEx(hkeyLocalMachine, subKeyPtr, 0, keyQueryValue, &handle)
	if err != nil {
		return nil, err
	}

	defer syscall.RegCloseKey(handle)

	productName := queryRegString(handle, "ProductName")
	installType := queryRegString(handle, "InstallationType")
	editionID := queryRegString(handle, "EditionID")
	currentBuildStr := queryRegString(handle, "CurrentBuildNumber")
	if currentBuildStr == "" {
		currentBuildStr = queryRegString(handle, "CurrentBuild")
	}

	buildNum, _ := strconv.Atoi(currentBuildStr)
	detail := ParseWindowsDetail(productName, installType, editionID, buildNum)
	detail.CompositionEdition = queryRegString(handle, "CompositionEditionID")
	detail.ReleaseId = queryRegString(handle, "ReleaseId")
	detail.BuildBranch = queryRegString(handle, "BuildBranch")
	detail.SystemRoot = queryRegString(handle, "SystemRoot")
	detail.RegisteredOwner = queryRegString(handle, "RegisteredOwner")

	return detail, nil
}

func queryRegString(handle syscall.Handle, valueName string) string {
	valNamePtr, err := syscall.UTF16PtrFromString(valueName)
	if err != nil {
		return ""
	}

	var valType uint32
	var bufLen uint32
	err = syscall.RegQueryValueEx(handle, valNamePtr, nil, &valType, nil, &bufLen)
	if err != nil || bufLen == 0 {
		return ""
	}

	buf := make([]uint16, bufLen/2+1)
	err = syscall.RegQueryValueEx(handle, valNamePtr, nil, &valType, (*byte)(unsafe.Pointer(&buf[0])), &bufLen)
	if err != nil {
		return ""
	}

	return syscall.UTF16ToString(buf)
}

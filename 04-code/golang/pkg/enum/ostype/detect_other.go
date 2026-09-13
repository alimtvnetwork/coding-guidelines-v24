//go:build !windows

package ostype

func queryWindowsRegistry() (*WindowsSystemDetail, error) {
	return nil, nil
}

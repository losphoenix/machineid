// +build windows

package machineid

import (
	"golang.org/x/sys/windows/registry"
)

// machineID returns the key MachineGuid in registry `HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Cryptography`.
// If there is an error running the commad an empty string is returned.
func machineID() (string, error) {
	t, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Services\Tcpip\Parameters`, registry.QUERY_VALUE|registry.WOW64_64KEY)
	if err != nil {
		return "", err
	}
	defer t.Close()
	y, _, err := t.GetStringValue("Domain")
	if err != nil {
		return "", err
	}
	if y != "" {
		return y, nil
	}

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Cryptography`, registry.QUERY_VALUE|registry.WOW64_64KEY)
	if err != nil {
		return "", err
	}
	defer k.Close()

	s, _, err := k.GetStringValue("MachineGuid")
	if err != nil {
		return "", err
	}
	return s, nil
}

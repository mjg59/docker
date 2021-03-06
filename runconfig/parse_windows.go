package runconfig

import (
	"fmt"
	"strings"
)

// ValidateNetMode ensures that the various combinations of requested
// network settings are valid.
func ValidateNetMode(c *Config, hc *HostConfig) error {
	// We may not be passed a host config, such as in the case of docker commit
	if hc == nil {
		return nil
	}
	parts := strings.Split(string(hc.NetworkMode), ":")
	switch mode := parts[0]; mode {
	case "default", "none":
	default:
		return fmt.Errorf("invalid --net: %s", hc.NetworkMode)
	}
	return nil
}

// ValidateIsolationLevel performs platform specific validation of the
// isolation level in the hostconfig structure. Windows supports 'default' (or
// blank), and 'hyperv'. These refer to Windows Server Containers and
// Hyper-V Containers respectively.
func ValidateIsolationLevel(hc *HostConfig) error {
	// We may not be passed a host config, such as in the case of docker commit
	if hc == nil {
		return nil
	}
	if !hc.Isolation.IsValid() {
		return fmt.Errorf("invalid --isolation: %q. Windows supports 'default' (Windows Server Container) or 'hyperv' (Hyper-V Container)", hc.Isolation)
	}
	return nil
}

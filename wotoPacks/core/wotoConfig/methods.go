package wotoConfig

import "net"

func (c *Config) IsServerExternal() bool {
	ip := net.ParseIP(c.Bind)
	return ip != nil && ip.IsGlobalUnicast()
}

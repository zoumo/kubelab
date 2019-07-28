package apps

import "github.com/zoumo/kubelab/apps/v1"

// Interface provides access to each of this group's versions.
type Interface interface {
	V1() v1.Interface
}

type group struct {
}

// New returns a new Interface.
func New() Interface {
	return &group{}
}

// V1 returns a new v1.Interface.
func (g *group) V1() v1.Interface {
	return v1.New()
}

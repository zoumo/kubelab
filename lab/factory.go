package lab

import "github.com/zoumo/kubelab/lab/apps"

// Interface provides useful utils for resources in all known API group versions
type Interface interface {
	Apps() apps.Interface
}

type kubelab struct{}

// New constructs a new instance of a kubelab
func New() Interface {
	return &kubelab{}
}

func (l *kubelab) Apps() apps.Interface {
	return apps.New()
}

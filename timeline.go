package capsule

import "github.com/sirupsen/logrus"

type TimelineStructure struct{}

func (ts *TimelineStructure) Build(n *Node, structure map[string]interface{}) error {
	logrus.Infof("building '%v'", n.FullPath())
	return nil
}


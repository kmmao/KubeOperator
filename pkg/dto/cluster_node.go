package dto

import (
	"github.com/KubeOperator/KubeOperator/pkg/model"
	v1 "k8s.io/api/core/v1"
)

type Node struct {
	model.ClusterNode
	Info v1.Node `json:"info"`
	Ip   string  `json:"ip"`
}

type NodeCreation struct {
	Hosts      []string `json:"hosts"`
	Nodes      []string `json:"nodes"`
	Increase   int      `json:"increase"`
	SupportGpu string   `json:"supportGpu"`
}

type NodePage struct {
	Items []Node `json:"items"`
	Total int    `json:"total"`
}

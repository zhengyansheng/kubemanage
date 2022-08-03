package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/public"
)

type DeploymentNameNS struct {
	Name      string `json:"name" form:"name" comment:"无状态控制器名称" validate:"required"`
	NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" validate:"required"`
}

type DeployCreateInput struct {
	//DeploymentNameNS
	Name          string            `json:"name" form:"name" comment:"无状态控制器名称" validate:"required"`
	NameSpace     string            `json:"namespace" form:"namespace" comment:"命名空间" validate:"required"`
	Replicas      int32             `json:"replicas" validate:"required" comment:"副本数"`
	Image         string            `json:"image" validate:"required" comment:"镜像名"`
	Labels        map[string]string `json:"labels" validate:"required" comment:"标签"`
	Cpu           string            `json:"cpu" validate:"required" comment:"Cpu限制"`
	Memory        string            `json:"memory" validate:"required" comment:"内存限制"`
	ContainerPort int32             `json:"container_port" validate:"required" comment:"容器端口"`
	HealthCheck   bool              `json:"health_check" validate:"required" comment:"健康检查开关"`
	HealthPath    string            `json:"health_path" validate:"required" comment:"Http健康检查路径"`
}

type UpdateDeployInput struct {
	Name      string `json:"name" form:"name" comment:"无状态控制器名称" validate:"required"`
	NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" validate:"required"`
	Content   string `json:"content" validate:"required" comment:"更新内容"`
}

type DeployScaleInput struct {
	Name      string `json:"name" form:"name" comment:"无状态控制器名称" validate:"required"`
	NameSpace string `json:"namespace" form:"namespace" comment:"命名空间" validate:"required"`
	ScaleNum  int    `json:"scale_num" form:"scale_num" comment:"期望副本数" validate:"required"`
}

type DeployListInput struct {
	FilterName string `json:"filter_name" form:"filter_name" validate:"required" comment:"过滤名"`
	NameSpace  string `json:"namespace" form:"namespace" validate:"required" comment:"命名空间"`
	Limit      int    `json:"limit" form:"limit" validate:"required" comment:"分页限制"`
	Page       int    `json:"page" form:"page" validate:"required" comment:"页码"`
}

func (params *DeployListInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

func (params *UpdateDeployInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

func (params *DeployCreateInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

func (params *DeployScaleInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

func (params *DeploymentNameNS) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

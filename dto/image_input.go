package dto

// 定义 ImageInput 结构体，用于从前端接收镜像相关信息
type ImageInput struct {
	Name            string   `json:"name"`            // 镜像名称
	Tag             string   `json:"tag"`             // 镜像标签
	TemplateName    string   `json:"templateName"`    // 镜像模板名称
	TemplateCommand []string `json:"templateCommand"` // 需要追加到模板的命令
}

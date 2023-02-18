package repositories

import (
	"go-to-cloud/internal/models/pipeline"
	"gorm.io/datatypes"
	"time"
)

type PipelineHistory struct {
	Model
	PipelineID   uint                    `json:"pipeline_id" gorm:"column:pipeline_id;index:pipeline_history_pipeline_id_index"`
	ProjectID    uint                    `json:"project_id" gorm:"column:project_id;index:pipeline_history_project_id_index"`
	Name         string                  `json:"name" gorm:"column:name;type:nvarchar(64)"`      // 计划名称
	Env          string                  `json:"env" gorm:"column:env;type:nvarchar(64)"`        // 运行环境(模板), e.g. dotnet:6; go:1.17
	SourceCodeID uint                    `json:"source_code_id" gorm:"column:source_code_id"`    // 代码仓库ID
	Branch       string                  `json:"branch" gorm:"column:branch;type:nvarchar(64)"`  // 分支名称
	Params       datatypes.JSON          `json:"params" gorm:"column:params"`                    // 本次运行的参数(json格式）
	CreatedBy    uint                    `json:"created_by" gorm:"column:created_by"`            // 构建人
	Remark       string                  `json:"remark" gorm:"column:remark"`                    // 备注
	BuildAt      time.Time               `json:"build_at" gorm:"column:build_at"`                // 运行时间
	BuildResult  pipeline.BuildingResult `json:"run_result" gorm:"column:run_result;default:99"` // 运行结果; 1：成功；2：取消；3：失败；0：从未执行; 99：正在构建
	BuildLog     string                  `json:"log" gorm:"column:log;type:text"`                // 构建日志
}

func (m *PipelineHistory) TableName() string {
	return "pipeline_history"
}

func UpdateBuildLog(historyId uint, log *string) {

}

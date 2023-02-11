package repositories

import (
	"go-to-cloud/conf"
	"go-to-cloud/internal/models/pipeline"
	"go-to-cloud/internal/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ArtifactScript struct {
	Dockerfile string `json:"dockerfile"`
	Registry   string `json:"registry"`
	IsSecurity bool   `json:"isSecurity"`
	Account    string `json:"account"`
	Password   string `json:"password"`
}

type Pipeline struct {
	Model
	PipelineSteps []PipelineSteps         `json:"-" gorm:"foreignKey:pipeline_id"`
	ProjectID     uint                    `json:"project_id" gorm:"column:project_id"`
	Name          string                  `json:"name" gorm:"column:name"` // 计划名称
	Env           string                  `json:"env" gorm:"column:env"`   // 运行环境(模板), e.g. dotnet:6; go:1.17
	SourceCodeID  uint                    `json:"source_code_id" gorm:"column:source_code_id"`
	Branch        string                  `json:"branch" gorm:"column:branch"` // 分支名称
	CreatedBy     uint                    `json:"created_by" gorm:"column:created_by"`
	Remark        string                  `json:"remark" gorm:"column:remark"`
	LastRunAt     *utils.JsonTime         `json:"last_run_at" gorm:"column:last_run_at"`         // 最近一次运行时间
	LastRunResult pipeline.BuildingResult `json:"last_run_result" gorm:"column:last_run_result"` // 最近一次运行结果; 1：成功；2：取消；3：失败；0：从未执行
}

func (m *Pipeline) TableName() string {
	return "pipeline"
}

// NewPlan 新建构建计划
func NewPlan(projectId uint, currentUserId uint, model *pipeline.PlanModel) (err error) {
	steps := make(steps, 0)
	sort := 0
	err = steps.qaStep(model, &sort)
	if err != nil {
		return err
	}
	err = steps.artifactStep(model, &sort)
	if err != nil {
		return err
	}

	plan := Pipeline{
		ProjectID:     projectId,
		Name:          model.Name,
		Env:           model.Env,
		SourceCodeID:  model.SourceCodeID,
		Branch:        model.Branch,
		CreatedBy:     currentUserId,
		Remark:        model.Remark,
		LastRunResult: 0,
		PipelineSteps: steps,
	}

	tx := conf.GetDbClient()

	err = tx.Omit("updated_at").Model(&Pipeline{}).Create(&plan).Error

	return
}

func QueryPlan(projectId uint) ([]Pipeline, error) {
	db := conf.GetDbClient()

	var plans []Pipeline

	tx := db.Model(&Pipeline{})

	tx = tx.Preload(clause.Associations)
	tx = tx.Where("project_id = ?", projectId)
	err := tx.Find(&plans).Error

	return returnWithError(plans, err)
}

func DeletePlan(projectId uint, planId uint) error {
	db := conf.GetDbClient()

	tx := db.Model(&Pipeline{})

	tx = tx.Preload(clause.Associations)
	tx = tx.Where("project_id = ?", projectId)
	err := tx.Delete(&Pipeline{Model: Model{ID: planId}}).Error

	if err == nil {
		tx = tx.Session(&gorm.Session{NewDB: true})
		tx.Model(&PipelineSteps{}).Where("ci_plan_id = ?", planId).Delete(&PipelineSteps{})
	}
	return err
}
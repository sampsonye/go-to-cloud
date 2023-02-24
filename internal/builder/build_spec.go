package builder

import (
	"encoding/json"
	"fmt"
	lang2 "go-to-cloud/internal/builder/lang"
	"go-to-cloud/internal/models/pipeline"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
	"strconv"
)

func ResetIdle(node *repositories.BuilderNode) {
	idleNodes.Delete(strconv.Itoa(int(node.ID)))
}

// BuildPodSpec 创建构建模板 k8s pod spec
// buildId: pipeline_history.ID
func BuildPodSpec(buildId uint, node *repositories.BuilderNode, plan *repositories.Pipeline) *kube.PodSpecConfig {
	var lang lang2.Tpl
	switch plan.Env {
	case lang2.DotNet3, lang2.DotNet5, lang2.DotNet6, lang2.DotNet7:
		lang = &lang2.DotNet{}
	case lang2.Go120, lang2.Go116, lang2.Go119, lang2.Go118, lang2.Go117:
		lang = &lang2.Golang{}
	}
	padLeftBuildIdStr := utils.DockerImageTagBuild(buildId)
	return &kube.PodSpecConfig{
		LabelFlag:    NodeSelectorLabel,
		LabelBuildId: BuildIdSelectorLabel,
		BuildId:      buildId,
		Namespace:    node.K8sWorkerSpace,
		TaskName:     plan.Name + "-" + padLeftBuildIdStr,
		SourceCode:   plan.SourceCode.GitUrl,
		Sdk:          lang.Sdk(plan.Env),
		Steps: func() []kube.Step {
			kvp := lang.Steps(plan.Env, plan.PipelineSteps)
			steps := make([]kube.Step, len(kvp))
			i := 0
			for t, cmd := range kvp {
				steps[i] = kube.Step{
					Command:     cmd,
					CommandText: (&t).GetTypeName(),
					CommandType: t,
				}
				switch t {
				case pipeline.Image:
					steps[i].Command = ""
					var model pipeline.ArtifactScript
					if err := json.Unmarshal([]byte(cmd), &model); err == nil {
						steps[i].Dockerfile = model.Dockerfile
						steps[i].ArtifactName = fmt.Sprintf("%s:v%s", plan.ArtifactName, padLeftBuildIdStr)
						steps[i].Registry.Url = model.Registry
						steps[i].Registry.User = model.Account
						steps[i].Registry.Password = model.Password
						steps[i].Registry.Security = model.IsSecurity
					}
				}

				i++
			}
			return steps
		}(),
	}
}
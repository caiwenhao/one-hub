package task

import (
    "errors"
    "one-api/common/config"
    "one-api/model"
    "one-api/relay/task/base"
    "one-api/relay/task/kling"
    "one-api/relay/task/minimax"
    "one-api/relay/task/sora"
    "one-api/relay/task/suno"
    "one-api/relay/task/vidu"
    "one-api/relay/task/volcark"

	"github.com/gin-gonic/gin"
)

func GetTaskAdaptor(relayType int, c *gin.Context) (base.TaskInterface, error) {
	switch relayType {
	case config.RelayModeSuno:
		return &suno.SunoTask{
			TaskBase: getTaskBase(c, model.TaskPlatformSuno),
		}, nil
	case config.RelayModeKling:
		return &kling.KlingTask{
			TaskBase: getTaskBase(c, model.TaskPlatformKling),
		}, nil
	case config.RelayModeVidu:
		return &vidu.ViduTask{
			TaskBase: getTaskBase(c, model.TaskPlatformVidu),
		}, nil
	case config.RelayModeVolcArkVideo:
		return &volcark.VolcArkTask{
			TaskBase: getTaskBase(c, model.TaskPlatformVolcArk),
		}, nil
    case config.RelayModeMiniMaxVideo:
        return &minimax.MiniMaxTask{TaskBase: getTaskBase(c, model.TaskPlatformMiniMax)}, nil
	default:
		return nil, errors.New("adaptor not found")
	}
}

func GetTaskAdaptorByPlatform(platform string) (base.TaskInterface, error) {
	relayType := config.RelayModeUnknown

	switch platform {
	case model.TaskPlatformSuno:
		relayType = config.RelayModeSuno
	case model.TaskPlatformKling:
		relayType = config.RelayModeKling
	case model.TaskPlatformVidu:
		relayType = config.RelayModeVidu
	case model.TaskPlatformVolcArk:
		relayType = config.RelayModeVolcArkVideo
	case model.TaskPlatformMiniMax:
		relayType = config.RelayModeMiniMaxVideo
	case model.TaskPlatformSora:
		// 使用 OpenAI 视频路由，单独实现 Sora 任务适配器
		return &sora.SoraTask{TaskBase: getTaskBase(nil, model.TaskPlatformSora)}, nil
	}

	return GetTaskAdaptor(relayType, nil)
}

func getTaskBase(c *gin.Context, platform string) base.TaskBase {
	return base.TaskBase{
		Platform: platform,
		C:        c,
	}
}

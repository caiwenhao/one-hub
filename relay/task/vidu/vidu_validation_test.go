package vidu

import (
	"encoding/json"
	"testing"

	"one-api/model"
	viduProvider "one-api/providers/vidu"
)

func TestActionValidateImg2VideoSingleImage(t *testing.T) {
	task := &ViduTask{
		Action: viduProvider.ViduActionImg2Video,
		Request: &viduProvider.ViduTaskRequest{
			Images: []string{"a", "b"},
		},
	}

	if err := task.actionValidate(); err == nil {
		t.Fatalf("expected error when more than one image provided")
	}
}

func TestActionValidateImg2VideoDefaults(t *testing.T) {
	task := &ViduTask{
		Action: viduProvider.ViduActionImg2Video,
		Request: &viduProvider.ViduTaskRequest{
			Images: []string{"a"},
		},
	}

	if err := task.actionValidate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if task.OriginalModel != "vidu-img2video-vidu1.5-4s-360p" {
		t.Fatalf("unexpected original model: %s", task.OriginalModel)
	}
}

func TestActionValidateReference2VideoValidations(t *testing.T) {
	baseReq := func(images int, prompt string) *viduProvider.ViduTaskRequest {
		imgs := make([]string, images)
		for i := range imgs {
			imgs[i] = "img"
		}
		return &viduProvider.ViduTaskRequest{
			Images: imgs,
			Prompt: prompt,
		}
	}

	cases := []struct {
		name      string
		req       *viduProvider.ViduTaskRequest
		wantError bool
	}{
		{"no images", baseReq(0, "prompt"), true},
		{"too many images", baseReq(8, "prompt"), true},
		{"blank prompt", baseReq(3, " \t\n"), true},
		{"valid", baseReq(2, "prompt"), false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			task := &ViduTask{
				Action:  viduProvider.ViduActionReference2Video,
				Request: tc.req,
			}
			err := task.actionValidate()
			if tc.wantError && err == nil {
				t.Fatalf("expected error but got none")
			}
			if !tc.wantError && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestActionValidateText2VideoPromptRequired(t *testing.T) {
	task := &ViduTask{
		Action:  viduProvider.ViduActionText2Video,
		Request: &viduProvider.ViduTaskRequest{Prompt: ""},
	}

	if err := task.actionValidate(); err == nil {
		t.Fatalf("expected error when prompt is empty")
	}
}

func TestTaskModel2Dto(t *testing.T) {
	credits := 4
	legacy := viduProvider.ViduResponse{
		TaskID:  "legacy",
		State:   "SUCCESS",
		Credits: &credits,
	}
	legacyData, _ := json.Marshal(legacy)

	newResp := viduProvider.ViduQueryResponse{
		ID:    "new",
		State: viduProvider.ViduStatusProcessing,
	}
	newData, _ := json.Marshal(newResp)

	tests := []struct {
		name string
		task *model.Task
		want string
	}{
		{
			name: "new schema",
			task: &model.Task{TaskID: "new", Status: model.TaskStatusInProgress, Data: newData},
			want: viduProvider.ViduStatusProcessing,
		},
		{
			name: "legacy schema",
			task: &model.Task{TaskID: "legacy", Status: model.TaskStatusSuccess, Data: legacyData},
			want: viduProvider.ViduStatusSuccess,
		},
		{
			name: "fallback",
			task: &model.Task{TaskID: "fallback", Status: model.TaskStatusFailure},
			want: viduProvider.ViduStatusFailed,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dto := TaskModel2Dto(tc.task)
			if dto == nil {
				t.Fatalf("expected dto")
			}
			if dto.ID == "" {
				t.Fatalf("dto.ID should not be empty")
			}
			if dto.State != tc.want {
				t.Fatalf("expected state %s, got %s", tc.want, dto.State)
			}
		})
	}
}

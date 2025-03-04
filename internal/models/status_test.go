package models

import (
	"testing"
	"time"

	"github.com/yohamta/dagu/internal/config"
	"github.com/yohamta/dagu/internal/scheduler"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPid(t *testing.T) {
	if PidNotRunning.IsRunning() {
		t.Error()
	}
	var pid Pid = Pid(-1)
	require.Equal(t, "", pid.String())

	pid = Pid(12345)
	require.Equal(t, "12345", pid.String())
}

func TestStatusSerialization(t *testing.T) {
	start, end := time.Now(), time.Now().Add(time.Second*1)
	cfg := &config.Config{
		ConfigPath:  "",
		Name:        "",
		Description: "",
		Env:         []string{},
		LogDir:      "",
		HandlerOn:   config.HandlerOn{},
		Steps: []*config.Step{
			{
				Name: "1", Description: "", Variables: []string{},
				Dir: "dir", Command: "echo 1", Args: []string{},
				Depends: []string{}, ContinueOn: config.ContinueOn{},
				RetryPolicy: &config.RetryPolicy{}, MailOnError: false,
				RepeatPolicy: config.RepeatPolicy{}, Preconditions: []*config.Condition{},
			},
		},
		MailOn:            config.MailOn{},
		ErrorMail:         &config.MailConfig{},
		InfoMail:          &config.MailConfig{},
		Smtp:              &config.SmtpConfig{},
		Delay:             0,
		HistRetentionDays: 0,
		Preconditions:     []*config.Condition{},
		MaxActiveRuns:     0,
		Params:            []string{},
		DefaultParams:     "",
	}
	st := NewStatus(cfg, nil, scheduler.SchedulerStatus_Success, 10000, &start, &end)

	js, err := st.ToJson()
	require.NoError(t, err)

	st_, err := StatusFromJson(string(js))
	require.NoError(t, err)

	assert.Equal(t, st.Name, st_.Name)
	require.Equal(t, 1, len(st_.Nodes))
	assert.Equal(t, cfg.Steps[0].Name, st_.Nodes[0].Name)
}

func TestCorrectRunningStatus(t *testing.T) {
	cfg := &config.Config{Name: "test"}
	status := NewStatus(cfg, nil, scheduler.SchedulerStatus_Running,
		10000, nil, nil)
	status.CorrectRunningStatus()
	require.Equal(t, scheduler.SchedulerStatus_Error, status.Status)
}

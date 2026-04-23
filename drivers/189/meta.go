package _189

import (
	"github.com/OpenListTeam/OpenList/v4/internal/driver"
	"github.com/OpenListTeam/OpenList/v4/internal/op"
)

type Addition struct {
	Username string `json:"username" required:"true"`
	Password string `json:"password" required:"true"`
	Cookie   string `json:"cookie" help:"Fill in the cookie if need captcha"`
	DownloadThread string `json:"download_thread" default:"32" help:"1<=thread<=128"`
	DownloadPartSize string `json:"download_part_size" default:"16" help:"Part size in MB, 1<=size<=64"`
	driver.RootID
}

var config = driver.Config{
	Name:        "189Cloud",
	LocalSort:   true,
	DefaultRoot: "-11",
	Alert:       `info|You can try to use 189PC driver if this driver does not work.`,
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &Cloud189{}
	})
}

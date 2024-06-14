package dialog

import (
	"context"

	"gtools-wails/backend/utils"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type BackendDialog struct {
	Ctx context.Context
}

func (d *BackendDialog) Startup(ctx context.Context) {
	d.Ctx = ctx
}

// displayName : Image Files (*.jpg, *.png)
// pattern : *.jpg;*.png
func (d *BackendDialog) OneFileDialog(displayName, pattern string) string {
	res, err := runtime.OpenFileDialog(d.Ctx, runtime.OpenDialogOptions{
		DefaultDirectory: "",
		DefaultFilename:  "",
		Title:            "选择一个文件",
		Filters: []runtime.FileFilter{
			{
				DisplayName: displayName,
				Pattern:     pattern,
			},
		},
		ShowHiddenFiles:            false,
		CanCreateDirectories:       false,
		ResolvesAliases:            false,
		TreatPackagesAsDirectories: false,
	})
	return utils.Resp(res, err)
}

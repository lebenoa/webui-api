package extra

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Meonako/webui-api"
	"github.com/Meonako/webui-api/utils"
)

// Convenience function to build []ImageData from base64-encoded image
func BuildBatch(imageList ...string) (res []api.ImageData) {
	for _, image := range imageList {
		res = append(res, api.ImageData{Data: "data:image/png;base64," + image})
	}

	return
}

// Convenience function to build []ImageData from files
func BuildBatchFromFiles(files ...string) (res []api.ImageData, err error) {
	for _, path := range files {
		encoded, err := utils.Base64FromFile(path)
		if err != nil {
			return []api.ImageData{}, err
		}
		res = append(res, api.ImageData{Data: "data:image/png;base64," + encoded})
	}

	return
}

// Convenience function to build []ImageData from files and IGNORE ANY ERRORS that might occur.
//
// This may be helpful when don't want your app to crash when file doesn't exists.
func BuildBatchFromFilesIgnore(files ...string) (res []api.ImageData) {
	for _, path := range files {
		encoded, _ := utils.Base64FromFile(path)
		res = append(res, api.ImageData{Data: "data:image/png;base64," + encoded})
	}

	return
}

// EXP: Convenience function to build []ImageData from directory
func BuildBatchFromDir(pathToDir string) (res []api.ImageData, err error) {
	files, readErr := os.ReadDir(pathToDir)
	if readErr != nil {
		return []api.ImageData{}, readErr
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		ext := filepath.Ext(fileName)
		if ext == "" && strings.ToLower(ext) != "png" && strings.ToLower(ext) != "jpg" && strings.ToLower(ext) != "jpeg" {
			continue
		}

		encoded, encErr := utils.Base64FromFile(filepath.Join(pathToDir, fileName))
		if encErr != nil {
			return []api.ImageData{}, encErr
		}
		res = append(res, api.ImageData{Data: "data:image/png;base64," + encoded})
	}

	return
}

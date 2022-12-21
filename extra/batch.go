package extra

import (
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

// TODO: BuildBatchFromDir // Convenience function to build []ImageData from directory

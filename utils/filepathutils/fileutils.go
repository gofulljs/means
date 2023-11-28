package filepathutils

import (
	"golang.org/x/xerrors"
	"os"
	"path/filepath"
)

/**
 * @description: 从完整文件名称来检测路径，防止路径不存在
 * @param {string} fullName 完整文件名称
 * @return {*}
 */
func CheckPathFromFullName(fullName string) (err error) {
	dir, _ := filepath.Split(fullName)
	return CheckPath(dir)
}

/**
 * @description: 获取文件是否存在
 * @param {string} filepath
 * @return {*}
 */
func GetFileExist(filepath string) (bool, error) {
	fs, exist, err := GetPathInfo(filepath)
	if err != nil {
		return false, xerrors.Errorf("%w", err)
	}
	if exist {
		if fs.IsDir() {
			return false, xerrors.Errorf("filepath %v is a directory", filepath)
		}
	}

	return exist, nil
}

/**
 * @description: 删除文件
 * @param {string} fileFullPath
 * @return {*}
 */
func DeleteFileIfExist(fileFullPath string) error {
	isExist, err := GetFileExist(fileFullPath)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}
	if isExist {
		err = os.Remove(fileFullPath)
		if err != nil {
			return xerrors.Errorf("%w", err)
		}
	}

	return nil
}

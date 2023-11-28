package filepathutils

import (
	"io/fs"
	"os"

	"golang.org/x/xerrors"
)

/**
 * @description: 获取路径信息
 * @param {string} path
 * @return {*}
 */
func GetPathInfo(path string) (fs fs.FileInfo, exist bool, err error) {
	fs, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, false, nil
		}
		return nil, false, xerrors.Errorf("%w", err)
	}

	return fs, true, nil
}

/**
 * @description: 获取路径是否存在
 * @param {string} path
 * @return {*}
 */
func GetPathExist(path string) (bool, error) {
	fs, exist, err := GetPathInfo(path)
	if err != nil {
		return false, xerrors.Errorf("%w", err)
	}
	if exist {
		if !fs.IsDir() {
			return false, xerrors.Errorf("path %v is not a directory", path)
		}
	}

	return exist, nil
}

/**
 * @description: 检测路径，没有就创建
 * @param {string} path
 * @return {*}
 */
func CheckPath(path string) (err error) {
	fs, exist, err := GetPathInfo(path)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	if exist {
		if !fs.IsDir() {
			return xerrors.Errorf("path %v is not a directory", path)
		}
	} else {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return xerrors.Errorf("%w", err)
		}
	}

	return
}

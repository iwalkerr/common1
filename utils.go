package common

// 获取ssl文件路径
func AuthFilePath(key, pem string) (string, string) {
	path := env(certPath)
	return path + key, path + pem
}

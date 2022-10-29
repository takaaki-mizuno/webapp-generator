package services

// MockGitService ...
type MockGitService struct {
}

// DownloadBoilerplate ...
func (service *MockGitService) DownloadBoilerplate(path string, projectName string, templateName string) error {
	return nil
}

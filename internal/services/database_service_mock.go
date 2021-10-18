package services

// MockDatabaseService ...
type MockDatabaseService struct {
}

// GenerateDatabase ...
func (service *MockDatabaseService) GenerateDatabase(path string, databaseDefinitionPath string, language string, projectName string) error {
	return nil
}

// GenerateAdminAPI ...
func (service *MockDatabaseService) GenerateAdminAPI(path string, databaseDefinitionPath string, language string, projectName string) error {
	return nil
}

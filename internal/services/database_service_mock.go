package services

// MockDatabaseService ...
type MockDatabaseService struct {
}

func (service *MockDatabaseService) GenerateDatabase(path string, databaseDefinitionPath string, language string, projectName string) error {
	return nil
}
func (service *MockDatabaseService) GenerateAdminAPI(path string, databaseDefinitionPath string, language string, projectName string) error {
	return nil
}

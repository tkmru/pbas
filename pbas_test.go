package pbas

import "testing"

func TestAddAndFindConfig(t *testing.T) {
  configName := "README.md"
  SetConfigName(configName)
  if p.configName != configName {
    t.Error("pbas.configName doesn't match SetConfigName arg.")
    t.Log("pbas.configName:", p.configName)
  }

  configPaths := []string{"./", "/etc/pbas"}
  for _, path := range configPaths {
    AddConfigPath(path)
  }

  if len(p.configPaths) != len(configPaths) {
    t.Error("Length of pbas.configPaths isn't correct.")
    t.Log("pbas.configPaths:", p.configPaths)
  }

  configPath := FindConfigPath()

  if configPath == "" {
    t.Error("The config path doesn't found.")
  }
}

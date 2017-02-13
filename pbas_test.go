package pbas

import "testing"

func TestAddAndFindConfig(t *testing.T) {
  configName := "README.md"
  SetConfigName(configName)
  if p.configName != configName {
    t.Error("pbas.configName doesn't match SetConfigName arg.")
    t.Log("pbas.configName:", p.configName)
  }

  AddConfigPath("./")
  AddConfigPath("/etc/pbas")

  if len(p.configPaths) != 2 {
    t.Error("Length of pbas.configPaths isn't correct.")
    t.Log("pbas.configPaths:", p.configPaths)
  }

  configPath := FindConfigPath()

  if configPath == "" {
    t.Error("The config path doesn't found.")
  }
}

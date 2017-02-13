package pbas

import (
	"io/ioutil"
	"path"
)

type Pbas struct {
	configName  string
	configPaths []string
}

var p *Pbas

func init() {
  p = new(Pbas)
  p.configName = ""
  p.configPaths = []string{}
}

func SetConfigName(fileName string) { p.SetConfigName(fileName) }
func (p *Pbas) SetConfigName(fileName string) {
	if fileName != "" {
		p.configName = fileName
	}
}

func AddConfigPath(filePath string) { p.AddConfigPath(filePath) }
func (p *Pbas) AddConfigPath(filePath string) {
	if filePath != "" {
		p.configPaths = append(p.configPaths, filePath)
	}
}

func FindConfigPath() string { return p.FindConfigPath() }
func (p *Pbas) FindConfigPath() string {
	for _, configPath := range p.configPaths {
		files, _ := ioutil.ReadDir(configPath)
		for _, f := range files {
			if f.Name() == p.configName {
				return path.Join(configPath, p.configName)
			}
		}
	}
	return ""
}

package file

import (
	"BilibiliBot/util/log"
	"io/ioutil"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

const logTag = "[file_handle]"

var (
	// ConfigTypeJSON ...
	ConfigTypeJSON ConfigType = "json"
	// ConfigTypeToml ...
	ConfigTypeToml ConfigType = "toml"
	// ConfigTypeYaml ...
	ConfigTypeYaml ConfigType = "yaml"
)

type FileHandle struct {
	path        string
	enableWatch bool
	changed     chan struct{}
}

func init() {
}

func extParser(cfgPath string) ConfigType {
	ext := filepath.Ext(cfgPath)
	switch ext {
	case ".json":
		return ConfigTypeJSON
	case ".toml":
		return ConfigTypeToml
	case ".yaml":
		return ConfigTypeYaml
	default:
		log.Panic("data source: invalid configuration type")
	}
	return ""
}

func (self *FileHandle) Parse(path string, watch bool) ConfigType {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		log.Panic("new datasource", err)
	}
	self.path = absolutePath
	self.enableWatch = watch

	if watch {
		self.changed = make(chan struct{}, 1)
		log.Info("go self.watch()", absolutePath)
		go self.watch()
	}

	return extParser(path)
}

func (self *FileHandle) ReadConfig() (content []byte, err error) {
	return ioutil.ReadFile(self.path)
}

func (self *FileHandle) Close() error {
	close(self.changed)
	return nil
}

func (self *FileHandle) IsConfigChanged() <-chan struct{} {
	return self.changed
}

// Watch file and automate update.
func (self *FileHandle) watch() {
	w, err := fsnotify.NewWatcher()
	log.Info("go watch 0")

	if err != nil {
		log.Fatal("new file watcher", err)
	}
	log.Info("go watch 1")
	defer w.Close()

	configFile := filepath.Clean(self.path)
	log.Info("go watch 2")

	realConfigFile, _ := filepath.EvalSymlinks(self.path)

	log.Info("read watch",
		"configFile", configFile,
		"realConfigFile", realConfigFile,
		"fppath", self.path,
	)

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-w.Events:
				currentConfigFile, _ := filepath.EvalSymlinks(self.path)

				log.Info("read watch event",
					"event", filepath.Clean(event.Name),
					"path", filepath.Clean(self.path),
					"currentConfigFile", currentConfigFile,
					"realConfigFile", realConfigFile,
				)
				// we only care about the config file with the following cases:
				// 1 - if the config file was modified or created
				// 2 - if the real path to the config file changed (eg: k8s ConfigMap replacement)
				const writeOrCreateMask = fsnotify.Write | fsnotify.Create
				if (filepath.Clean(event.Name) == configFile &&
					event.Op&writeOrCreateMask != 0) ||
					(currentConfigFile != "" && currentConfigFile != realConfigFile) {

					realConfigFile = currentConfigFile
					log.Info("modified file", event.Name, realConfigFile)

					select {
					case self.changed <- struct{}{}:
					default:

					}
				}
			case err := <-w.Errors:
				log.Error("read watch error", err)
			}
		}
	}()
	err = w.Add(self.path)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

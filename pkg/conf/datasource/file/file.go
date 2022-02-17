package file

import (
	"errors"
	"fmt"
	"github.com/whileW/core-go/pkg/util/xcodec"
	"github.com/whileW/core-go/pkg/util/xfile"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// DataSourceFile defines file scheme
const DataSourceFile = "file"

type fileDataSource struct {
	path        string
	dir         string
	codec 		string
	//enableWatch bool
	//changed     chan struct{}
}

// NewDataSource returns new fileDataSource.
func NewDataSource() (*fileDataSource,error) {
	absolutePath, err := filepath.Abs(config_file_path)
	if err != nil {
		return nil,errors.New(fmt.Sprintf("get file config absolute path error. path:%s, err: %v",config_file_path,err))
	}
	fileType := strings.Split(config_file_path,".")
	var content_codec string
	switch fileType[len(fileType)-1] {
	case "yaml":
		content_codec = xcodec.XcodecYaml
	case "json":
		content_codec = xcodec.XcodecJson
	default:
		return nil,errors.New(fmt.Sprintf("no codec adaptate this file type [%s]",config_file_path))
	}

	dir := xfile.CheckAndGetParentDir(absolutePath)
	ds := &fileDataSource{path: absolutePath, dir: dir,codec:content_codec}
	//ds.changed = make(chan struct{}, 1)
	//go ds.watch()
	return ds,nil
}

func (fp *fileDataSource) ReadConfig() (map[string]interface{},error) {
	content,err := ioutil.ReadFile(fp.path)
	if err != nil {
		return nil,errors.New(fmt.Sprintf("read file by absolute path error. path:%s, absolute_path:%s, err: %v",config_file_path,fp.path,err))
	}
	conf := map[string]interface{}{}
	if err := xcodec.Decode(fp.codec,content,&conf);err != nil{
		return nil,errors.New(fmt.Sprintf("decode file config data failed: %v，codec: %s", err,fp.codec))
	}
	return conf,nil
}

//// Close ...
//func (fp *fileDataSource) Close() error {
//	close(fp.changed)
//	return nil
//}
//
//// IsConfigChanged ...
//func (fp *fileDataSource) IsConfigChanged() <-chan struct{} {
//	return fp.changed
//}
//
//// Watch file and automate update.
//func (fp *fileDataSource) watch() {
//	w, err := fsnotify.NewWatcher()
//	if err != nil {
//		fmt.Println("new file watcher error:"+err.Error())
//		return
//	}
//
//	defer w.Close()
//	done := make(chan bool)
//	go func() {
//		for {
//			select {
//			case event := <-w.Events:
//				fmt.Println(fmt.Sprintf("read watch event,event:%s,path:%s",filepath.Clean(event.Name),filepath.Clean(fp.path)))
//				// we only care about the config file with the following cases:
//				// 1 - if the config file was modified or created
//				// 2 - if the real path to the config file changed
//				const writeOrCreateMask = fsnotify.Write | fsnotify.Create
//				if event.Op&writeOrCreateMask != 0 && filepath.Clean(event.Name) == filepath.Clean(fp.path) {
//					log.Println("modified file: ", event.Name)
//					select {
//					case fp.changed <- struct{}{}:
//					default:
//					}
//				}
//			case err := <-w.Errors:
//				// log.Println("error: ", err)
//				fmt.Println("read watch error:"+err.Error())
//			}
//		}
//		done<-true
//	}()
//
//	err = w.Add(fp.dir)
//	if err != nil {
//		fmt.Println(err)
//	}
//	<-done
//}
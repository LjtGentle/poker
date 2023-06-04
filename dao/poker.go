package dao

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"poker/model"
	"runtime"
)

// ReadFile 把数据从文件中读取出来 分别放在切片中返回
func ReadFile(filename string) (alices, bobs []string, results []model.Result) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var matches model.Match
	err = json.Unmarshal(buf, &matches)
	if err != nil {
		panic(err)
	}

	alices = make([]string, len(matches.Matches))
	bobs = make([]string, len(matches.Matches))
	results = make([]model.Result, len(matches.Matches))

	for k, v := range matches.Matches {
		alices[k] = v.Alice
		bobs[k] = v.Bob
		results[k] = v.Result
	}
	return
}

// GetCurrentAbPathByCaller 得到项目的路径
func GetCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
		abPath = path.Dir(abPath)
	}
	return abPath
}

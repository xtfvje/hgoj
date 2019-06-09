package tools

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
	"io/ioutil"
	"os"
	"time"
)

func StartCron() {
	tk := toolbox.NewTask("clearDownData", "0 0 */1 * * *", func() error { clearDownData(); return nil })
	err := tk.Run()
	if err != nil {
		fmt.Println(err)
	}
	toolbox.AddTask("clearDownData", tk)
	toolbox.StartTask()
	//time.Sleep(6 * time.Second)
	//toolbox.StopTask()
}


func clearDownData() {
	downDir := "./static/down"

	dir_list, e := ioutil.ReadDir(downDir)
	if e != nil {
		fmt.Println("read dir error")
		return
	}
	for _, v := range dir_list {
		finfo, _ := os.Stat(downDir+"/"+v.Name())
		var fctime time.Time
		fctime = finfo.ModTime()
		t := time.Now().Sub(fctime).Minutes()
		logs.Info(t)
		if t > 10 {
			err := os.Remove(downDir+"/"+v.Name())
			if err != nil {
				logs.Error(err)
			}
		}
		logs.Info("delete file ",downDir+"/"+v.Name())
	}
}
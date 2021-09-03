package conf

import (
	"fmt"
	"testing"
)

func TestInitConfg(t *testing.T) {
	conf := InitConfg()
	fmt.Println(conf)
}

func TestSettings_Getd(t *testing.T) {
	v := GetConf().Setting.GetChildd("test1").GetChildd("test2").GetStringd("test3","value")
	fmt.Println(v)
}
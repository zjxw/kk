package controller

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func DelUser(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id := vars.Get("Id")
	if id == "" {
		rs := strings.NewReader("Id 不能为空")
		_,_ = io.Copy(w,rs)
		return
	}
	fid, _ := strconv.Atoi(id)
	err := Udb.Del(fid)
	if err != nil {
		rs := strings.NewReader(fmt.Sprintf("%s",err))
		_,_ = io.Copy(w,rs)
		return
	}
	err = Udb.Sync()
	if err == nil {
		http.Redirect(w,r,"/",302)
	}else {
		rs := strings.NewReader(fmt.Sprintf("%s",err))
		_,_ = io.Copy(w,rs)
		return
	}

}

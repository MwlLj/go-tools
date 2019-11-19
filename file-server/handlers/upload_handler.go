package handlers

import (
	"fmt"
	"github.com/MwlLj/gotools/ospath"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	html string = `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Insert title here</title>
</head>
<body>

    <form action="/file/upload" id="upfile"  method="post" enctype="multipart/form-data">
        目标路径 (默认为程序启动目录):<br>
        <input type="text" id="dest" name="dest" />
        <br>
        <br>
        <input type="file" id="file" name="file" value="浏览图片" />
        <br>
        <br>
        <input type="submit" value="提交"/>
    </form>

</body>
</html>
	`
)

const (
	FormDataFile string = "file"
	FormDataDest string = "dest"
)

func UploadHandler(w http.ResponseWriter, r *http.Request, userName *string, userPwd *string) {
	authHandler(w, r, userName, userPwd, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			io.WriteString(w, html)
		} else {
			var dest string = "."
			{
				err := r.ParseMultipartForm(32 << 20)
				if err != nil {
					io.WriteString(w, "parse error")
					return
				}
				multi := r.MultipartForm
				d, ok := multi.Value[FormDataDest]
				if !ok || len(d) == 0 {
					io.WriteString(w, "dest is null")
					return
				}
				de := d[0]
				if de != "" {
					dest = de
				}
			}
			file, head, err := r.FormFile(FormDataFile)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()
			path := strings.Join([]string{dest, head.Filename}, "/")
			ospath.CreateDirsIfNotExists(&dest)
			fW, err := os.Create(path)
			if err != nil {
				fmt.Println(path, err)
				io.WriteString(w, "file create failed")
				return
			}
			defer fW.Close()
			_, err = io.Copy(fW, file)
			if err != nil {
				io.WriteString(w, "file save failed")
				return
			}
			io.WriteString(w, "upload success")
		}
	})
}

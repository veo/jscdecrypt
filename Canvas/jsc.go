package Canvas

import (
	"bytes"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	dialog2 "github.com/sqweek/dialog"
	"github.com/xxtea/xxtea-go/xxtea"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

var Key = widget.NewEntry()
var Jscpath = widget.NewEntry()
var Outputpath = widget.NewEntry()
var Cmdout = widget.NewLabel("")
var Fileslist = widget.NewLabel("")
var Iszip = widget.NewSelect([]string{"yes", "no"}, nil)

func Decrypt(jscpath string) {
	var filelist string
	files := GetFileList(jscpath)
	if Jscpath.Text == "" || Outputpath.Text == "" {
		Cmdout.SetText("error: please input Jscpath or Outputpath")
	}
	for _, file := range files {
		filelist = filelist + "\n" + file
	}
	Fileslist.SetText("files: \n" + filelist)
	for _, file := range files {
		newpath := Outputpath.Text + strings.Replace(file, jscpath, "", -1)
		b, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Print(err)
		}
		decrypt_data := xxtea.Decrypt(b, []byte(Key.Text))
		err = os.MkdirAll(path.Dir(newpath), os.ModePerm)
		if err != nil {
			Cmdout.SetText("Decrypt error: unable to creat dir")
			return
		}
		if decrypt_data != nil {
			if Iszip.Selected == "yes" {

				err = ioutil.WriteFile(newpath, decrypt_data, 0666)
				if err != nil {
					Cmdout.SetText("Decrypt error: unable to write new file")
					return
				}
				_ = Unzip(newpath, Outputpath.Text)
				err = os.Rename(Outputpath.Text+"/encrypt.js", strings.Replace(newpath, "jsc", "js", -1))
				_ = os.Remove(newpath)
				if err != nil {
					Cmdout.SetText("Decrypt error: unable to unzip")
					return
				} else {
					Cmdout.SetText("Decrypt sucess")
				}
			} else {
				err = ioutil.WriteFile(strings.Replace(newpath, "jsc", "js", -1), decrypt_data, 0666)
				if err != nil {
					Cmdout.SetText("Decrypt error: unable to write new file,may be zip")
					return
				}
			}
		} else {
			Cmdout.SetText("key is wrong")
			return
		}
	}

}

//func Getkey()  {
//	widget.NewButton("get key", func() {
//		fileOpen := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
//			if err != nil {
//				println(err)
//			}
//			if closer != nil {
//				defer closer.Close()
//				bc := bufio.NewScanner(closer)
//				//var n = 0
//				for bc.Scan() {
//					//n = n+1
//					//if n < 3 {
//					//	fmt.Println(bc.Text())
//					//	encodedStr := hex.EncodeToString(bc.Bytes())
//					//	fmt.Println(encodedStr)
//					//	test, _ := hex.DecodeString(encodedStr)
//					//	fmt.Println(string(test))
//					//}
//					encodedStr := hex.EncodeToString(bc.Bytes())
//					reg := regexp.MustCompile("636f6")
//					//fmt.Printf("%s\n", reg.Find(bc.Bytes()))
//					if reg.FindAllString(encodedStr,-1) != nil {
//						fmt.Println("yes")
//						fmt.Printf("%s\n", reg.FindAllString(encodedStr,-1))
//					}
//					//println(bc.Text())
//				}
//			}
//		}, w)
//		fileOpen.SetDismissText("open fail")
//		fileOpen.Show()
//	})
//}

func JscfileCanvas() fyne.CanvasObject {
	Jscpath.SetPlaceHolder("jsc path")
	jscPathButton := widget.NewButton("select jsc path", func() {
		path, err := dialog2.Directory().Title("select jsc path").Browse()
		if err != nil {
			fmt.Println(err)
		} else {
			Jscpath.SetText(path)
		}
	})
	jscinput := fyne.NewContainerWithLayout(
		layout.NewGridLayout(2),
		Jscpath, jscPathButton,
	)
	return jscinput
}
func OutfileCanvas() fyne.CanvasObject {
	Outputpath.SetPlaceHolder("output path")
	outputButton := widget.NewButton("select output path", func() {
		path, err := dialog2.Directory().Title("select output path").Browse()
		if err != nil {
			fmt.Println(err)
		} else {
			Outputpath.SetText(path)
		}
	})
	output := fyne.NewContainerWithLayout(
		layout.NewGridLayout(2),
		Outputpath, outputButton,
	)
	return output
}

func CmdtestCanvas() fyne.CanvasObject {
	test := widget.NewEntry()
	test.SetPlaceHolder("cmdshell")
	cmdButton := widget.NewButton("select output path", func() {
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		cmd := exec.Command("bash", "-c", test.Text)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		} else {
			Cmdout.SetText(stdout.String())
		}
	})
	cmdtset := fyne.NewContainerWithLayout(
		layout.NewGridLayout(2),
		test, cmdButton,
	)
	return cmdtset
}

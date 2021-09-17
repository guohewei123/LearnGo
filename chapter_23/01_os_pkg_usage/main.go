package main

import (
	"fmt"
	"io"
	"os"
)

/*
1. 使用OpenFile打开文件进行读操作
*/

func OpenFileReadUsage() {
	f, err := os.OpenFile("./learnGo/chapter_23/01_os_pkg_usage/test.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	b := make([]byte, 10)
	for i := 0; i < 10; i++ {
		n, err := f.Read(b)
		if err != nil {
			if err == io.EOF{
				fmt.Println("Complete Read: ", err)
				return
			}
			fmt.Println(err)
			return
		}
		fmt.Printf("%s    %d\n", string(b), n)
	}
}

/*
1. 使用OpenFile打开文件进行写操作
*/

func OpenFileWriteUsage() {
	f, err := os.OpenFile("./learnGo/chapter_23/01_os_pkg_usage/test.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	//seek, err := f.Seek(50, io.SeekCurrent)
	//seek, err := f.Seek(50, io.SeekEnd)
	seek, err := f.Seek(-50, io.SeekEnd)
	if err != nil {
		panic(err)
	}
	fmt.Println("Current seek: ", seek)
	n, err := f.Write([]byte("hello world"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Write byte num: ", n)
}

/*
1. 使用bufio.NewReader打开文件进行写操作
*/

func BufioNewReaderUsage() {
	f, err := os.OpenFile("./learnGo/chapter_23/01_os_pkg_usage/test.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()



	// ------------ 1. bufio  ReadString ---------------------
	/*
	reader := bufio.NewReader(f)
	for i := 0; i < 10; i++ {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF{
				fmt.Printf("%s \n", str)
				fmt.Println("Complete Read: ", err)
				return
			}
			fmt.Println(err)
			return
		}
		fmt.Printf("%s \n", str)
	}*/

	// ------------ 2. bufio  ReadLine ---------------------
	/*
	reader := bufio.NewReader(f)
	for i := 0; i < 10; i++ {
		b, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF{
				fmt.Printf("%s %t\n", string(b), isPrefix)
				fmt.Println("Complete Read: ", err)
				return
			}
			fmt.Println(err)
			return
		}
		fmt.Printf("%s %t\n", string(b), isPrefix)
	}*/

	// ------------ 3. ioutil.ReadAll == io.ReadAll  ---------------------
	/*
	//all, err := ioutil.ReadAll(f)
	all, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(all))*/

	// ------------ 4. os  ReadFile ---------------------
	/*
	all, err := os.ReadFile("./learnGo/chapter_23/01_os_pkg_usage/test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(all))*/


}

/*
4. os.ReadDir() == ioutil.ReadDir 使用
*/

func osReadDir() {
	dirs, err := os.ReadDir("./learnGo/chapter_23")
	if err != nil {
		panic(err)
	}

	/*
	Name() string

		// IsDir reports whether the entry describes a directory.
		IsDir() bool

		// Type returns the type bits for the entry.
		// The type bits are a subset of the usual FileMode bits, those returned by the FileMode.Type method.
		Type() FileMode

		// Info returns the FileInfo for the file or subdirectory described by the entry.
		// The returned FileInfo may be from the time of the original directory read
		// or from the time of the call to Info. If the file has been removed or renamed
		// since the directory read, Info may return an error satisfying errors.Is(err, ErrNotExist).
		// If the entry denotes a symbolic link, Info reports the information about the link itself,
		// not the link's target.
		Info() (FileInfo, error)*/
	for _, d := range dirs {
		fmt.Println("d.Name() :", d.Name())
		fmt.Println("d.IsDir() :", d.IsDir())
		fmt.Println("d.Type() :", d.Type())
		fmt.Println(d.Info())
		fmt.Println("-----------------------------------")
	}
}



func main() {
	//OpenFileReadUsage()
	//OpenFileWriteUsage()
	//BufioNewReaderUsage()
	//osReadDir()
}
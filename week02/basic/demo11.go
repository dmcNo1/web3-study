package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// testRead()
	// testReadString()
	// testReadFile()
	// testCreateWrite()
	// testOpenFile()
	// testFileInfo()
	testCopyFile()
}

func testRead() {
	// 打开一个文件，默认是以只读的形式打开的
	file, err := os.Open("D:/projects/jackpot/study/web3/web3-study/week02/demo01.go")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file)
	}

	// 关闭文件，一般都是放到defer里
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	buf := make([]byte, 50)
	for {
		count, err := file.Read(buf)
		fmt.Println(string(buf[:count]))
		// 判断是否到了文件末尾，如果是，则会抛出EOF异常
		if err == io.EOF {
			break
		}
	}
}

func testReadString() {
	file, _ := os.Open("D:/projects/jackpot/study/web3/web3-study/week02/demo01.go")

	// 关闭文件，一般都是放到defer里
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// 开启一个reader
	reader := bufio.NewReader(file)
	for {
		// buf, err := reader.ReadBytes('\n')
		str, err := reader.ReadString('\n')
		fmt.Println(str)
		if err == io.EOF {
			break
		}
	}
}

func testReadFile() {
	file, err := os.ReadFile("D:/projects/jackpot/study/web3/web3-study/week02/demo01.go")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(file))
	}
}

func testCreateWrite() {
	file, _ := os.Create("D:/projects/jackpot/study/web3/web3-study/week02/files/test.txt")
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	buf := []byte{'j', 'a', 'c', 'k', 'p', 'o', 't', '\n'}
	file.Write(buf)
	file.WriteString("yiya\n")
	file.WriteString("咿呀\n")
}

// 以不同的模式打开一个文件
func testOpenFile() {
	// 以不同的模式来打开文件，最后一个参数指定了文件权限，在Windows下是没有用的
	// file, err := os.OpenFile("D:/projects/jackpot/study/web3/web3-study/week02/files/test.txt", os.O_CREATE|os.O_WRONLY, 0666)
	file, _ := os.OpenFile("D:/projects/jackpot/study/web3/web3-study/week02/files/test.txt", os.O_CREATE|os.O_APPEND, 0666)
	defer func() {
		file.Close()
	}()
	file.WriteString("咿呀\n")

	// 通过writer的方式写入，这种方式可以使用缓冲池；writer不用关闭
	writer := bufio.NewWriter(file)
	writer.WriteString("devil may cry 6")
	// 还能这样，不过这种写法，会默认覆盖掉
	// os.WriteFile("D:/projects/jackpot/study/web3/web3-study/week02/files/test.txt", []byte{'j', 'a', 'c', 'k', 'p', 'o', 't', '\n'}, 0666)
	// 刷新
	writer.Flush()
}

func testFileInfo() {
	// 获取文件信息
	// type FileInfo interface {
	// 	Name() string       // base name of the file
	// 	Size() int64        // length in bytes for regular files; system-dependent for others
	// 	Mode() FileMode     // file mode bits
	// 	ModTime() time.Time // modification time
	// 	IsDir() bool        // abbreviation for Mode().IsDir()
	// 	Sys() any           // underlying data source (can return nil)
	// }
	fileInfo, err := os.Stat("D:/projects/jackpot/study/web3/web3-study/week02/files/test.txt")
	if err == nil {
		fmt.Println(fileInfo)
	} else if os.IsNotExist(err) {
		fmt.Println("file is not exists")
	} else {
		fmt.Println(err)
	}
}

func testCopyFile() {
	srcDir := "D:/projects/jackpot/study/web3/web3-study/week02/files/srcDir"
	dstDir := "D:/projects/jackpot/study/web3/web3-study/week02/files/dstDir"
	if err := copyDir(srcDir, dstDir); err != nil {
		fmt.Printf("copy failed, error = %v", err)
	}
}

func copyDir(srcDir string, dstDir string) error {
	// 源目录有问题
	dir, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}
	srcDirInfo, _ := os.Stat(srcDir)

	// 如果目标目录不存在，拷贝一份
	_, err = os.Stat(dstDir)
	if os.IsNotExist(err) {
		os.MkdirAll(dstDir, srcDirInfo.Mode())
	}

	// 遍历文件夹下的文件，如果是文件夹，递归创建，否则直接拷贝文件
	for _, file := range dir {
		srcAfterJoin := filepath.Join(srcDir, file.Name())
		dstAfterJoin := filepath.Join(dstDir, file.Name())
		if file.IsDir() {
			if err := copyDir(srcAfterJoin, dstAfterJoin); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcAfterJoin, dstAfterJoin); err != nil {
				return err
			}
		}
	}

	return nil
}

func copyFile(srcPath string, dstPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer func() { srcFile.Close() }()

	srcFileInfo, _ := os.Stat(srcPath)
	dstFile, _ := os.OpenFile(dstPath, os.O_CREATE|os.O_TRUNC, srcFileInfo.Mode())
	defer func() { dstFile.Close() }()

	// 这里可以直接用io.Copy()来的，或者是os.CopyFS()，底层也是用了一个1024*32的[]byte来处理的
	buf := make([]byte, 1024)
	for {
		count, err := srcFile.Read(buf)
		dstFile.Write(buf[:count])
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}

	return nil
}

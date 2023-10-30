package main

var (
	file1   *File
	file2   *File
	file3   *File
	folder1 *Folder
	folder2 *Folder
)

func Init() {
	file1 = &File{name: "file1"}
	file2 = &File{name: "file2"}
	file3 = &File{name: "file3"}

	folder1 = &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}
	folder2 = &Folder{
		children: []Inode{file2, file3},
		name:     "Folder2",
	}
}

func main() {
	Init() // 初始化
	folder2.print("  ")
	cloneFolder := folder2.clone()
	
	cloneFolder.print("  ")
}

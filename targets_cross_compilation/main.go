package main

import "fmt"
import "github.com/360EntSecGroup-Skylar/excelize"

// TODO: gonum

// GOOS=windows
// GOARCH=386, amd64
// GOOS=windows ARCH={X} go build
// targets_cross_compilation.exe: PE32 executable (console) Intel 80386 (stripped to external PDB), for MS Windows
// targets_cross_compilation.exe: PE32+ executable (console) x86-64 (stripped to external PDB), for MS Windows

// GOOS=linux
// GOARCH=386, amd64, arm64
// GOOS=linux GOARCH={X} go build
// targets_cross_compilation: ELF 32-bit LSB executable, Intel 80386, version 1 (SYSV), statically linked, not stripped
// targets_cross_compilation: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, not stripped
// targets_cross_compilation: ELF 64-bit LSB executable, ARM aarch64, version 1 (SYSV), statically linked, not stripped

/*
   $GOOS        $GOARCH
   aix          ppc64
   android      386
   android      amd64
   android      arm
   android      arm64
   darwin       amd64
   darwin       arm64
   dragonfy     amd64
   freebsd      386
   freebsd      amd64
   freebsd      arm
   illumos      amd64
   js           wasm
   linux        386
   linux        amd64
   linux        arm
   linux        arm64
   linux        ppc64
   linux        ppc64le
   linux        mips
   linux        mipsle
   linux        mips64
   linux        mips64le
   linux        s390x
   netbsd       386
   netbsd       amd64
   netbsd       arm
   openbsd      386
   openbsd      amd64
   openbsd      arm
   openbsd      arm64
   plan9        386
   plan9        amd64
   plan9        arm
   solaris      amd64
   windows      386
   windows      amd64
*/

func create_xlsx(fname string) {
	xlsx := excelize.NewFile()
	_ = xlsx.NewSheet("Sheet2")
	_ = xlsx.NewSheet("Sheet3")
	// index2 := xlsx.NewSheet("Sheet2")
	// index3 := xlsx.NewSheet("Sheet3")
	xlsx.SetCellValue("Sheet2", "A1", "Hello world1.")
	xlsx.SetCellValue("Sheet2", "A2", "Hello world2.")
	xlsx.SetCellValue("Sheet2", "A3", "Hello world3.")
	xlsx.SetCellValue("Sheet2", "A4", "Hello world4.")
	xlsx.SetCellValue("Sheet2", "B1", 1)
	xlsx.SetCellValue("Sheet2", "B2", 2)
	xlsx.SetCellValue("Sheet2", "B3", 3)
	xlsx.SetCellValue("Sheet2", "B4", 4)
	xlsx.SetCellValue("Sheet1", "B2", 100)
	xlsx.SetCellValue("Sheet1", "B2", 100)
	xlsx.SetCellValue("Sheet3", "B2", 0.100)
	// xlsx.SetActiveSheet(index3)
	// xlsx.SetActiveSheet(index2)
	err := xlsx.SaveAs(fname + ".xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func read_xlsx(fname string) {
	xlsx, err := excelize.OpenFile(fname)
	if err != nil {
		fmt.Println(err)
		return
	}
	cell := xlsx.GetCellValue("Sheet1", "B2")
	fmt.Println(cell)

	cell = xlsx.GetCellValue("Sheet3", "B2")
	fmt.Println(cell)

	rows := xlsx.GetRows("Sheet2")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

	// cols := xlsx.GetColumns("Sheet1")
}

func main() {
	create_xlsx("test")
	read_xlsx("test.xlsx")
	fmt.Print("in main")
}

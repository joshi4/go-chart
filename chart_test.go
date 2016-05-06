package chart

import "testing"

func TestextraxtGoStmt(t *testing.T) {
	src := `package main 
	import "fmt" 

	func main(){
		go foo()
	}

	func foo(){
		fmt.Println("test") 
	}
	`

	gostmts := extractGoStmnt(src)
	if len(gostmts) != 1 || gostmts[0].parentFn != "main" {
		t.Errorf("got: {len(gostmts)=%d , parentFn=%s}\nwant: {len(gostmts)=1, parentFn=\"main\"}", len(gostmts), gostmts[0].parentFn)
	}
}

//Tests for nucleotide

package tests

import ( "github.com/Winterflower/nucleotide/cmd" 
	"testing"
	"fmt" )

func TestProcess(t *testing.T){
	var dnaString string
	dnaString = "ATCG"
	result := cmd.ProcessDNA(dnaString)
	fmt.Println(result)
}

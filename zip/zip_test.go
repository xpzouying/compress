package zip

import (
	"fmt"
	"os"
	"testing"

	"github.com/xpzouying/compress"
)

func TestZipAFile(t *testing.T) {
	src := "zip_test.go"
	dst := "zip_test.go.zip"

	err := compress.Compress(dst, src)
	if err != nil {
		fmt.Println("Error: ", err)
		t.Error(err)
	} else {
		if err := os.Remove(dst); err != nil {
			panic(err)
		}
	}
}

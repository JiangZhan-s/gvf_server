package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("654884102"))
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$04$s1SDj/QAijcIIY3k2tNUaOCVXNNtqvjknu1lkha9xUsfCrn1GwdbW", "123456"))
}

package golang

import (
	"fmt"
	"github.com/moontrade/dtc-go/codegen"
	"testing"
)

func TestGenerator(t *testing.T) {
	schema, err := codegen.LoadSchema("../testdata/DTCProtocol.proto", "../testdata/DTCProtocol.h", "../testdata/DTCProtocolVLS.h", "../testdata/DTCProtocol_NonStandard.h")
	if err != nil {
		t.Fatal(err)
	}
	g, err := NewGenerator(DefaultConfig("../../model"), schema)
	if err != nil {
		t.Fatal(err)
	}
	err = g.Run()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(g)
}

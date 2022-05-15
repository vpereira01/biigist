//go:build ignore

package main

import (
	"path/filepath"

	"github.com/consensys/gnark-crypto/field"
	"github.com/consensys/gnark-crypto/field/generator"
)

func main() {
	fp, _ := field.NewField("fp1280", "Element", "20815864389328798163850480654728171077230524494533409610638224700807216119346720596024478883464648369684843227908562015582767132496646929816279813211354641525848259018778440691546366699323167100945918841095379622423387354295096957733925002768876520583464697770622321657076833170056511209332449663781837603694136444406281042053396870977465916057756101739472373801429441421111406337458175", false)
	generator.GenerateFF(fp, filepath.Join(".", "fp1280"))
}

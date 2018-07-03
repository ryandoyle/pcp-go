//Copyright (c) 2018 Ryan Doyle
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

package pmda
//#cgo CFLAGS: -I .
//#cgo LDFLAGS: -L . -lpcp -lpcp_pmda
//#include <pcp/pmapi.h>
//#include <pcp/pmda.h>
import "C"
import "unsafe"

type PmdaInterfaceVersion int

const (
	PmdaInterfaceV2 = PmdaInterfaceVersion(int(C.PMDA_INTERFACE_2))
	PmdaInterfaceV3 = PmdaInterfaceVersion(int(C.PMDA_INTERFACE_3))
	PmdaInterfaceV4 = PmdaInterfaceVersion(int(C.PMDA_INTERFACE_4))
	PmdaInterfaceV5 = PmdaInterfaceVersion(int(C.PMDA_INTERFACE_5))
	PmdaInterfaceV6 = PmdaInterfaceVersion(int(C.PMDA_INTERFACE_6))
	PmdaInterfaceV7 = PmdaInterfaceVersion(int(C.PMDA_INTERFACE_7))
	PmdaInterfaceVLatest = PmdaInterfaceVersion(int(C.PMDA_INTERFACE_LATEST))
)

type PMDA interface {
	PmdaConnect()
	PmdaMain()
}

type pmdaInstance struct {
	pmdaInterface *C.pmdaInterface
	domain uint
	name string
}

func New(version PmdaInterfaceVersion, name string, domain uint, logfile string, help_path string) PMDA {
	var pmdaInterface C.pmdaInterface

	name_ptr := C.CString(name)
	defer C.free(unsafe.Pointer(name_ptr))

	logfile_ptr := C.CString(logfile)
	defer C.free(unsafe.Pointer(logfile_ptr))

	help_path_ptr := C.CString(help_path)
	defer C.free(unsafe.Pointer(help_path_ptr))

	C.pmdaDaemon(&pmdaInterface, C.int(version), name_ptr, C.int(domain), logfile_ptr, help_path_ptr)
	return &pmdaInstance{pmdaInterface:&pmdaInterface}
}

func (i *pmdaInstance) PmdaConnect() {
	C.pmdaConnect(i.pmdaInterface)
}

func (i *pmdaInstance) PmdaMain() {
	C.pmdaMain(i.pmdaInterface)
}

func (i *pmdaInstance) PmdaSetFetchCallback(pmda_fetch_callback func() int) {

}
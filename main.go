package main

/*
#cgo LDFLAGS: -lwkhtmltox
#cgo darwin CFLAGS: -I/usr/local/include/wkhtmltox
#cgo linux CFLAGS: -I/usr/include/wkhtmltox
#include <stdbool.h>
#include <stdio.h>
#include <image.h>
#include <pdf.h>

void wkhtmltopdfStrCallback(wkhtmltopdf_converter * converter, const char * str);
void wkhtmltopdfIntCallback(wkhtmltopdf_converter * converter, int val);
void wkhtmltopdfFinishedCallback(wkhtmltopdf_converter * converter, int val);
void wkhtmltopdfVoidCallback(wkhtmltopdf_converter * converter);
*/
import "C"
import (
	"log"
	"unsafe"
)

//export goWkhtmltopdfStrCallback
func goWkhtmltopdfStrCallback(converter *C.wkhtmltopdf_converter, str *C.char) {
	log.Printf("Str: %v", C.GoString(str))
}

//export goWkhtmltopdfIntCallback
func goWkhtmltopdfIntCallback(converter *C.wkhtmltopdf_converter, percent C.int) {
	log.Printf("Progress: %d", percent)
}

//export goWkhtmltopdfFinishedCallback
func goWkhtmltopdfFinishedCallback(converter *C.wkhtmltopdf_converter, intval C.int) {
	log.Printf("Done! %v", intval)
}

//export goWkhtmltopdfVoidCallback
func goWkhtmltopdfVoidCallback(converter *C.wkhtmltopdf_converter) {
	var phaseCount = C.wkhtmltopdf_phase_count(converter) - 1
	var currentPhase = C.wkhtmltopdf_current_phase(converter)
	log.Printf("%v (%d/%d)", C.GoString(C.wkhtmltopdf_phase_description(converter, currentPhase)), currentPhase, phaseCount)
}

func main() {
	log.Printf("Version %v", C.GoString(C.wkhtmltopdf_version()))
	var isExtendedQT = C.wkhtmltopdf_extended_qt() == C.int(1)
	log.Printf("Extended QT? %v", isExtendedQT)

	C.wkhtmltopdf_init(C.int(map[bool]int{true: 1, false: 0}[false]))
	var out = "out.pdf"
	var settings = C.wkhtmltopdf_create_global_settings()
	var globalSettings = [][2]string{
		{"out", out},
		{"size.paperSize", "A4"},
		{"orientation", "Portrait"},
	}
	for _, setting := range globalSettings {
		C.wkhtmltopdf_set_global_setting(settings, C.CString(setting[0]), C.CString(setting[1]))
	}

	var pageSettings = C.wkhtmltopdf_create_object_settings()
	// {"page", "http://wkhtmltopdf.org/libwkhtmltox/"},
	var objectSettings = [][2]string{
		{"page", "http://www.google.de/"}, // - = read from stdin
		{"web.defaultEncoding", "utf-8"},
		{"web.enableJavascript", "true"},
		{"footer.center", "[page] / [toPage]"},
	}
	for _, setting := range objectSettings {
		C.wkhtmltopdf_set_object_setting(pageSettings, C.CString(setting[0]), C.CString(setting[1]))

	}
	var converter = C.wkhtmltopdf_create_converter(settings)

	C.wkhtmltopdf_add_object(converter, pageSettings, nil)

	C.wkhtmltopdf_set_warning_callback(converter, (C.wkhtmltopdf_str_callback)(unsafe.Pointer(C.wkhtmltopdfStrCallback)))
	C.wkhtmltopdf_set_error_callback(converter, (C.wkhtmltopdf_str_callback)(unsafe.Pointer(C.wkhtmltopdfStrCallback)))
	C.wkhtmltopdf_set_phase_changed_callback(converter, (C.wkhtmltopdf_void_callback)(unsafe.Pointer(C.wkhtmltopdfVoidCallback)))
	C.wkhtmltopdf_set_progress_changed_callback(converter, (C.wkhtmltopdf_int_callback)(unsafe.Pointer(C.wkhtmltopdfIntCallback)))
	C.wkhtmltopdf_set_finished_callback(converter, (C.wkhtmltopdf_int_callback)(unsafe.Pointer(C.wkhtmltopdfFinishedCallback)))

	log.Printf("%d", C.wkhtmltopdf_convert(converter)) // 1 success, 0 failure

	//  wkhtmltopdf_get_output if out is missing

	C.wkhtmltopdf_destroy_object_settings(pageSettings)
	C.wkhtmltopdf_destroy_converter(converter)
	C.wkhtmltopdf_destroy_global_settings(settings)
	C.wkhtmltopdf_deinit()
}

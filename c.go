package main

/*
#cgo LDFLAGS: -lwkhtmltox
#cgo darwin CFLAGS: -I/usr/local/include/wkhtmltox
#cgo linux CFLAGS: -I/usr/include/wkhtmltox
#include <pdf.h>

void goWkhtmltopdfStrCallback(wkhtmltopdf_converter * converter, const char * str);

void wkhtmltopdfStrCallback(wkhtmltopdf_converter * converter, const char * str) {
  goWkhtmltopdfStrCallback(converter, str);
}

void goWkhtmltopdfIntCallback(wkhtmltopdf_converter * converter, int str);

void wkhtmltopdfIntCallback(wkhtmltopdf_converter * converter, int str) {
  goWkhtmltopdfIntCallback(converter, str);
}

void goWkhtmltopdfFinishedCallback(wkhtmltopdf_converter * converter, int str);

void wkhtmltopdfFinishedCallback(wkhtmltopdf_converter * converter, int str) {
  goWkhtmltopdfFinishedCallback(converter, str);
}

void goWkhtmltopdfVoidCallback(wkhtmltopdf_converter * converter);

void wkhtmltopdfVoidCallback(wkhtmltopdf_converter * converter) {
  goWkhtmltopdfVoidCallback(converter);
}
*/
import "C"

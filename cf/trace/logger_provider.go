package trace

import (
	"strconv"

	. "github.com/cloudfoundry/cli/cf/i18n"

	"github.com/cloudfoundry/gofileutils/fileutils"
)

func NewLogger(verbose bool, cf_trace, config_trace string) Printer {
	LoggingToStdout = verbose

	var printers []Printer

	stdoutLogger := NewWriterPrinter(stdout)

	for _, path := range []string{cf_trace, config_trace} {
		b, err := strconv.ParseBool(path)
		LoggingToStdout = LoggingToStdout || b

		if path != "" && err != nil {
			file, err := fileutils.Open(path)

			if err == nil {
				printers = append(printers, NewWriterPrinter(file))
			} else {
				stdoutLogger.Printf(T("CF_TRACE ERROR CREATING LOG FILE {{.Path}}:\n{{.Err}}",
					map[string]interface{}{"Path": path, "Err": err}))

				LoggingToStdout = true
			}
		}
	}

	if LoggingToStdout {
		printers = append(printers, stdoutLogger)
	}

	return CombinePrinters(printers)
}
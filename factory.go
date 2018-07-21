package wcai

import (
	"strings"

	"github.com/object88/writing-compilers-and-interpreters/backend"
	"github.com/object88/writing-compilers-and-interpreters/backend/compiler"
	"github.com/object88/writing-compilers-and-interpreters/backend/interpreter"
	"github.com/object88/writing-compilers-and-interpreters/frontend"
	"github.com/object88/writing-compilers-and-interpreters/message"
	"github.com/pkg/errors"
)

func NewBackend(operation string, messageHandler *message.MessageHandler) (backend.Backend, error) {
	switch strings.ToLower(operation) {
	case "execute":
		return compiler.NewCodeGenerator(messageHandler), nil
	case "interpret":
		return interpreter.NewInterpreter(messageHandler), nil
	default:
		return nil, errors.Errorf("Operation '%s' is not supported", operation)

	}
}

func NewParser(language, parserType string, source *frontend.Source) (frontend.Parser, error) {
	switch strings.ToLower(language) {
	case "pascal":
		switch strings.ToLower(parserType) {
		case "top-down":
			s := frontend.NewPascalScanner(source)
			return frontend.NewPascalParserTD(s, nil), nil
		default:
			return nil, errors.Errorf("For language '%s'; parser type '%s' is not supported", language, parserType)
		}
	default:
		return nil, errors.Errorf("Language '%s' is not supported", language)
	}
}
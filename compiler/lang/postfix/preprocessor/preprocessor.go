package postfix

import "github.com/xBu3n0/base-compiler/compiler/analyzer/preprocessor"

type Preprocessor struct{}

func (pp *Preprocessor) Preprocess() {}

func NewPreprocessor() preprocessor.Preprocessor {
	return &Preprocessor{}
}

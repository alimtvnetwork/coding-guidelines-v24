package errcmd

import (
	"fmt"
	"io"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
)

// SafeClose closes an io.Closer and sets faultHolder if no existing fault is recorded.
func SafeClose(closer io.Closer, faultHolder **appfault.AppError) {
	if closer == nil {
		return
	}

	err := closer.Close()
	if err != nil && faultHolder != nil && *faultHolder == nil {
		*faultHolder = appfault.Wrap(errtype.IO, err, "failed to close resource")
	}
}

// SafeCloseIgnore closes an io.Closer ignoring errors.
func SafeCloseIgnore(closer io.Closer) {
	if closer != nil {
		_ = closer.Close()
	}
}

// SafeRecover captures an active panic and records it into faultHolder as a System fault.
func SafeRecover(faultHolder **appfault.AppError, contextMsg string) {
	r := recover()
	if r == nil || faultHolder == nil {
		return
	}

	msg := fmt.Sprintf("%s: recovered panic: %v", contextMsg, r)
	*faultHolder = appfault.New(errtype.Internal, msg)
}

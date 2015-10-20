package libsass

import "io"

type Compiler interface {
	Start() error
}

// SetPath specifies a file to read instead of using the provided
// io.Reader. This activates file compiling that includes line numbers
// in the resulting output.
func SetPath(path string) options {
	return func(c *comp) error {
		c.srcFile = path
		c.ctx.MainFile = path
		return nil
	}
}

type options func(*comp) error

func New(dst io.Writer, src io.Reader, options ...options) (Compiler, error) {

	c := &comp{
		dst: dst,
		src: src,
		ctx: NewContext(),
	}
	c.ctx.in = src
	c.ctx.out = dst

	for _, opt := range options {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

type comp struct {
	ctx     *Context
	dst     io.Writer
	src     io.Reader
	srcFile string
}

func (c comp) start() error {
	if len(c.srcFile) > 0 {
		return c.ctx.FileCompile(c.srcFile, c.dst)
	}
	return c.ctx.Compile(c.src, c.dst)
}

func (c comp) Start() error {
	return c.start()
}

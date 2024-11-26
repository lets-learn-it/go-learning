package math

import "errors"

type Args struct {
	A, B float64
}

type Math struct{}

func (m *Math) Add(args *Args, reply *float64) error {
	*reply = args.A + args.B
	return nil
}

func (m *Math) Sub(args *Args, reply *float64) error {
	*reply = args.A - args.B
	return nil
}

func (m *Math) Mul(args *Args, reply *float64) error {
	*reply = args.A * args.B
	return nil
}

func (m *Math) Div(args *Args, reply *float64) error {
	if args.B == 0.0 {
		return errors.New("divide by 0")
	}

	*reply = args.A / args.B
	return nil
}

package ghooks

import (
	"net"
	"strconv"
)

type Conf struct {
	Port int
}

type Event struct {
	Name string
}

type Hook struct {
	Name string
	Func func(req interface{})
}

type Hooks struct {
	Hooks []Hook
}

func SetConf(port int) (*Conf, error) {
	l, err := net.Listen("tcp", ":"+strconv.Itoa(port))

	if err != nil {
		return nil, err
	}
	l.Close()

	c := &Conf{port}
	return c, nil
}

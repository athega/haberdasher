package haberdasher

//go:generate protoc --proto_path=. --go_out=. --twirp_out=. rpc/haberdasher.proto

import (
	"context"
	"math/rand"

	"github.com/twitchtv/twirp"

	"github.com/athega/haberdasher/rpc"
)

type Logger interface {
	Printf(string, ...interface{})
}

// Service implements the Haberdasher service
type Service struct {
	logger Logger

	colors []string
	names  []string
}

func NewService(logger Logger) *Service {
	return &Service{
		logger: logger,
		colors: []string{"white", "black", "brown", "red", "blue"},
		names:  []string{"bowler", "baseball cap", "top hat", "derby"},
	}
}

func (s *Service) MakeHat(ctx context.Context, size *rpc.Size) (*rpc.Hat, error) {
	if size.Inches <= 0 {
		err := twirp.InvalidArgumentError("inches", "I can't make a hat that small!")

		s.log("\033[0;31mERROR\033[0m   code:%q message:%q", err.Code(), err.Msg())

		return nil, err
	}

	hat := &rpc.Hat{
		Inches: size.Inches,
		Color:  s.randomColor(),
		Name:   s.randomName(),
	}

	s.log("\033[0;32mNEW HAT\033[0m %v", hat)

	return hat, nil
}

func (s *Service) randomColor() string {
	if len(s.colors) == 0 {
		return ""
	}

	return s.colors[rand.Intn(len(s.colors))]
}

func (s *Service) randomName() string {
	if len(s.names) == 0 {
		return ""
	}

	return s.names[rand.Intn(len(s.names))]
}

func (s *Service) log(format string, args ...interface{}) {
	s.logger.Printf(format+"\n", args...)
}

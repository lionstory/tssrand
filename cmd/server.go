package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"runtime"
	"time"

	"tssrand/config"
	pb "tssrand/proto"
	"tssrand/tsslib/keygen"

	"google.golang.org/grpc"
)

const (
	defaultSafePrimeGenTimeout = 5 * time.Minute
)

type RandServer struct {
	pb.UnimplementedRandServer
}

func (s *RandServer) GetRand(ctx context.Context, in *pb.RandRequest) (*pb.RandReply, error) {
	log.Printf(">>>Received: %v", in.GetType())
	if in.GetType() != pb.RandType_ALL {
		return &pb.RandReply{
			Code: pb.ReplyCodeType_ERROR,
			Msg:  "Invalid RandType received.",
		}, errors.New("Invalid RandType received.")
	}

	log.Printf("Start generating randoms....")
	dt_start := time.Now()
	concurrency := runtime.GOMAXPROCS(0)
	preParams, err := keygen.GeneratePreParams(defaultSafePrimeGenTimeout, concurrency)
	if err != nil {
		log.Print("pre-params generation failed:", err)
		return &pb.RandReply{
			Code: pb.ReplyCodeType_ERROR,
			Msg:  "pre-params generation failed",
		}, err
	}
	log.Print("Completed generating randoms.", preParams)
	fmt.Println("===>Cost time: ", time.Since(dt_start))

	reply := &pb.RandReply{
		Code: pb.ReplyCodeType_OK,
		Msg:  "",
		Data: &pb.LocalPreParams{
			PaillierSK: &pb.PrivateKey{
				PublicKey: preParams.PaillierSK.PublicKey.N.Bytes(),
				LambdaN:   preParams.PaillierSK.LambdaN.Bytes(),
				PhiN:      preParams.PaillierSK.PhiN.Bytes(),
			},
			NTildei: preParams.NTildei.Bytes(),
			H1I:     preParams.H1i.Bytes(),
			H2I:     preParams.H2i.Bytes(),
			Alpha:   preParams.Alpha.Bytes(),
			Beta:    preParams.Beta.Bytes(),
			P:       preParams.P.Bytes(),
			Q:       preParams.Q.Bytes(),
		},
	}

	return reply, nil
}

func NewRandServer(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRandServer(s, &RandServer{})
	log.Printf(">>>>>Server started on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	var conf config.Conf
	config, err := conf.GetConf("conf/config.yaml")
	if err != nil {
		log.Fatalf("Conf GetConf err:%s", err)
	}

	NewRandServer(config.Port)
}

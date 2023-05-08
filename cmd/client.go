package main

import (
	"context"
	"log"
	"math/big"

	pb "tssrand/proto"
	"tssrand/tsslib/crypto/paillier"
	"tssrand/tsslib/keygen"

	"google.golang.org/grpc"
)

const (
	serverAddr      = "localhost:50051"
	defaultRandType = pb.RandType_ALL
)

func GetRandRemote(serverAddr string) (*keygen.LocalPreParams, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Printf("could not connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	c := pb.NewRandClient(conn)
	randType := defaultRandType
	reply, err := c.GetRand(context.Background(), &pb.RandRequest{Type: randType})
	if err != nil {
		log.Printf("could not GetRand: %v", err)
		return nil, err
	}

	preParams := &keygen.LocalPreParams{
		PaillierSK: &paillier.PrivateKey{
			PublicKey: paillier.PublicKey{
				N: new(big.Int).SetBytes(reply.Data.GetPaillierSK().PublicKey),
			},
			LambdaN: new(big.Int).SetBytes(reply.Data.GetPaillierSK().GetLambdaN()),
			PhiN:    new(big.Int).SetBytes(reply.Data.GetPaillierSK().GetPhiN()),
		},
		NTildei: new(big.Int).SetBytes(reply.Data.NTildei),
		H1i:     new(big.Int).SetBytes(reply.Data.H1I),
		H2i:     new(big.Int).SetBytes(reply.Data.H2I),
		Alpha:   new(big.Int).SetBytes(reply.Data.Alpha),
		Beta:    new(big.Int).SetBytes(reply.Data.Beta),
		P:       new(big.Int).SetBytes(reply.Data.P),
		Q:       new(big.Int).SetBytes(reply.Data.Q),
	}

	// log.Printf("Received: code: %s, msg: %s", reply.GetCode(), reply.GetMsg())
	log.Printf("Received: data:%v", preParams)

	return preParams, nil
}

func main() {
	GetRandRemote(serverAddr)
}

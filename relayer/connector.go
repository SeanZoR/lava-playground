package relayer

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
)

type Connector struct {
	lock        sync.Mutex
	freeClients []*rpc.Client
	usedClients int
}

func NewConnector(ctx context.Context, nConns uint, addr string) *Connector {
	connector := &Connector{
		freeClients: make([]*rpc.Client, 0, nConns),
	}

	for i := uint(0); i < nConns; i++ {
		var rpcClient *rpc.Client
		var err error
		for {
			if ctx.Err() != nil {
				connector.Close()
				return nil
			}
			nctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
			rpcClient, err = rpc.DialContext(nctx, addr)
			if err != nil {
				log.Println("retrying", err)
				cancel()
				continue
			}
			cancel()
			break
		}
		connector.freeClients = append(connector.freeClients, rpcClient)
	}

	go connector.connectorLoop(ctx)
	return connector
}

func (connector *Connector) connectorLoop(ctx context.Context) {
	<-ctx.Done()
	log.Println("connectorLoop ctx.Done")
	connector.Close()
}

func (connector *Connector) Close() {
	for {
		connector.lock.Lock()
		log.Println("Connector closing", len(connector.freeClients))
		for i := 0; i < len(connector.freeClients); i++ {
			connector.freeClients[i].Close()
		}
		connector.freeClients = []*rpc.Client{}

		if connector.usedClients > 0 {
			log.Println("Connector closing, waiting for in use clients", connector.usedClients)
			connector.lock.Unlock()
			time.Sleep(100 * time.Millisecond)
		} else {
			connector.lock.Unlock()
			break
		}
	}
}

func (connector *Connector) GetRpc(block bool) (*rpc.Client, error) {
	connector.lock.Lock()
	defer connector.lock.Unlock()
	countPrint := 0

	if len(connector.freeClients) == 0 {
		if !block {
			return nil, errors.New("out of clients")
		} else {
			for {
				if countPrint < 3 {
					countPrint++
				}
				connector.lock.Unlock()
				time.Sleep(50 * time.Millisecond)
				connector.lock.Lock()
				if len(connector.freeClients) != 0 {
					break
				}
			}
		}
	}

	ret := connector.freeClients[len(connector.freeClients)-1]
	connector.freeClients = connector.freeClients[:len(connector.freeClients)-1]
	connector.usedClients++

	return ret, nil
}

func (connector *Connector) ReturnRpc(rpc *rpc.Client) {
	connector.lock.Lock()
	defer connector.lock.Unlock()

	connector.usedClients--
	connector.freeClients = append(connector.freeClients, rpc)
}

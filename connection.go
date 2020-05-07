package impalathing

import (
	"context"
	"fmt"
	"log"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/MediaMath/impalathing/services/beeswax"
	impala "github.com/MediaMath/impalathing/services/impalaservice"
)

type Options struct {
	PollIntervalSeconds float64
	BatchSize           int64
}

var (
	DefaultOptions = Options{PollIntervalSeconds: 0.1, BatchSize: 10000}
)

type Connection struct {
	ctx       context.Context
	client    *impala.ImpalaServiceClient
	handle    *beeswax.QueryHandle
	transport thrift.TTransport
	options   Options
	Host      string
	Port      int
}

func Connect(ctx context.Context, host string, port int, options Options) (*Connection, error) {
	socket, err := thrift.NewTSocketTimeout(fmt.Sprintf("%s:%d", host, port), time.Second*30)

	if err != nil {
		return nil, err
	}

	transportFactory := thrift.NewTBufferedTransportFactory(16 * 1024)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, _ := transportFactory.GetTransport(socket)

	if err := transport.Open(); err != nil {
		if transport != nil {
			transport.Close()
		}
		return nil, err
	}

	client := impala.NewImpalaServiceClientFactory(transport, protocolFactory)

	return &Connection{ctx, client, nil, transport, options, host, port}, nil
}

func (c *Connection) isOpen() bool {
	return c.client != nil
}

func (c *Connection) Close() error {
	if c.isOpen() {
		if c.handle != nil {
			status, err := c.client.Cancel(c.ctx, c.handle)
			if err != nil {
				return err
			} else {
				log.Println(status)
			}
			c.handle = nil
		}

		c.transport.Close()
		c.client = nil
	}
	return nil
}

func (c *Connection) CloseQuery(ctx context.Context, handle *beeswax.QueryHandle) error {
	return c.client.Close(ctx, handle)
}

func (c *Connection) CloseInsert(ctx context.Context, handle *beeswax.QueryHandle) (map[string]int64, error) {
	result, err := c.client.CloseInsert(ctx, handle)
	if err != nil {
		return nil, err
	}
	return result.RowsAppended, nil
}

func (c *Connection) ExecuteAndWait(ctx context.Context, query string) (RowSet, error) {
	bquery := beeswax.Query{}

	bquery.Query = query
	bquery.Configuration = []string{}

	handle, err := c.client.ExecuteAndWait(ctx, &bquery, "impala")

	if err != nil {
		return nil, err
	}

	return newRowSet(ctx, c.client, handle, c.options, c.Host, c.Port), nil
}
func (c *Connection) Query(ctx context.Context, query string) (RowSet, error) {
	bquery := beeswax.Query{}

	bquery.Query = query
	bquery.Configuration = []string{}

	handle, err := c.client.Query(ctx, &bquery)

	if err != nil {
		return nil, err
	}

	return newRowSet(ctx, c.client, handle, c.options, c.Host, c.Port), nil
}

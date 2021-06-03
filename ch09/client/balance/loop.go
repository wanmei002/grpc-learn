package balance

import (
    "google.golang.org/grpc/resolver"
    "log"
)

const (
    LoopScheme      = "loop"
    LoopServiceName = "loop.grpc.balance"
)

var addrs = []string{":8093", ":8094"}

type loopBuilder struct{}

func (lb *loopBuilder) Build(target resolver.Target, cc resolver.ClientConn,
    opts resolver.BuildOptions) (resolver.Resolver, error) {
    log.Println("i am build")
    ll := &loopResolver{
        target:     target,
        cc:         cc,
        addrsStore: map[string][]string{
            LoopServiceName: addrs,
        },
    }
    ll.start()
    return ll, nil
}

func (*loopBuilder) Scheme() string {
    log.Println("i am scheme")
    return LoopScheme
}

type loopResolver struct {
    target  resolver.Target
    cc      resolver.ClientConn
    addrsStore map[string][]string
}

func (ll *loopResolver) start() {
    addrList := ll.addrsStore[ll.target.Endpoint]
    log.Printf("loop resolver %+v\n", ll)
    addr := make([]resolver.Address, 0)
    for _, s := range addrList {
        addr = append(addr, resolver.Address{Addr: s})
    }
    log.Println("start start")
    
    ll.cc.UpdateState(resolver.State{Addresses: addr})
}

func (*loopResolver) ResolveNow(o resolver.ResolveNowOptions) {
    log.Println(" i am in resolveNew")
}

func (*loopResolver) Close(){
    log.Println(" i am in close ")
}

func init(){
    resolver.Register(&loopBuilder{})
}



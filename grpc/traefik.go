package grpc

import (
	"fmt"
	"github.com/whileW/core-go/conf"
	"google.golang.org/grpc/resolver"
)

type TraefikBuilder struct{}

var traefik_scheme = "traefik"
func GetTraefikTarget(server_name string) string {
	return fmt.Sprintf("%s:///grpc_%s",traefik_scheme,server_name)
}
func (t *TraefikBuilder)Scheme() string {
	return traefik_scheme
}
func (t *TraefikBuilder)Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &traefikResolver{
		target: target,
		cc:     cc,
	}
	r.start()
	return r,nil
}
type traefikResolver struct {
	target resolver.Target
	cc     resolver.ClientConn
}
func (r *traefikResolver) start() {
	addr := conf.GetConf().Setting.GetStringd("traefik","127.0.0.1:80")
	r.cc.UpdateState(resolver.State{Addresses: []resolver.Address{{Addr: addr,ServerName:r.target.Endpoint}}})
}
func (*traefikResolver) ResolveNow(o resolver.ResolveNowOptions) {}
func (*traefikResolver) Close() {}


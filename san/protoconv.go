package san

import (
	"net"

	"github.com/IPA-CyberLab/kmgm/pb"
)

func FromProtoStruct(s *pb.Names) (Names, error) {
	var ns Names
	if s == nil {
		return ns, nil
	}

	ns.DNSNames = append(ns.DNSNames, s.Dnsname...)
	for _, e := range s.Ipaddr {
		if ipaddr := net.ParseIP(e); ipaddr != nil {
			ns.IPAddrs = append(ns.IPAddrs, ipaddr)
		}
	}
	if err := ns.Verify(); err != nil {
		return Names{}, err
	}

	return ns, nil
}

func (ns Names) ToProtoStruct() *pb.Names {
	ss := make([]string, 0, len(ns.IPAddrs))
	for _, ip := range ns.IPAddrs {
		ss = append(ss, ip.String())
	}

	return &pb.Names{
		Dnsname: append([]string{}, ns.DNSNames...),
		Ipaddr:  ss,
	}
}

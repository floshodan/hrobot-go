package hrobot

import "github.com/floshodan/hrobot-go/hrobot/schema"

//Converts an json Key to SSHKey struct
func SSHKeyFromSchema(s schema.Key) *SSHKey {
	sshKey := &SSHKey{
		Name:        s.Key.Name,
		Fingerprint: s.Key.Fingerprint,
		Type:        s.Key.Type,
		Size:        s.Key.Size,
		Data:        s.Key.Data,
	}
	return sshKey
}

func ServerMarketOrderFromSchema(s schema.Server_market_product) *ServerMarketProduct {
	server := &ServerMarketProduct{
		ID:             s.Product.ID,
		Name:           s.Product.Name,
		Description:    s.Product.Description,
		Traffic:        s.Product.Traffic,
		Dist:           s.Product.Dist,
		Arch:           s.Product.Arch,
		Lang:           s.Product.Lang,
		CPU:            s.Product.CPU,
		CPUBenchmark:   s.Product.CPUBenchmark,
		MemorySize:     s.Product.MemorySize,
		HddSize:        s.Product.HddSize,
		HddText:        s.Product.HddText,
		HddCount:       s.Product.HddCount,
		Datacenter:     s.Product.Datacenter,
		NetworkSpeed:   s.Product.NetworkSpeed,
		Price:          s.Product.Price,
		PriceSetup:     s.Product.PriceSetup,
		PriceVat:       s.Product.PriceVat,
		PriceSetupVat:  s.Product.PriceSetupVat,
		FixedPrice:     s.Product.FixedPrice,
		NextReduce:     s.Product.NextReduce,
		NextReduceDate: s.Product.NextReduceDate,
	}
	return server
}

func ServerProductFromSchema(s schema.ServerProduct) *ServerProduct {
	server := &ServerProduct{
		ID:          s.Product.ID,
		Name:        s.Product.Name,
		Description: s.Product.Description,
		Traffic:     s.Product.Traffic,
		Dist:        s.Product.Dist,
		Arch:        s.Product.Arch,
		Lang:        s.Product.Lang,
		Location:    s.Product.Location,
		Prices: []struct {
			Location string
			Price    struct {
				Net   string
				Gross string
			}
			PriceSetup struct {
				Net   string
				Gross string
			}
		}(s.Product.Prices),
	}
	return server
}

func ServerOrderTransactionFromSchema(s schema.ServerOrderTransaction) *ServerOrderTransaction {
	transaction := &ServerOrderTransaction{
		ID:           s.Transaction.ID,
		Date:         s.Transaction.Date,
		Status:       s.Transaction.Status,
		ServerNumber: s.Transaction.ServerNumber,
		AuthorizedKey: []struct {
			Key struct {
				Name        string
				Fingerprint string
				Type        string
				Size        int
			}
		}(s.Transaction.AuthorizedKey),
		HostKey: s.Transaction.HostKey,
		Comment: s.Transaction.Comment,
		Product: struct {
			ID          string
			Name        string
			Description []string
			Traffic     string
			Dist        string
			Arch        string
			Lang        string
			Location    string
		}{},
	}
	return transaction
}

func ServerMarketTransactionFromSchema(s schema.ServerMarketTransaction) *ServerMarketTransaction {
	transaction := &ServerMarketTransaction{
		ID:           s.Transaction.ID,
		Date:         s.Transaction.Date,
		Status:       s.Transaction.Status,
		ServerNumber: s.Transaction.ServerNumber,
		ServerIP:     s.Transaction.ServerIP,
		AuthorizedKey: []struct {
			Key struct {
				Name        string
				Fingerprint string
				Type        string
				Size        int
			}
		}(s.Transaction.AuthorizedKey),
		HostKey: s.Transaction.HostKey,
		Comment: s.Transaction.Comment,
		Product: struct {
			ID           int
			Name         string
			Description  []string
			Traffic      string
			Dist         string
			Arch         string
			Lang         string
			CPU          string
			CPUBenchmark int
			MemorySize   int
			HddSize      int
			HddText      string
			HddCount     int
			Datacenter   string
			NetworkSpeed string
		}{},
	}
	return transaction

}

func ServerFromSchema(s schema.Server) *Server {
	server := &Server{
		ServerIP:      s.Server.ServerIP,
		ServerIpv6Net: s.Server.ServerIpv6Net,
		ServerNumber:  s.Server.ServerNumber,
		ServerName:    s.Server.ServerName,
		Product:       s.Server.Product,
		Dc:            s.Server.Dc,
		Traffic:       s.Server.Traffic,
		Status:        s.Server.Status,
		Cancelled:     s.Server.Cancelled,
		PaidUntil:     s.Server.PaidUntil,
		IP:            s.Server.IP,
		Subnet: []struct {
			IP   string
			Mask string
		}(s.Server.Subnet),
		LinkedStoragebox: s.Server.LinkedStoragebox,
	}
	return server
}

func SingleServerFromSchema(s schema.SingleServer) *SingleServer {
	server := &SingleServer{
		ServerIP:      s.Server.ServerIP,
		ServerIpv6Net: s.Server.ServerIpv6Net,
		ServerNumber:  s.Server.ServerNumber,
		ServerName:    s.Server.ServerName,
		Product:       s.Server.Product,
		Dc:            s.Server.Dc,
		Traffic:       s.Server.Traffic,
		Status:        s.Server.Status,
		Cancelled:     s.Server.Cancelled,
		PaidUntil:     s.Server.PaidUntil,
		IP:            s.Server.IP,
		Subnet: []struct {
			IP   string
			Mask string
		}(s.Server.Subnet),
		Reset:            s.Server.Reset,
		Rescue:           s.Server.Rescue,
		Vnc:              s.Server.Vnc,
		Windows:          s.Server.Windows,
		Plesk:            s.Server.Plesk,
		Cpanel:           s.Server.Cpanel,
		Wol:              s.Server.Wol,
		HotSwap:          s.Server.HotSwap,
		LinkedStoragebox: s.Server.LinkedStoragebox,
	}
	return server
}

func CancellationFromSchema(s schema.Cancellation) *Cancellation {
	cancellation := &Cancellation{
		ServerIP:                 s.Cancellation.ServerIP,
		ServerIpv6Net:            s.Cancellation.ServerIpv6Net,
		ServerNumber:             s.Cancellation.ServerNumber,
		ServerName:               s.Cancellation.ServerName,
		EarliestCancellationDate: s.Cancellation.EarliestCancellationDate,
		Cancelled:                s.Cancellation.Cancelled,
		ReservationPossible:      s.Cancellation.ReservationPossible,
		Reserved:                 s.Cancellation.Reserved,
		CancellationDate:         s.Cancellation.CancellationDate,
		CancellationReason:       s.Cancellation.CancellationReason,
	}
	return cancellation
}

func RescueFromSchema(s schema.BootRescue) *BootRescue {
	rescue := &BootRescue{
		ServerIP:      s.Rescue.ServerIP,
		ServerIpv6Net: s.Rescue.ServerIpv6Net,
		ServerNumber:  s.Rescue.ServerNumber,
		Os:            s.Rescue.Os,
		Arch:          s.Rescue.Arch,
		Active:        s.Rescue.Active,
		Password:      s.Rescue.Password,
		AuthorizedKey: s.Rescue.AuthorizedKey,
		HostKey:       s.Rescue.HostKey,
	}
	return rescue
}

func IPFromSchema(s schema.IP) *IP {
	ip := &IP{
		IP:              s.IP.IP,
		ServerIP:        s.IP.ServerIP,
		ServerNumber:    s.IP.ServerNumber,
		Locked:          s.IP.Locked,
		SeparateMac:     s.IP.SeparateMac,
		TrafficWarnings: s.IP.TrafficWarnings,
		TrafficHourly:   s.IP.TrafficHourly,
		TrafficDaily:    s.IP.TrafficDaily,
		TrafficMonthly:  s.IP.TrafficMonthly,
	}
	return ip
}

func MACFromSchema(s schema.MAC) *MAC {
	mac := &MAC{
		IP:  s.Mac.IP,
		MAC: s.Mac.Mac,
	}
	return mac
}

func SubnetFromSchema(s schema.Subnet) *Subnet {
	subnet := &Subnet{
		IP:              s.Subnet.IP,
		Mask:            s.Subnet.Mask,
		Gateway:         s.Subnet.Gateway,
		ServerIP:        s.Subnet.ServerIP,
		ServerNumber:    s.Subnet.ServerNumber,
		Failover:        s.Subnet.Failover,
		Locked:          s.Subnet.Locked,
		TrafficWarnings: s.Subnet.TrafficWarnings,
		TrafficHourly:   s.Subnet.TrafficHourly,
		TrafficDaily:    s.Subnet.TrafficDaily,
		TrafficMonthly:  s.Subnet.TrafficMonthly,
	}
	return subnet
}

func ResetFromSchema(s schema.Reset) *Reset {
	reset := &Reset{
		ServerIP:        s.Reset.ServerIP,
		ServerIpv6Net:   s.Reset.ServerIpv6Net,
		ServerNumber:    s.Reset.ServerNumber,
		Type:            s.Reset.Type,
		OperatingStatus: s.Reset.OperatingStatus,
	}
	return reset
}

func WOLFromSchema(s schema.WOL) *WOL {
	wol := &WOL{
		ServerIP:      s.Wol.ServerIP,
		ServerIpv6Net: s.Wol.ServerIpv6Net,
		ServerNumber:  s.Wol.ServerNumber,
	}
	return wol
}

func FirewallFromSchema(s schema.Firewall) *Firewall {
	firewall := &Firewall{
		ServerIP:     s.Firewall.ServerIP,
		ServerNumber: s.Firewall.ServerNumber,
		Status:       s.Firewall.Status,
		WhitelistHos: s.Firewall.WhitelistHos,
		Port:         s.Firewall.Port,
		Rules: struct {
			Input []struct {
				IPVersion string
				Name      string
				DstIP     interface{}
				SrcIP     string
				DstPort   string
				SrcPort   interface{}
				Protocol  interface{}
				TCPFlags  interface{}
				Action    string
			}
		}(s.Firewall.Rules),
	}
	return firewall
}

func FirewallTemplateFromSchema(s schema.FirewallTemplate) *FirewallTemplate {
	firewalltemplate := &FirewallTemplate{
		ID:           s.FirewallTemplate.ID,
		Name:         s.FirewallTemplate.Name,
		WhitelistHos: s.FirewallTemplate.WhitelistHos,
		IsDefault:    s.FirewallTemplate.IsDefault,
	}
	return firewalltemplate
}

func FirewallTemplateWithRulesFromSchema(s schema.FirewallTemplateWithRules) *FirewallTemplateWithRules {
	firewall := &FirewallTemplateWithRules{
		ID:           s.FirewallTemplate.ID,
		WhitelistHos: s.FirewallTemplate.WhitelistHos,
		IsDefault:    s.FirewallTemplate.IsDefault,
		Rules: struct {
			Input []struct {
				IPVersion string
				Name      string
				DstIP     interface{}
				SrcIP     string
				DstPort   string
				SrcPort   interface{}
				Protocol  interface{}
				TCPFlags  interface{}
				Action    string
			}
		}(s.FirewallTemplate.Rules),
	}
	return firewall
}

func VswitchFromSchema(s schema.VSwitch) *VSwitch {
	vswitch := &VSwitch{
		ID:        s.ID,
		Name:      s.Name,
		Vlan:      s.Vlan,
		Cancelled: s.Cancelled,
	}
	return vswitch
}

func VSwitchSingleFromSchema(s schema.VSwitchSingle) *VSwitchSingle {
	vswitch := &VSwitchSingle{
		ID:           s.ID,
		Name:         s.Name,
		Vlan:         s.Vlan,
		Cancelled:    s.Cancelled,
		Server:       s.Server,
		Subnet:       s.Subnet,
		CloudNetwork: s.CloudNetwork,
	}
	return vswitch
}

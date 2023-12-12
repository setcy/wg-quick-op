package quick

import (
	"github.com/sirupsen/logrus"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func PeerStatus(iface string) (map[wgtypes.Key]*wgtypes.Peer, error) {
	c, err := wgctrl.New()
	defer func(c *wgctrl.Client) {
		err := c.Close()
		if err != nil {
			logrus.WithError(err).WithField("iface", iface).Error("failed to close client: ", err.Error())
		}
	}(c)
	if err != nil {
		return nil, err
	}
	device, err := c.Device(iface)
	if err != nil {
		return nil, err
	}

	peers := make(map[wgtypes.Key]*wgtypes.Peer)
	for _, p := range device.Peers {
		peers[p.PublicKey] = &p
	}
	return peers, nil
}

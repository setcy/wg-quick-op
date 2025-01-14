package cmd

import (
	"github.com/hdu-dn11/wg-quick-op/quick"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync (deprecated)",
	Short: "sync [interface name]",
	Long: `sync [interface name], sync link,address,device and route. Notice that PostUp and PreUp won't run
it may result in address added by PostUp being deleted.'`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			logrus.Errorln("up command requires exactly one interface name")
			return
		}
		cfgs := quick.MatchConfig(args[0])
		for iface, cfg := range cfgs {
			err := quick.Sync(cfg, iface, logrus.WithField("iface", iface))
			if err != nil {
				logrus.WithError(err).Errorln("failed to sync interface")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}

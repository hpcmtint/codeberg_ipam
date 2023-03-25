/*
Copyright © 2023 Laura Kalb <dev@lauka.net>
*/

package cmd

import (
    "encoding/json"
    "fmt"
    "net/netip"
    "os"
    "path"
    "time"

    "github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
    Use:   "export",
    Short: "Export ipam configuration",
    Long: `Export current ipam contents to importable data format.
You can either export a single subnet or all subnets.`,
    Example: "ipam export\nipam export 192.168.0.0/24",
    Args:    cobra.RangeArgs(0, 1),
    Aliases: []string{"ex"},
    Run: func(cmd *cobra.Command, args []string) {
        var jsonoutput []Subnet
        if len(args) == 1 {
            prefix, parseerr := netip.ParsePrefix(args[0])
            if parseerr != nil {
                fmt.Println("[ERROR]", parseerr)
                os.Exit(1)
            }

            subnet, geterr := GetSubnet(prefix)
            if geterr != nil {
                fmt.Println("[ERROR]", geterr)
                os.Exit(1)
            }
            jsonoutput = append(jsonoutput, subnet)
        } else {
            subnetlist := ListSubnets()
            for _, net := range subnetlist {
                prefix, parseerr := netip.ParsePrefix(net)
                if parseerr != nil {
                    fmt.Println("[ERROR]", parseerr)
                    os.Exit(1)
                }
                subnet, geterr := GetSubnet(prefix)
                if geterr != nil {
                    fmt.Println("[ERROR]", geterr)
                    os.Exit(1)
                }
                jsonoutput = append(jsonoutput, subnet)
            }
        }

        //workingdir, _ := os.Getwd()
        //timestamp := time.Now().Format("2006-01-02_15-04")
        //exportfilename := workingdir + "/ipam_export_" + timestamp + ".json"

        var exportname string

        flagpath, _ := cmd.Flags().GetString("file")
        if path.IsAbs(flagpath) {
            exportname = flagpath
        } else {
            wd, _ := os.Getwd()
            exportname = path.Join(wd, flagpath)
        }

        data, _ := json.MarshalIndent(jsonoutput, "", "  ")

        file, fileerr := os.Create(exportname)
        if fileerr != nil {
            fmt.Println("[ERROR]", fileerr)
            os.Exit(1)
        }
        defer file.Close()

        _, writeerr := file.Write(data)
        if writeerr != nil {
            fmt.Println("[ERROR]", writeerr)
            os.Exit(1)
        }
        fmt.Printf("[INFO] Data was exported to file %v\n", exportname)
    },
}

func init() {
    rootCmd.AddCommand(exportCmd)

    timestamp := time.Now().Format("2006-01-02_15-04")
    exportCmd.Flags().StringP("file", "f", "./ipam_export_"+timestamp+".json", "File name for exported data.\nCan be both absolute or relative path.")
}

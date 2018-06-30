/*
Copyright 2018 Safewrd Ventures OÜ

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	goflag "flag"
	"fmt"
	"os"

	"github.com/golang/glog"

	"github.com/SAFEWRD/enterperm/pkg/utils"
	"github.com/SAFEWRD/enterperm/pkg/version"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var rootCmd = &cobra.Command{
	Use:   "enterperm",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// tbd
	},
}

var podCountCmd = &cobra.Command{
	Use:   "pod-count",
	Short: "Counts all pods in the kubernetes cluster",
	Long: `Call used to test connectivity to kubernetes clusters. 
		Simply outputs the number of running pods on the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		k8s := utils.GetClientExternal()
		pods, err := k8s.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			glog.Exitln(err)
		}

		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of enterperm",
	Long:  `All software has versions. This is enterperm's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Version())
	},
}

func init() {
	rootCmd.AddCommand(podCountCmd)
	rootCmd.AddCommand(versionCmd)
}

// Execute runs cobra to process commands and flags
func Execute() {
	// add compatibility for glog flags
	goflag.Set("logtostderr", "true")
	goflag.CommandLine.Parse([]string{})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/arken/arkstrap/config"
	"github.com/arken/arkstrap/ipfs"
)

func main() {
	node, err := ipfs.CreateNode(config.Global.Ipfs.Path, ipfs.NodeConfArgs{
		Addr:           config.Global.Ipfs.Addr,
		PeerID:         config.Global.Ipfs.PeerID,
		PrivKey:        config.Global.Ipfs.PrivateKey,
		SwarmKey:       config.Manifest.ClusterKey,
		BootstrapPeers: config.Manifest.BootstrapPeers,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Print Node ID
	fmt.Printf("ID: %s\n", node.ID())

	// Wait for a SIGINT or SIGTERM signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")
	os.Exit(0)
}

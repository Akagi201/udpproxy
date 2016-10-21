package main

import (
	"net"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	Source string `long:"source" default:":2203" description:"Source port to listen on"`
	Target string `long:"target" default:"127.0.0.1:2202" description:"Target address to forward to"`
	Quiet  bool   `long:"quiet" description:"whether to print logging info or not"`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		if !strings.Contains(err.Error(), "Usage") {
			log.Printf("error: %v\n", err.Error())
			os.Exit(1)
		} else {
			log.Printf("%v\n", err.Error())
			os.Exit(0)
		}
	}

	if opts.Quiet {
		log.SetLevel(log.WarnLevel)
	}

	sourceAddr, err := net.ResolveUDPAddr("udp", opts.Source)
	if err != nil {
		log.WithError(err).Fatal("Could not resolve source address:", opts.Source)
		return
	}

	targetAddr, err := net.ResolveUDPAddr("udp", opts.Target)
	if err != nil {
		log.WithError(err).Fatal("Could not resolve target address:", opts.Target)
		return
	}

	sourceConn, err := net.ListenUDP("udp", sourceAddr)
	if err != nil {
		log.WithError(err).Fatal("Could not listen on address:", opts.Source)
		return
	}

	defer sourceConn.Close()

	targetConn, err := net.DialUDP("udp", nil, targetAddr)
	if err != nil {
		log.WithError(err).Fatal("Could not connect to target address:", opts.Target)
		return
	}

	defer targetConn.Close()

	log.Printf(">> Starting udpproxy, Source at %v, Target at %v...", opts.Source, opts.Target)

	for {
		b := make([]byte, 10240)
		n, addr, err := sourceConn.ReadFromUDP(b)

		if err != nil {
			log.WithError(err).Error("Could not receive a packet")
			continue
		}

		log.WithField("addr", addr.String()).WithField("bytes", n).WithField("content", string(b)).Info("Packet received")
		if _, err := targetConn.Write(b[0:n]); err != nil {
			log.WithError(err).Warn("Could not forward packet.")
		}
	}
}

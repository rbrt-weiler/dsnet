package dsnet

func Add(hostname string, owner string, description string) {//, publicKey string) {
	conf := MustLoadDsnetConfig()

	privateKey := GenerateJSONPrivateKey()
	presharedKey := GenerateJSONKey()
	publicKey := privateKey.PublicKey()

	conf.ChooseIP()

	peer := PeerConfig{
		Owner: owner,
		Hostname: hostname,
		Description: description,
		PublicKey: publicKey,
		PresharedKey: presharedKey,
		// TODO Endpoint:
		// TODO pick an available IP AllowedIPs
	}

	conf.MustAddPeer(peer)
	conf.MustSave()
}

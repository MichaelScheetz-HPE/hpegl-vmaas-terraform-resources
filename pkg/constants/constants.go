// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

// Package constants - constants that are used in pkg/client and pkg/resources
package constants

const (
	// ServiceName - the service mnemonic
	ServiceName    = "vmaas"
	AccServiceURL  = "https://iac-vmaas.dev.hpehcss.net"
	IntgServiceURL = "https://iac-vmaas.intg.hpedevops.net"
	ServiceURL     = "https://iac-vmaas.us1.greenlake-hpe.com"

	LOCATION    = "location"
	SPACENAME   = "space_name"
	INSECURE    = "allow_insecure"
	SpaceKey    = "space"
	LocationKey = "location"

	AccLocation = "HPEGL-VMAAS-TERRAFORM"
	AccSpace    = "Default"

	MockIAMKey    = "TF_ACC_MOCK_IAM"
	CmpSubjectKey = "TF_ACC_CMP_SUBJECT"
)

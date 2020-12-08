package azure

// LoadBalancerIPType enumerator for types Public, Private or No IP.
type LoadBalancerIPType string

// LoadBalancerIPType values
const (
	PublicIP  LoadBalancerIPType = "PublicIP"
	PrivateIP LoadBalancerIPType = "PrivateIP"
	NoIP      LoadBalancerIPType = "NoIP"
)

package object

type ApplicationsXml struct {
	Application ApplicationItem `xml:"application"`
}

type ApplicationItem struct {
	Name string `xml:"name"`
	Instance[] InstanceItem `xml:"instance"`
}
type InstanceItem struct {
	HostName string `xml:"hostName"`
	App string `xml:"app"`
	IpAddr string `xml:"ipAddr"`
	Status string `xml:"status"`
	Port string `xml:"port"`
	RealLink string `xml:"realLink"`
}

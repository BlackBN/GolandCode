package main
//
import (
	"fmt"
	"strconv"
	"strings"

	//"encoding/json"
	//"fmt"
	////"github.com/caicloud/nirvana/log"
	////"github.com/willf/bitset"
	//"net"
	//"strings"
	"github.com/containernetworking/cni/pkg/types"
	"net"
)
//type NetworkPart struct {
//	Name        string `json:"name"`
//	DisplayName string `json:"displayName"`
//	Description string `json:"description"`
//	IsDefault   bool   `json:"isDefaultNetwork"`
//	IsFixedIP   bool   `json:"isFixedIP"`
//	Type        string `json:"type"`
//}
//
//
//
//type NetworkList struct {
//	Items []NetworkInfo `json:"items"`
//}
//
//type NetworkInfo struct {
//	name string `json:"name"`
//	age int `json:"age"`
//	sex string`json:"sex"`
//}
//
//
//type void struct {
//
//}
//
//type FixedIPIPAM struct {
//	tesxt string
//	networks      map[string]*Network
//}
//
//type Network struct {
//	Subnets           []*Subnet          `json:"subnets"`
//	TenantSubnets     map[string]*Subnet `json:"tenantSubnets"`
//}
//
//type Subnet struct {
//	subpools     map[string]*WorkloadIPPool
//	allocator    *ipAllocator
//}
//
//type WorkloadIPPool struct {
//	poolID     string
//	savedReversion uint32
//}
//
//type ipAllocator struct {
//	baseIP uint
//	prefix uint
//	min    uint
//	max    uint
//	bitSet bitset.BitSet
//	used   uint
//}
//
//const wordSize = uint(64)
//const log2WordSize = uint(6)
//
//func main(){
//	//c := 18446744073709551615
//	//a := ^uint(0)
//	//log.Infof("a : %d",a)
//	//if 18446744073709551615 == -1 {
//	//	log.Infof("a equal -1")
//	//}else {
//	//	log.Infof("no")
//	//}
//	//b := a - wordSize + 1
//	//log.Infof("b : %d",b)
//	//
//	//log.Infof( "b1 : %d " , a << 6 )
//	//
//	//
//	//
//	//c := int(a >> log2WordSize)
//	//log.Infof("c : %d",c)
//	//
//	//log.Infof(" 2 ^ 58 :               %d", 1 << 58 - 1)
//	//log.Infof("uint(2 ^ 64 -1 ) >> 6 : %d" , ^uint(0) >> 6)
//	//
//	//
//	//e := int((4096 + (wordSize - 1)) >> log2WordSize)
//	//log.Infof("e : %d",e)
//	//i := uint(10)
//	//i1 := i >> log2WordSize
//	//log.Infof("i1 : %d ",i1)
//	//i2 := i & (wordSize - 1)
//	//log.Infof("i2 : %d ",i2)
//	//i3 := 1 << i2
//	//log.Infof("i3 : %d ",i3)
//
//
//
//	//var bit bitset.BitSet
//	//bit.Set(10)
//	//bit.Set(68)
//	//bit.Set(4)
//	//bit.Set(56)
//	//bit.Set(90)
//	//a,d,_ := GetBitset(&bit)
//	//fmt.Println(a)
//	//fmt.Println(d)
//	//bit.set[i>>log2WordSize] |= 1 << (i & (wordSize - 1))
//
//	//log.Infof("end")
//
//	// ip := "192.168.0.0"
//
//
//	// prefix, _ := cidr.Mask.Size()
//	// baseIP := ip4ToInt(cidr.IP.Mask(cidr.Mask))
//	// maxSize := (1 << (32 - uint(prefix))) - 1
//
//	// log.Infof("perfix : %d ; baseIP := %d ; maxSize := %d",prefix,baseIP,maxSize)
//
//
//	// workloads := make([][3]string, 9)
//	//for  i := 0; i < 9; i++ {
//	//	for j := 0 ; j < 3 ; j++ {
//	//		 workloads[i][j]  = fmt.Sprintf("%d",i+j+2)
//	//	}
//	//}
//	//
//	//for  i := 0; i < 9; i++ {
//	//	for j := 0 ; j < 3 ; j++ {
//	//		fmt.Printf("a[%d][%d] = %s\n", i , j , workloads[i][j])
//	//	}
//	//}
//	/*var empty void
//	set := make(map[string]void)
//	set["age"]=empty
//	for k := range set {
//		fmt.Println(k)
//	}
//	//delete(set,"age")
//	size := len(set)
//	fmt.Println(size)
//	exists:= set["age"]
//	fmt.Println(exists)*/
//	//p := &FixedIPIPAM{
//	//	tesxt:"aaa",
//	//	networks: map[string]*Network{
//	//		"one":{
//	//			Subnets:       []*Subnet{
//	//				{
//	//					subpools: map[string]*WorkloadIPPool{
//	//						"one": {
//	//							poolID:         "abcdefg",
//	//							savedReversion: 1,
//	//						},
//	//						"two":{
//	//							poolID:"sadfasdfads",
//	//							savedReversion:4,
//	//						},
//	//					},
//	//					allocator: &ipAllocator{
//	//						baseIP: 4,
//	//						prefix: 5,
//	//						min:    5,
//	//						max:    5,
//	//						used:   5,
//	//					},
//	//				},
//	//
//	//			},
//	//		},
//	//	},
//	//}
//	//log.Infof("aaa : %v",*p)
//	//for k,v := range p.networks {
//	//	log.Infof("k : %s",k)
//	//	subnets := v.Subnets
//	//	for _,v := range subnets {
//	//		pool := v.subpools
//	//		log.Infof("%v",pool)
//	//	}
//	//}
//	IP := net.ParseIP("1.1.1.8")
//	ip := IP.To4()
//	ipUint := uint(ip[3]) | (uint(ip[2]) << 8) | (uint(ip[1]) << 16) | (uint(ip[0]) << 24)
//	log.Infof("p.network.tenantSubnet range start : %d",ipUint)
//
//
//	aaa:= IP.String()
//	log.Infof("aa : %s",aaa)
//
//	cidrIP := "192.168.14.93"
//	startIP := "192.168.0.30"
//	endIP := "192.168.0.32"
//	gatewayIP := "192.168.0.1"
//	cidr :=  &net.IPNet{
//		IP: net.ParseIP(cidrIP),
//		Mask: net.IPv4Mask(byte(255), byte(255), byte(253), byte(0)),
//	}
//
//	prefix, _ := cidr.Mask.Size()
//	fmt.Printf("prefix : %d",prefix)
//	fmt.Println()
//	maxSize := (1 << (32 - uint(prefix))) - 1
//	fmt.Printf("maxsize : %d",maxSize)
//	fmt.Println()
//
//
//	start , end,gateway := net.ParseIP(startIP),net.ParseIP(endIP),net.ParseIP(gatewayIP)
//	ipAcc,_ := newIPAllocator(cidr,start,end,gateway)
//	fmt.Println(ipAcc.baseIP)
//	fmt.Println(ipAcc.bitSet.Len())
//	fmt.Println(ipAcc.bitSet.Count())
//	fmt.Println(ipAcc.bitSet.Len()-ipAcc.bitSet.Count())
//
//	str := "AAAAAAAAAQD_____P____________________________________w=="
//
//	bitset := new(bitset.BitSet)
//	bitset.UnmarshalJSON([]byte(str))
//	fmt.Println(bitset.Len())
//	fmt.Println(bitset.Count())
//
//
//	deletedWorkloads := make([][3]string, 0)
//	a:=[3]string{
//		"a",
//		"b",
//		"c",
//	}
//
//	deletedWorkloads = append(deletedWorkloads, a)
//
//	ListNetworkUsedRange()
//
//	tenantTest := TenantSubnet{
//		Spec:TenantSubnetSpec{
//			Tenant:  "aba",
//			Subnets: map[string]SubSpec{
//				"/default/abcd":{
//					ID:"aaa",
//
//				},
//				"/default/edf":{
//					ID:"bbb",
//				},
//			},
//		},
//	}
//
//	clusterNetworkList := NetworkList{
//		Items: []NetworkInfo{
//			{name:"abcd",age:12,sex:"man"},
//			{name:"edf",age:23,sex:"woman"},
//			{name:"efdfa"},
//		},
//	}
//
//	subnets := tenantTest.Spec.Subnets
//
//	log.Infof("tenant subnet length : %d",len(subnets))
//
//	networks := make([]*NetworkPart, len(subnets))
//
//	networkMap := make(map[string]NetworkInfo,len(clusterNetworkList.Items))
//	for _ ,network  := range clusterNetworkList.Items {
//		log.Infof("network name is : %s",network.name)
//		networkMap[network.name] = network
//	}
//
//	index := 0
//	for networkId := range subnets {
//		_, _ , networkName := SplitID(networkId)
//		net := networkMap[networkName]
//
//		defaultNet := false
//		if net.name == "abcd" {
//			defaultNet = true
//		}
//		networks[index] = &NetworkPart{
//			Name:        net.name,
//			DisplayName: net.sex,
//			IsDefault:   defaultNet,
//
//		}
//		index = index + 1
//	}
//	log.Infof("%d",len(networks))
//
//
//
//
//
//}
//type IPRange struct {
//	Start net.IP `json:"start"`
//	End   net.IP `json:"end"`
//}
//
//type NetworkIPRange struct {
//	Network string `json:"network"`
//	IPRangeList []IPRange `json:"items"`
//}
//
//type IPRangeList struct {
//	Items    []NetworkIPRange `json:"items"`
//}
//
//type TenantSubnetList struct {
//	// Items is the list of Networks
//	Items []TenantSubnet `json:"items"`
//}
//
//
//type TenantSubnet struct {
//
//	Spec              TenantSubnetSpec   `json:"spec"`
//}
//
//type TenantSubnetSpec struct {
//	Tenant  string             `json:"tenant"`
//	Subnets map[string]SubSpec `json:"subnets,omitempty"`
//}
//
//type SubSpec struct {
//	ID         string `json:"id"`
//	CIDR       string `json:"cidr"`
//	Gateway    net.IP `json:"gateway,omitempty"`
//	RangeStart net.IP `json:"rangeStart,omitempty"`
//	RangeEnd   net.IP `json:"rangeEnd,omitempty"`
//	SubStart   net.IP `json:"subStart,omitempty"`
//	SubEnd     net.IP `json:"subEnd,omitempty"`
//}
//
//
//func ListNetworkUsedRange() (*IPRangeList, error) {
//
//	tsList := TenantSubnetList{}
//
//
//	tsOne := TenantSubnet{Spec:TenantSubnetSpec{
//		Tenant:  "abcdefg",
//		Subnets: map[string]SubSpec{
//			"AAAA": {
//				ID:         "AAAA",
//				CIDR:       "AAAA",
//				Gateway:    net.ParseIP("127.0.0.0"),
//				SubStart:   net.ParseIP("192.168.0.20"),
//				SubEnd:     net.ParseIP("192.168.0.34"),
//			},
//			"BBBB": {
//				ID:         "BBBB",
//				CIDR:       "BBBB",
//				Gateway:    net.ParseIP("127.0.0.0"),
//				SubStart:   net.ParseIP("192.168.0.45"),
//				SubEnd:     net.ParseIP("192.168.0.60"),
//			},
//		},
//	}}
//
//	tsTwo := TenantSubnet{Spec:TenantSubnetSpec{
//		Tenant:  "ddddddd",
//		Subnets: map[string]SubSpec{
//			"AAAA": {
//				ID:         "AAAA",
//				CIDR:       "AAAA",
//				Gateway:    net.ParseIP("127.0.0.0"),
//				SubStart:   net.ParseIP("192.168.0.89"),
//				SubEnd:     net.ParseIP("192.168.0.100"),
//			},
//			"CCCC": {
//				ID:         "CCCC",
//				CIDR:       "CCCC",
//				Gateway:    net.ParseIP("127.0.0.0"),
//				SubStart:   net.ParseIP("192.168.0.124"),
//				SubEnd:     net.ParseIP("192.168.0.145"),
//			},
//		},
//	}}
//	tsList.Items = []TenantSubnet{tsOne,tsTwo}
//
//	networkUsedRangeMap := make(map[string][]IPRange)
//	for _, v := range tsList.Items {
//		// need to judge this network has allocated to the tenant or not
//		// if tenant don't has this network, just skip and loop next
//		for networkId, subSpec := range v.Spec.Subnets {
//			if networkId == "" {
//				continue
//			}
//			start := subSpec.SubStart
//			end := subSpec.SubEnd
//			ipRange := IPRange{
//				Start: start,
//				End:   end,
//			}
//			_,ok := networkUsedRangeMap[networkId]
//			if !ok {
//				networkUsedRangeMap[networkId] = []IPRange{ipRange}
//			}else {
//				networkUsedRangeMap[networkId] = append(networkUsedRangeMap[networkId],ipRange)
//			}
//		}
//	}
//	total := len(networkUsedRangeMap)
//	var rangeList = new(IPRangeList)
//	rangeList.Items = make([]NetworkIPRange,0,total)
//	for networkId,v := range  networkUsedRangeMap{
//		if networkId == "" || v == nil {
//			continue
//		}
//		networkIPRange := NetworkIPRange{
//			Network:     networkId,
//			IPRangeList: v,
//		}
//		rangeList.Items = append(rangeList.Items,networkIPRange)
//	}
//	return rangeList, nil
//}
//
//func GetBitset(bitSet *bitset.BitSet) (string, uint, error) {
//	out, err := bitSet.MarshalJSON()
//	if err != nil {
//		return "", 0, err
//	}
//	// must use json unmarshal to a string
//	var str string
//	err = json.Unmarshal(out, &str)
//	if err != nil {
//		log.Errorf("unmarshal bitset to string error %v", err)
//		return "", 0, err
//	}
//	available := bitSet.Len() - bitSet.Count()
//	return str, available, nil
//}
//
//
//
//
//func ip4ToInt(ip net.IP) uint {
//	ip = ip.To4()
//	ipUint := uint(ip[3]) | (uint(ip[2]) << 8) | (uint(ip[1]) << 16) | (uint(ip[0]) << 24)
//	return ipUint
//}
//
//
//func newIPAllocator(cidr *net.IPNet, start net.IP, end net.IP, gateway net.IP) (*ipAllocator, error) {
//	var ia ipAllocator
//
//	prefix, _ := cidr.Mask.Size()
//	baseIP := ip4ToInt(cidr.IP.Mask(cidr.Mask))
//	maxSize := (1 << (32 - uint(prefix))) - 1
//
//	ia.prefix = uint(prefix)
//	ia.baseIP = baseIP
//	ia.bitSet.Set(uint(maxSize))
//	ia.bitSet.Set(uint(0))
//
//	ia.setBitsOutsideRange(start, end)
//
//	// make gateway unavailable
//	ia.AllocateSpecifiedIP(gateway)
//
//	log.Infof("ipallcator: %v/%d, min:%v, max:%v\n", intToIP4(ia.baseIP), ia.prefix, intToIP4(ia.min), intToIP4(ia.max))
//	return &ia, nil
//}
//
//func intToIP4(ipUint uint) net.IP {
//	b1, b2, b3, b4 := byte(ipUint>>24), byte(ipUint>>16), byte(ipUint>>8), byte(ipUint)
//	return net.IPv4(b1, b2, b3, b4)
//}
//
//func (ia *ipAllocator) setBitsOutsideRange(start net.IP, end net.IP) {
//
//	firstAddr, lastAddr := getFirstAndLastAddress(ia.baseIP, ia.prefix)
//
//	min := firstAddr
//	max := lastAddr
//
//	if start != nil {
//		minAddr := ip4ToInt(start)
//		if minAddr > firstAddr {
//			min = minAddr
//		}
//	}
//	if end != nil {
//		maxAddr := ip4ToInt(end)
//		if maxAddr < lastAddr {
//			max = maxAddr
//		}
//	}
//
//	ia.min = min
//	ia.max = max
//
//	for i := firstAddr; i < min; i++ {
//		if i > lastAddr {
//			break
//		}
//		index := i - ia.baseIP
//		ia.bitSet.Set(uint(index))
//	}
//
//	for j := lastAddr; j > max; j-- {
//		if j < firstAddr {
//			break
//		}
//		index := j - ia.baseIP
//		ia.bitSet.Set(uint(index))
//	}
//}
//
//
//func getFirstAndLastAddress(baseIP uint, prefix uint) (uint, uint) {
//	firstAddr := baseIP + 1
//	lastAddr := (baseIP | 0xffffffff>>prefix) - 1
//	return firstAddr, lastAddr
//}
//
//func (ia *ipAllocator) AllocateSpecifiedIP(ip net.IP) error {
//	if !ia.Contains(ip) {
//		return fmt.Errorf("specified ip %s out of range", ip)
//	}
//	v := ip4ToInt(ip)
//	v = v - ia.baseIP
//	if ia.bitSet.Test(uint(v)) {
//		return fmt.Errorf("specified ip %s has alreadly been allocated", ip)
//	}
//	ia.bitSet.Set(uint(v))
//	ia.used++
//	return nil
//}
//
//func (ia *ipAllocator) Contains(ip net.IP) bool {
//	if len(ip) == 0 {
//		return false
//	}
//	v := ip4ToInt(ip)
//	return v >= ia.min && v <= ia.max
//}
//func SplitID(id string) (string, string, string) {
//	ss := strings.Split(id, "/")
//	return ss[0], ss[1], ss[2]
//}
func intToIP4(ipUint uint) net.IP {
	b1, b2, b3, b4 := byte(ipUint>>24), byte(ipUint>>16), byte(ipUint>>8), byte(ipUint)
	return net.IPv4(b1, b2, b3, b4)
}

func ip4ToInt(ip net.IP) uint {
	ip = ip.To4()
	ipUint := uint(ip[3]) | (uint(ip[2]) << 8) | (uint(ip[1]) << 16) | (uint(ip[0]) << 24)
	return ipUint
}
func hexString(b []byte) string {
	s := make([]byte, len(b)*2)
	for i, tn := range b {
		s[i*2], s[i*2+1] = hexDigit[tn>>4], hexDigit[tn&0xf]
	}
	return string(s)
}
const hexDigit = "0123456789abcdef"
func Hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}
//  http://play.golang.org/p/m8TNTtygK0
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}


// Convert CIDR to IPv4 range
func CIDRRangeToIPv4Range(CIDR string) (ipStart string, ipEnd string, err error) {

	var ip uint32        // ip address

	var ipS uint32		 // Start IP address range
	var ipE uint32 		 // End IP address range

	cidrParts := strings.Split(CIDR, "/")

	ip = iPv4ToUint32(cidrParts[0])
	bits, _ := strconv.ParseUint(cidrParts[1], 10, 32)

	if ipS == 0 || ipS > ip {
		ipS = ip
	}

	ip = ip | (0xFFFFFFFF >> bits)

	if ipE < ip {
		ipE = ip
	}

	ipStart = uInt32ToIPv4(ipS)
	ipEnd = uInt32ToIPv4(ipE)

	return ipStart, ipEnd, err
}

//Convert IPv4 to uint32
func iPv4ToUint32(iPv4 string ) uint32 {

	ipOctets := [4]uint64{}

	for i, v := range strings.SplitN(iPv4,".", 4) {
		ipOctets[i], _  = strconv.ParseUint(v, 10, 32)
	}

	result := (ipOctets[0] << 24) | (ipOctets[1] << 16) | (ipOctets[2] << 8) | ipOctets[3]

	return uint32(result)
}

//Convert uint32 to IP
func uInt32ToIPv4(iPuInt32 uint32) (iP string) {
	iP =  fmt.Sprintf ("%d.%d.%d.%d",
		iPuInt32 >> 24,
		(iPuInt32 & 0x00FFFFFF)>> 16,
		(iPuInt32 & 0x0000FFFF) >> 8,
		iPuInt32 & 0x000000FF)
	return iP
}


func main()  {
	start,end,_ := CIDRRangeToIPv4Range("1.5.4.7/13")
	fmt.Println("start :",start)
	fmt.Println("end :",end)


	aa,_ :=Hosts("1.0.0.0/8")
	length := len(aa)
	fmt.Println(aa[0])
	fmt.Println(aa[length-1])
	cidr, _ := types.ParseCIDR("127.27.0.5/23")
	fmt.Println(cidr.IP.String())
	fmt.Println(hexString(cidr.Mask))
	//ip4ToInt(cidr.IP)
	//
	//subStart := cidr
	//subEnd := cidr.IP.Mask(cidr.Mask)

}
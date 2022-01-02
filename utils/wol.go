package utils

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"strings"
)

func GetMagicPacket(mac string) ([]byte, error) {

	// 替换掉 MAC 地址中的 - 或者 :
	hardwareAddr := strings.Replace(strings.Replace(mac, ":", "", -1), "-", "", -1)
	if len(hardwareAddr) != 12 {
		return nil, errors.New(fmt.Sprintf("MAC %s 格式不正确", mac))
	}

	// hex.DecodeString
	macHex, err := hex.DecodeString(hardwareAddr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("MAC %s 格式不正确", mac))
	}

	// Broadcast FF:FF:FF:FF:FF:FF
	var broadcast = []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
	var buff bytes.Buffer
	buff.Write(broadcast)
	for i := 0; i < 16; i++ {
		buff.Write(macHex)
	}

	// 获得 MagicPacket
	magicPacket := buff.Bytes()
	if len(magicPacket) != 102 {
		return nil, errors.New(fmt.Sprintf("MAC %s 格式不正确", mac))
	}

	return magicPacket, nil

}

func SendMagicPacket(magicPacket []byte, nicName string) (bool, error) {

	sender := net.UDPAddr{}
	if len(nicName) != 0 {
		address, err := getNicIpv4AddressByName(nicName)
		if err != nil {
			return false, errors.New(fmt.Sprintf("网卡 %s 错误: %v", nicName, err))
		}
		sender.IP = address
	}

	target := net.UDPAddr{
		IP:   net.IPv4bcast,
		Port: 9,
	}
	conn, err := net.DialUDP("udp", &sender, &target)
	if err != nil {
		return false, errors.New(fmt.Sprintf("UDP 创建失败: %v", err))
	}
	defer func() {
		_ = conn.Close()
	}()

	_, err = conn.Write(magicPacket)
	if err != nil {
		return false, errors.New(fmt.Sprintf("MagicPacket 发送失败: %v", err))
	}
	return true, nil

}

func getNicIpv4AddressByName(nicName string) (net.IP, error) {

	nic, err := net.InterfaceByName(nicName)
	if err != nil {
		return nil, err
	}

	if (nic.Flags & net.FlagUp) == 0 {
		return nil, errors.New(fmt.Sprintf("网卡 %s 未处于工作状态", nicName))
	}

	address, err := nic.Addrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range address {
		if ip, ok := addr.(*net.IPNet); ok {
			if ipv4 := ip.IP.To4(); ipv4 != nil {
				return ipv4, nil
			}
		}
	}

	return nil, errors.New(fmt.Sprintf("网卡 %s IPV4 地址获取失败", nicName))

}

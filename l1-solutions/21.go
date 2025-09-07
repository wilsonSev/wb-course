package main

import (
	"fmt"
	"time"
)

// Интерфейс для объектов с USB-C разъемами
type UsbDevice interface {
	ChargeWithUsbCable()
}

// Смартфон Samsung реализует данный интерфейс, поскольку обладает необходимым разъемом
type Samsung struct {
	battery int
}

func (s *Samsung) ChargeWithUsbCable() {
	time.Sleep(1 * time.Second)
	s.battery = 100
	fmt.Println("Samsung has been charged!")
}

// Смартфон Iphone (до 14 модели включительно) имеет другой тип разъема, значит не реализует интерфейс
type Iphone struct {
	battery int
}

func (i *Iphone) ChargeWithLightningCable() {
	time.Sleep(1 * time.Second)
	i.battery = 100
	fmt.Println("Iphone has been charged!")
}

// Кабель, подключенный к розетке, который умеет заряжать устройства с разъемом USB
type UsbCable struct{}

func (uc *UsbCable) Charge(device UsbDevice) {
	device.ChargeWithUsbCable()
}

type UsbAdapter struct {
	iphone *Iphone
}

func NewUsbAdapter(iphone *Iphone) *UsbAdapter {
	return &UsbAdapter{iphone}
}

func (ua *UsbAdapter) ChargeWithUsbCable() {
	ua.iphone.ChargeWithLightningCable()
}

func main() {
	cable := UsbCable{}
	samsung := new(Samsung)
	iphone := new(Iphone)
	adapter := NewUsbAdapter(iphone)
	cable.Charge(samsung)
	// cable.Charge(iphone) эта запись вызовет ошибку, поскольку айфон не реализует интерфейс USB-устройств
	cable.Charge(adapter)
}

package main

type IphoneFactory struct{}

func (i *IphoneFactory) Create() {
	i.createCamera()
	i.createChip()
}
func (i *IphoneFactory) createCamera() {

}
func (i *IphoneFactory) createChip() {

}

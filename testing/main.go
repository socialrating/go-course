package main

import "fmt"

type Player interface {
	Fly()
	RunOnPet()
	UsePetSkills()
	UseSkillsEllam()
	Run()
	Swim()
	UseSkills()
	UseInventory()
	UseGlidersSkill()
	Chat(message string)
	ChangeChat(name string)
	AddFriend(name string)
	BlockUser(name string)
	InviteGuild(name string)
}

type PlayerImpl struct {
	Name string
}

func (p *PlayerImpl) Fly() {
	fmt.Println("Flying", p.Name)
}

func (p *PlayerImpl) RunOnPet() {
	fmt.Println("Running on pet", p.Name)
}

func (p *PlayerImpl) UsePetSkills() {
	fmt.Println("Using pet skills", p.Name)
}

func (p *PlayerImpl) UseSkillsEllam() {
	fmt.Println("Using SkillsEllam", p.Name)
}

func (p *PlayerImpl) Run() {
	fmt.Println("Running", p.Name)
}

func (p *PlayerImpl) Swim() {
	fmt.Println("Swimming", p.Name)
}

func (p *PlayerImpl) UseSkills() {
	fmt.Println("Using skills", p.Name)
}

func (p *PlayerImpl) UseInventory() {
	fmt.Println("Using inventory", p.Name)
}

func (p *PlayerImpl) UseGlidersSkill() {
	fmt.Println("Using gliders skill", p.Name)
}

func (p *PlayerImpl) Chat(message string) {
	fmt.Println("Chat:", message)
}

func (p *PlayerImpl) ChangeChat(name string) {
	fmt.Println("Changed chat:", name)
}

func (p *PlayerImpl) AddFriend(name string) {
	fmt.Println("Added friend:", name)
}

func (p *PlayerImpl) BlockUser(name string) {
	fmt.Println("Blocked user:", name)
}

func (p *PlayerImpl) InviteGuild(name string) {
	fmt.Println("Invited to guild:", name)
}

func main() {
	player := &PlayerImpl{Name: "Игрок"}

	var chatMessage string
	fmt.Print("Что вы хотите написать в чат? ")
	fmt.Scanln(&chatMessage)
	player.Chat(chatMessage)

	var friendName string
	fmt.Print("Кого вы хотите добавить в друзья: ")
	fmt.Scanln(&friendName)
	player.AddFriend(friendName)
}

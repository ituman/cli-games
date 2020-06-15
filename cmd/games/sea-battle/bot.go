package seabattle

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Bot struct {
	name     string
	field    *Field
	fleet    Fleet
	enemy    Player
	attempts []string
}

func (b *Bot) initAttempts() {
	for i := 10; i < 100; i++ {
		b.attempts = append(b.attempts, strconv.Itoa(i))
	}
}

func (b *Bot) initField() {
	b.initAttempts()
	defaultPos := [][]int{
		{0, 0},
		{0, 2},
		{0, 4},
		{0, 6},
		{0, 8},
		{5, 0},
		{5, 2},
		{5, 4},
		{5, 6},
		{5, 8},
	}
	for i, s := range b.fleet {
		b.field.SetShip(defaultPos[i][0], defaultPos[i][1], true, s)
	}
}

func (b *Bot) shot() bool {
	time.Sleep(2 * time.Second)
	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(b.attempts))
	nums := strings.Split(b.attempts[i], "")
	b.attempts = append(b.attempts[:i], b.attempts[i+1:]...) // Убираем использованную попытку
	x, _ := strconv.Atoi(nums[0])
	y, _ := strconv.Atoi(nums[1])
	fmt.Println("Стреляет: ", x, y)
	isHit := b.getEnemy().getField().Shoot(x, y)
	return isHit
}

func (b *Bot) isLose() bool {
	if b.fleet.StillAlive() > 0 {
		return false
	} else {
		return true
	}
}

func (b *Bot) getEnemy() Player {
	return b.enemy
}

func (b *Bot) getField() *Field {
	return b.field
}

func (b *Bot) getFleet() Fleet {
	return b.fleet
}

func (b *Bot) getName() string {
	return b.name
}

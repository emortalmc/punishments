package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type PunishmentLadder struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Punishments []Punishment `json:"punishments"`
}

type PunishmentType string

const (
	PunishmentTypeKick PunishmentType = "kick"
	PunishmentTypeBan  PunishmentType = "ban"
	PunishmentTypeMute PunishmentType = "mute"
)

type Punishment struct {
	Type PunishmentType `json:"type"`
}

type DurationPunishment struct {
	Punishment
	Duration time.Duration `json:"duration"`
}

type PunishmentConfig struct {
	Ladders map[string]PunishmentLadder `json:"ladders"`
}

func LoadPunishmentConfig() (config *PunishmentConfig, err error) {
	file, err := os.Open("punishments.json")
	if err != nil {
		return nil, fmt.Errorf("failed to open punishments file: %s", err)
	}

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("faile to read punishments file: %s", err)
	}

	err = json.Unmarshal(b, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal punishments JSON: %s", err)
	}

	if err := file.Close(); err != nil {
		return nil, fmt.Errorf("failed to close punishments file: %s", err)
	}

	return
}

func (c *PunishmentConfig) GetPunishmentFromLadder(ladderId string, index int) (punishment Punishment, ok bool) {
	ladder, ok := c.Ladders[ladderId]
	if !ok {
		return
	}
	punishment = ladder.Punishments[index]
	return
}

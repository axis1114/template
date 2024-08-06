package config

type System struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Env       string `mapstructure:"env"`
	StartTime string `mapstructure:"start_time"`
	MachineID uint64 `mapstructure:"machine_id"`
}

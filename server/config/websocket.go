package config

type Websocket struct {
	Url             string `mapstructure:"url" json:"url" yaml:"url"`
	Register        string `mapstructure:"register" json:"register" yaml:"register"`
	Send_to_client  string `mapstructure:"send_to_client" json:"send_to_client" yaml:"send_to_client"`
	Send_to_clients string `mapstructure:"send_to_clients" json:"send_to_clients" yaml:"send_to_clients"`
	Send_to_group   string `mapstructure:"send_to_group" json:"send_to_group" yaml:"send_to_group"`
	Bind_to_group   string `mapstructure:"bind_to_group" json:"bind_to_group" yaml:"bind_to_group"`
	Get_online_list string `mapstructure:"get_online_list" json:"get_online_list" yaml:"get_online_list"`
	Close_client    string `mapstructure:"close_client" json:"close_client" yaml:"close_client"`
}

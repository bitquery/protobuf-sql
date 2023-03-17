package generators

import "testing"

func Test_genFileName(t *testing.T) {
	type args struct {
		messageName string
		suffix      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"first", args{"TestMessage", "Message"}, "tests.sql"},
		{"second", args{"TransactionMessage", "Message"}, "transactions.sql"},
		{"third", args{"BalanceUpdateMessage", "Message"}, "balance_updates.sql"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileName(tt.args.messageName, tt.args.suffix); got != tt.want {
				t.Errorf("genFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

package hashutils

import "testing"

func TestBase64(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				input: "my secret key",
			},
			want: "bXkgc2VjcmV0IGtleQ==",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64(tt.args.input); got != tt.want {
				t.Errorf("Base64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecret2Password(t *testing.T) {
	type args struct {
		username string
		secret   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				username: "wubin",
				secret:   "my secret",
			},
			want: "cff23c519b29a0e0c0304ff1a3d795f171b9c919",
		},
		{
			name: "",
			args: args{
				username: "wubin",
				secret:   "",
			},
			want: "f85610573ac9cda1a0e27e27406e9125e0e2403d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Secret2Password(tt.args.username, tt.args.secret); got != tt.want {
				t.Errorf("Secret2Password() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUIDByString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				input: "http://219.142.101.111/gongkai/fdzdgknr/zfhcxjsbwj/202110/20211013_762458.html",
			},
			want: "bddcb874-43f2-5fee-aeab-b0a6c4e5103b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UUIDByString(tt.args.input); got != tt.want {
				t.Errorf("UUIDByString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package domain

import (
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
)

// weaver.InstanceOf checks.

// weaver.Router checks.

// Local stub implementations.

// Client stub implementations.

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][24]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.24.5 (codegen
version v0.24.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

// Server stub implementations.

// Reflect stub implementations.

// AutoMarshal implementations.

var _ codegen.AutoMarshal = (*Config)(nil)

type __is_Config[T ~struct {
	weaver.AutoMarshal
	Supabase       SupabaseConfig "json:\"supabase\""
	DefaultAccount DefaultAccount "json:\"defaultAccount\""
}] struct{}

var _ __is_Config[Config]

func (x *Config) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Config.WeaverMarshal: nil receiver"))
	}
	(x.Supabase).WeaverMarshal(enc)
	(x.DefaultAccount).WeaverMarshal(enc)
}

func (x *Config) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Config.WeaverUnmarshal: nil receiver"))
	}
	(&x.Supabase).WeaverUnmarshal(dec)
	(&x.DefaultAccount).WeaverUnmarshal(dec)
}

var _ codegen.AutoMarshal = (*DefaultAccount)(nil)

type __is_DefaultAccount[T ~struct {
	weaver.AutoMarshal
	Email    string "json:\"email\""
	Password string "json:\"password\""
}] struct{}

var _ __is_DefaultAccount[DefaultAccount]

func (x *DefaultAccount) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("DefaultAccount.WeaverMarshal: nil receiver"))
	}
	enc.String(x.Email)
	enc.String(x.Password)
}

func (x *DefaultAccount) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("DefaultAccount.WeaverUnmarshal: nil receiver"))
	}
	x.Email = dec.String()
	x.Password = dec.String()
}

var _ codegen.AutoMarshal = (*SupabaseConfig)(nil)

type __is_SupabaseConfig[T ~struct {
	weaver.AutoMarshal
	Postgres   string "json:\"postgres\""
	Api        string "json:\"api\""
	JwtSecret  string "json:\"jwtSecret\""
	AnonKey    string "json:\"anonKey\""
	ServiceKey string "json:\"serviceKey\""
	S3Access   string "json:\"s3Access\""
	S3Secret   string "json:\"s3Secret\""
	S3Region   string "json:\"s3Region\""
}] struct{}

var _ __is_SupabaseConfig[SupabaseConfig]

func (x *SupabaseConfig) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("SupabaseConfig.WeaverMarshal: nil receiver"))
	}
	enc.String(x.Postgres)
	enc.String(x.Api)
	enc.String(x.JwtSecret)
	enc.String(x.AnonKey)
	enc.String(x.ServiceKey)
	enc.String(x.S3Access)
	enc.String(x.S3Secret)
	enc.String(x.S3Region)
}

func (x *SupabaseConfig) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("SupabaseConfig.WeaverUnmarshal: nil receiver"))
	}
	x.Postgres = dec.String()
	x.Api = dec.String()
	x.JwtSecret = dec.String()
	x.AnonKey = dec.String()
	x.ServiceKey = dec.String()
	x.S3Access = dec.String()
	x.S3Secret = dec.String()
	x.S3Region = dec.String()
}

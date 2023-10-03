package interfaces

import "reflect"

type Options struct {
	// ShouldInclude tests whether a method should be included in the generated interfaces. If it returns true,
	// the method will be included. MethodHasPrefix and MethodHasSuffix can be used inside a custom function here
	// to customize the behavior.
	ShouldInclude func(reflect.Method) bool

	// ExtraImports can add extra imports for a method
	ExtraImports func(reflect.Method) []string

	// SinglePackage allows to generate all passed clients into a single package.
	// The clients will get their package name as prefix to the interface name (e.g., s3.Client -> S3Client)
	SinglePackage string
}

func (o *Options) SetDefaults() {
	if o.ShouldInclude == nil {
		o.ShouldInclude = func(reflect.Method) bool { return true }
	}
	if o.ExtraImports == nil {
		o.ExtraImports = func(reflect.Method) []string { return nil }
	}
}

type Option func(*Options)

func WithIncludeFunc(f func(reflect.Method) bool) Option {
	return func(o *Options) {
		o.ShouldInclude = f
	}
}

func WithExtraImports(f func(reflect.Method) []string) Option {
	return func(o *Options) {
		o.ExtraImports = f
	}
}

func WithSinglePackage(name string) Option {
	return func(o *Options) {
		o.SinglePackage = name
	}
}

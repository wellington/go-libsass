package libs

func SetIncludePaths(goopts SassOptions, paths []string) {
	for _, inc := range paths {
		SassOptionSetIncludePath(goopts, inc)
	}
}

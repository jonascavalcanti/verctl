package help

func Default() string {
	defaultMsg := "xVersioner (CLI): xversioner controls application version\n" +
		"\n" +
		"Find more information at:  \n" +
		"\n" +
		"Basic Commands:\n" +
		"\tget		get application version \n" +
		"\tupdate	Update application version \n"

	return defaultMsg
}

func Update() string {
	updateMsg := "xVersioner (CLI): xversioner controls application version\n" +
		"update command: \n" +
		"\n" +
		"Examples:\n" +
		"\t# Generate patch version of the application\n" +
		"\txversioner update filepath ./app/file_contain_app_version type patch\n" +
		"\n" +
		"\txversioner update -f ./app/file_contain_app_version -t patch\n" +
		"Find more information at:  \n" +
		"\n" +
		"Basic Commands:\n" +
		"\tfilepath,	-f	path of the file where contains application version\n" +
		"\ttype, 		-t	type of version that will incremented (major, minor, patch) \n"

	return updateMsg
}

func Get() string {
	getMsg := "xVersioner (CLI): xversioner controls application version\n" +
		"get command: \n" +
		"\n" +
		"Examples:\n" +
		"\t# Generate patch version of the application\n" +
		"\txversioner get filepath ./app/file_contain_app_version \n" +
		"\t or\n" +
		"\txversioner get -f ./app/file_contain_app_version \n" +
		"Find more information at:  \n" +
		"\n" +
		"Basic Commands:\n" +
		"\tfilepath,	-f	path of the file where contains application version\n"

	return getMsg
}

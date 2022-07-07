package help

func Default() string {
	defaultMsg := "xversioner manages version control of applications based on the `Semantica Version Specification` and `Commit Convention`\n" +
		"\n" +
		"Find more information at:  \n" +
		"\n" +
		"Basic Commands:\n" +
		"\tget		get application version \n" +
		"\tupdate	Update application version \n"

	return defaultMsg
}

func Update() string {
	updateMsg := "xversioner manages version control of applications based on the `Semantica Version Specification` and `Commit Convention`\n" +
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
		"\tincrement, 	-i	type of version that will incremented (major, minor, patch) \n"

	return updateMsg
}

func Get() string {
	getMsg := "xversioner manages version control of applications based on the `Semantica Version Specification` and `Commit Convention`\n" +
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
		"\tfilepath, -f		path of the file where contains application version\n"

	return getMsg
}

package e

import "errors"

var msgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	ERROR_PARSING_CONFIG_FILE:                     "Config file could not be parsed",
	ERROR_GETTING_CURRENT_DIR:                     "Current directory could not get",
	ERROR_CREATING_KUBECONFIG:                     "Config could not be created from kubeconfig file",
	ERROR_CREATING_CLIENTSET:                      "Clientset could not be created",
	ERROR_CREATING_NAMESPACE:                      "Namespace could not be created",
	ERROR_CONNECTING_DATABASE:                     "Database could not be connected",
	ERROR_SETUP_DATABASE:                          "Database could not be setup",
	ERROR_CREATING_PROJECT_WRONG_INPUT:            "Project could not be created (Wrong input)",
	ERROR_GETTING_PROJECTS:                        "Projects could not be retrieved from database",
	ERROR_GETTING_PROJECT_NOT_INTEGER_PROJECT_ID:  "Project could not be retrieved (Project ID must be an integer)",
	ERROR_GETTING_PROJECT:                         "Project could not be retrieved from database",
	ERROR_CREATING_ARTICLE_NOT_INTEGER_PROJECT_ID: "Article could not be created (Project ID must be an integer)",
	ERROR_CREATING_ARTICLE_NOT_EXISTS_PROJECT:     "Article could not be created (Project must exists)",
	ERROR_CREATING_ARTICLE_WRONG_INPUT:            "Article could not be created (Wrong input)",
	ERROR_GETTING_ARTICLES_NOT_INTEGER_PROJECT_ID: "Articles could not be retrieved (Project ID must be an integer)",
	ERROR_GETTING_ARTICLES:                        "Articles could not be retrieved from database",
	ERROR_GETTING_ARTICLE_NOT_INTEGER_PROJECT_ID:  "Article could not be retrieved (Project ID must be an integer)",
	ERROR_GETTING_ARTICLE_NOT_INTEGER_ARTICLE_ID:  "Article could not be retrieved (Article ID must be an integer)",
	ERROR_GETTING_ARTICLE:                         "Article could not be retrieved from database",
}

func GetMsg(code int) string {
	msg, ok := msgFlags[code]
	if ok {
		return msg
	}

	return msgFlags[ERROR]
}

func Error(code int) error {
	return errors.New(GetMsg(code))
}

package main

import (
	"os"

	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/recipe/passwordless/plessmodels"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartypasswordless"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartypasswordless/tplmodels"
	"github.com/supertokens/supertokens-golang/supertokens"

	_ "github.com/joho/godotenv/autoload"
)

func createSuperTokensConfig() supertokens.TypeInput {
	var superTokensConnectionURI = os.Getenv("SUPERTOKENS_CONNECTION_URI")
	var superTokensAPIKey = os.Getenv("SUPERTOKENS_API_KEY")

	var superTokensAPIDomain = os.Getenv("AUTH_API_DOMAIN")
	var superTokensUIDomain = os.Getenv("AUTH_UI_DOMAIN")
	
	var SuperTokensConfig = supertokens.TypeInput{
		
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: superTokensConnectionURI,
			APIKey: superTokensAPIKey,
		},
		AppInfo: supertokens.AppInfo{
			AppName:       "Scott Benton Auth",
			APIDomain:     superTokensAPIDomain,
			WebsiteDomain: superTokensUIDomain,
		},
		RecipeList: []supertokens.Recipe{
			thirdpartypasswordless.Init(tplmodels.TypeInput{
				FlowType: "USER_INPUT_CODE_AND_MAGIC_LINK",
				ContactMethodEmailOrPhone: plessmodels.ContactMethodEmailOrPhoneConfig{
					Enabled: true,
				},
				Providers: []tpmodels.ProviderInput{},
			}),
			session.Init(nil), // initializes session features
			dashboard.Init(nil),
		},
	}
	
	return SuperTokensConfig;
}
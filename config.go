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
	var superTokensUIPath = os.Getenv("AUTH_UI_PATH")

	var googleOAuthClientId = os.Getenv("GOOGLE_OAUTH_CLIENT_ID");
	var googleOAuthClientSecret = os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET");
	
	var githubOAuthClientId = os.Getenv("GITHUB_OAUTH_CLIENT_ID");
	var githubOAuthClientSecret = os.Getenv("GITHUB_OAUTH_CLIENT_SECRET");

	var discordOAuthClientId = os.Getenv("DISCORD_OAUTH_CLIENT_ID");
	var discordOAuthClientSecret = os.Getenv("DISCORD_OAUTH_CLIENT_SECRET");
	
	var SuperTokensConfig = supertokens.TypeInput{
		
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: superTokensConnectionURI,
			APIKey: superTokensAPIKey,
		},
		AppInfo: supertokens.AppInfo{
			AppName:       "Scott Benton Auth",
			APIDomain:     superTokensAPIDomain,
			WebsiteDomain: superTokensUIDomain,
			WebsiteBasePath: &superTokensUIPath,
		},
		RecipeList: []supertokens.Recipe{
			thirdpartypasswordless.Init(tplmodels.TypeInput{
				FlowType: "USER_INPUT_CODE_AND_MAGIC_LINK",
				ContactMethodEmailOrPhone: plessmodels.ContactMethodEmailOrPhoneConfig{
					Enabled: true,
				},
				Providers: []tpmodels.ProviderInput{
					{
						Config: tpmodels.ProviderConfig{
							ThirdPartyId: "google",
							Clients: []tpmodels.ProviderClientConfig{
								{
									ClientID: googleOAuthClientId,
									ClientSecret: googleOAuthClientSecret,
								},
							},
						},
					},
					{
						Config: tpmodels.ProviderConfig{
							ThirdPartyId: "github",
							Clients: []tpmodels.ProviderClientConfig{
								{
									ClientID: githubOAuthClientId,
									ClientSecret: githubOAuthClientSecret,
								},
							},
						},
					},
					{
						Config: tpmodels.ProviderConfig{
							ThirdPartyId: "discord",
							Clients: []tpmodels.ProviderClientConfig{
								{
									ClientID: discordOAuthClientId,
									ClientSecret: discordOAuthClientSecret,
								},
							},
						},
					},
				},
			}),
			session.Init(nil), // initializes session features
			dashboard.Init(nil),
		},
	}
	
	return SuperTokensConfig;
}
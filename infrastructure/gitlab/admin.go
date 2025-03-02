package gitlab

import (
	"strconv"
	"strings"

	sdk "github.com/xanzy/go-gitlab"

	"github.com/opensourceways/xihe-server/domain/platform"
	userdomain "github.com/opensourceways/xihe-server/user/domain"
)

var (
	admin     *administrator
	obsHelper *obsService

	endpoint        string
	defaultBranch   string
	graphqlEndpoint string

	maxFileCount int
)

func NewUserSerivce() platform.User {
	return admin
}

func Init(cfg *Config) error {
	v, err := sdk.NewClient(cfg.RootToken, sdk.WithBaseURL(cfg.Endpoint))
	if err != nil {
		return err
	}

	s, err := initOBS(cfg)
	if err != nil {
		return err
	}

	obsHelper = &s

	admin = &administrator{v}
	endpoint = strings.TrimSuffix(cfg.Endpoint, "/")
	maxFileCount = cfg.MaxFileCount
	defaultBranch = cfg.DefaultBranch
	graphqlEndpoint = strings.TrimSuffix(cfg.GraphqlEndpoint, "/")

	return nil
}

type administrator struct {
	cli *sdk.Client
}

func (m *administrator) New(u platform.UserOption) (r userdomain.PlatformUser, err error) {
	name := u.Name.Account()
	email := u.Email.Email()
	pass := u.Password.Password()
	b := true

	v, _, err := m.cli.Users.CreateUser(&sdk.CreateUserOptions{
		Name:             &name,
		Email:            &email,
		Username:         &name,
		Password:         &pass,
		SkipConfirmation: &b,
	})

	if err != nil {
		return
	}

	r.Id = strconv.Itoa(v.ID)
	r.NamespaceId = strconv.Itoa(v.NamespaceID)

	return
}

func (m *administrator) NewToken(u userdomain.PlatformUser) (string, error) {
	uid, err := strconv.Atoi(u.Id)
	if err != nil {
		return "", err
	}

	name := "___"
	scope := []string{"api"}

	v, _, err := m.cli.Users.CreatePersonalAccessToken(
		uid, &sdk.CreatePersonalAccessTokenOptions{
			Name:   &name,
			Scopes: &scope,
		},
	)

	if err != nil {
		return "", err
	}

	return v.Token, nil
}

func (m *administrator) RefreshToken(userId string) (string, error) {
	uid, err := strconv.Atoi(userId)
	if err != nil {
		return "", err
	}

	opt := sdk.ListPersonalAccessTokensOptions{
		UserID: &uid,
	}
	tokens, _, err := m.cli.PersonalAccessTokens.ListPersonalAccessTokens(&opt)
	if err != nil {
		return "", nil
	}

	for _, token := range tokens {
		if token.Active {
			_, err := m.cli.PersonalAccessTokens.RevokePersonalAccessToken(token.ID)
			if err != nil {
				return "", err
			}
		}

	}

	name := "___"
	scope := []string{"api"}

	v, _, err := m.cli.Users.CreatePersonalAccessToken(
		uid, &sdk.CreatePersonalAccessTokenOptions{
			Name:   &name,
			Scopes: &scope,
		},
	)

	if err != nil {
		return "", err
	}

	return v.Token, nil
}

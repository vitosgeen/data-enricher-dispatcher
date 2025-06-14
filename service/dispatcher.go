package service

import (
	"fmt"

	"data-enricher-dispatcher/apperrors"
	"data-enricher-dispatcher/client"
	"data-enricher-dispatcher/config"
	"data-enricher-dispatcher/logger"
	"data-enricher-dispatcher/model"
)

const infoSkipping = "skipping user with email: %s due to special postfix exclusion"

type Dispatcher interface {
	Start() error
}

type dispatcher struct {
	apiClient client.APIClient
	logger    logger.Logger
	cfg       *config.Config
}

func NewDispatcher(apiClient client.APIClient, logger logger.Logger, cfg *config.Config) Dispatcher {
	return &dispatcher{
		apiClient: apiClient,
		logger:    logger,
		cfg:       cfg,
	}
}

func (d *dispatcher) Start() error {
	users, err := d.apiClient.GetUsers()
	if err != nil {
		return apperrors.ServiceDispatcherGetUsersError.AppendMessage(err)
	}
	for _, user := range users {
		if !model.UserEmailHasSpecialPostfix(&user, d.cfg.ExcludePostfixes) {
			err := fmt.Errorf(infoSkipping, user.Email)
			fmt.Println(err.Error())
			continue
		}
		if !user.IsValid() {
			d.logger.Println(apperrors.ServiceDispatcherInvalidUserError.AppendMessage(user))
			continue
		}

		err := d.apiClient.PostUser(user)
		if err != nil {
			d.logger.Error(apperrors.ServiceDispatcherPostUserError.AppendMessage(err, user))
		}
	}

	return nil
}

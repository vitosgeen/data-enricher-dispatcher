package service

import (
	"context"
	"fmt"
	"time"

	"data-enricher-dispatcher/apperrors"
	"data-enricher-dispatcher/client"
	"data-enricher-dispatcher/config"
	"data-enricher-dispatcher/logger"
	"data-enricher-dispatcher/model"
)

const (
	defaultTimeout = 10 * time.Second
	infoSkipping   = "skipping user with email: %s due to special postfix exclusion"
)

type Dispatcher interface {
	Start(ctx context.Context) error
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

func (d *dispatcher) Start(ctx context.Context) error {
	userCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	users, err := d.apiClient.GetUsers(userCtx)
	if err != nil {
		return apperrors.ServiceDispatcherGetUsersError.AppendMessage(err)
	}
	for _, user := range users {
		if !model.UserEmailHasSpecialPostfix(&user, d.cfg.ExcludePostfixes) {
			err := fmt.Errorf(infoSkipping, user.Email)
			d.logger.Info(err.Error())
			// show to the console
			fmt.Println(err.Error())
			continue
		}
		if !user.IsValid() {
			d.logger.Println(apperrors.ServiceDispatcherInvalidUserError.AppendMessage(user))
			continue
		}

		postUserCtx, postCancel := context.WithTimeout(ctx, defaultTimeout)
		defer postCancel()
		err := d.apiClient.PostUser(postUserCtx, user)
		if err != nil {
			d.logger.Error(apperrors.ServiceDispatcherPostUserError.AppendMessage(err, user))
		}
	}

	return nil
}

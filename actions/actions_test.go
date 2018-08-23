package actions

import (
	"testing"
	"time"

	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/suite"
	"github.com/robbyoconnor/golangnyc_cfp_aug23/models"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	action, err := suite.NewActionWithFixtures(App(), packr.NewBox("../fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	as := &ActionSuite{
		Action: action,
	}
	suite.Run(t, as)
}

func (as *ActionSuite) CreateUser() *models.User {
	user := &models.User{
		Name:       "Mark",
		Email:      nulls.NewString("mark@example.com"),
		Provider:   "faux",
		ProviderID: time.Now().String(),
	}
	as.NoError(as.DB.Create(user))
	return user
}

func (as *ActionSuite) Login() *models.User {
	user := as.CreateUser()
	as.Session.Set("current_user_id", user.ID)
	return user
}

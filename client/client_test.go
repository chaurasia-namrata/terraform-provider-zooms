package client

import(
	"github.com/stretchr/testify/assert"
	"testing"
	"os"
	"log"
)

func init(){
	os.Setenv("ZOOM_TOKEN", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MjM0NzQ4MDksImlhdCI6MTYyMjg3MDAwOX0.m3I7dQYBh96qoiKBfG7Vvw1v5G1_4-RUZOWVIs5MrU4")
}

func TestClient_GetItem(t *testing.T) {
	testCases := []struct {
		testName     string
		itemName     string
		seedData     map[string]User
		expectErr    bool
		expectedResp *User
	}{
		{
			testName: "user exists",
			itemName: "ui17co14@iiitsurat.ac.in",
			seedData: map[string]User{
				"ekansh0786@gmail.com": {
					Id:        "t2OUx6lvTMedrAiW2ffURA",
					Email:   "ui17co14@iiitsurat.ac.in",
					FirstName: "ekansh",
					LastName:  "singh",
					Type:        1,
					RoleName:    "Member",
					Status  :    "active",
					Department:  "devops",
					JobTitle:    "Engineer",
					Location:	"Delhi",
				},
			},
			expectErr: false,
			expectedResp: &User{
				Id:        "t2OUx6lvTMedrAiW2ffURA",
				Email:   "ui17co14@iiitsurat.ac.in",
				FirstName: "ekansh",
				LastName:  "singh",
				Type:        1,
				RoleName:    "Member",
				Status  :    "active",
				Department:  "devops",
				JobTitle:    "Engineer",
				Location:	"Delhi",
			},
		},
		
		{
			testName:     "user does not exist",
			itemName:     "ui17co14@itsurat.ac.in",
			seedData:     nil,
			expectErr:    true,
			expectedResp: nil,
		},
		
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("ZOOM_TOKEN"))

			item, err := client.GetItem(tc.itemName)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, item)
		})
	}
}

func TestClient_NewItem(t *testing.T) {
	testCases := []struct {
		testName  string
		newItem   *User
		seedData  map[string]User
		expectErr bool
	}{
		{
			testName: "success",
			newItem: &User{
				Id:        "t2OUx6lvTMedrAiW2ffURA",
				Email:   "ui17co14@iiitsurat.ac.in",
				FirstName: "ekansh",
				LastName:  "singh",
				Type:        1,
				RoleName:    "Member",
				Status  :    "active",
				Department:  "devops",
				JobTitle:    "Engineer",
				Location:	"Delhi",
			},
			seedData:  nil,
			expectErr: false,
		},
		{
			testName: "item already exists",
			newItem: &User{
				Id:        "t2OUx6lvTMedrAiW2ffURA",
				Email:   "ui17co14@iiitsurat.ac.in",
				FirstName: "ekansh",
				LastName:  "singh",
				Type:        1,
				RoleName:    "Member",
				Status  :    "active",
				Department:  "devops",
				JobTitle:    "Engineer",
				Location:	"Delhi",
			},
			seedData: map[string]User{
				"item1": {
					Id:        "t2OUx6lvTMedrAiW2ffURA",
					Email:   "ui17co14@iiitsurat.ac.in",
					FirstName: "ekansh",
					LastName:  "singh",
					Type:        1,
					RoleName:    "Member",
					Status  :    "active",
					Department:  "devops",
					JobTitle:    "Engineer",
					Location:	"Delhi",
				},
			},
			expectErr: true,
		},
		
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("ZOOM_TOKEN"))


			err := client.NewItem(tc.newItem)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			item, err := client.GetItem(tc.newItem.Email)
			assert.NoError(t, err)
			assert.Equal(t, tc.newItem, item)
		})
	}
}

func TestClient_UpdateItem(t *testing.T) {
	testCases := []struct {
		testName    string
		updatedItem *User
		seedData    map[string]User
		expectErr   bool
	}{
		{
			testName: "item exists",
			updatedItem: &User{
				Id:        "t2OUx6lvTMedrAiW2ffURA",
					Email:   "ui17co14@iiitsurat.ac.in",
					FirstName: "ekansh",
					LastName:  "singh",
					Type:        1,
					RoleName:    "Member",
					Status  :    "active",
					Department:  "devops",
					JobTitle:    "Engineer",
					Location:	"Delhi",
			},
			seedData: map[string]User{
				"item1": {
					Id:        "t2OUx6lvTMedrAiW2ffURA",
					Email:   "ui17co14@iiitsurat.ac.in",
					FirstName: "ekansh",
					LastName:  "singh",
					Type:        1,
					RoleName:    "Member",
					Status  :    "active",
					Department:  "devops",
					JobTitle:    "Engineer",
					Location:	"Delhi",
				},
			},
			expectErr: false,
		},
		{
			testName: "item does not exist",
			updatedItem: &User{
				Id :       "dfhjjddfjsd",
				Email:   "ui17ec38@iitsurat.ac.in",
				FirstName: "ekansh",
				LastName:  "rock",
				Type:        1,
				RoleName:    "Member",
				Status  :    "active",
				Department:  "devops",
				JobTitle:    "Engineer",
				Location:	"Delhi",
			},
			seedData:  nil,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("ZOOM_TOKEN"))
			err := client.UpdateItem(tc.updatedItem)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			item, err := client.GetItem(tc.updatedItem.Email)
			assert.NoError(t, err)
			assert.Equal(t, tc.updatedItem, item)
		})
	}
}

func TestClient_DeleteItem(t *testing.T) {
	testCases := []struct {
		testName  string
		itemName  string
		seedData  map[string]User
		expectErr bool
	}{
		{
			testName: "user exists",
			itemName: "ui17co14@iiitsurat.ac.in",
			seedData: map[string]User{
				"user1": {
					Id:        "t2OUx6lvTMedrAiW2ffURA",
					Email:   "ui17co14@iiitsurat.ac.in",
					FirstName: "ekansh",
					LastName:  "singh",
					Type:        1,
					RoleName:    "Member",
					Status  :    "active",
					Department:  "devops",
					JobTitle:    "Engineer",
					Location:	"Delhi",
				},
			},
			expectErr: false,
		},
		
		
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("ZOOM_TOKEN"))
			err := client.DeleteItem(tc.itemName)
			log.Println(err)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			_, err = client.GetItem(tc.itemName)
			assert.Error(t, err)
		})
	}
}



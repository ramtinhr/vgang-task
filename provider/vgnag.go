package provider

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/google/uuid"
	"github.com/ramtinhr/vgang-task/models"
	"github.com/ramtinhr/vgang-task/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type VgangUser struct {
	Username string
	Password string
}

type d struct {
	Data authData
}

type authData struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type product struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
}

type prodD struct {
	Products []product `json:"products"`
}

type catId uint

var (
	ClothingCat = catId(1)
)


// Login
// login to vgang account and create a indexer to use access token to continue the process
func (u *VgangUser) Login() (*models.Indexer, error) {
	url := fmt.Sprintf("%s/auth/login/retailer/vgang", os.Getenv("VGANG_API_ENDPOINT"))
	uuid := uuid.New()
	uuidBytes := uuid[:]
	licenseKeyBytes := append(uuidBytes)
	licenseKey := base64.StdEncoding.EncodeToString(licenseKeyBytes)
	data := map[string]string{
		"email":    u.Username,
		"password": u.Password,
		"deviceID": licenseKey,
	}

	jData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(jData))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respData d
	if err = json.Unmarshal(body, &respData); err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		pass, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
		if err != nil {
			return nil, err
		}

		var indexer = &models.Indexer{
			Username:     u.Username,
			Password:     string(pass),
			AccessToken:  respData.Data.AccessToken,
			RefreshToken: respData.Data.RefreshToken,
		}

		models.AddIndexer(indexer)

		logrus.Info("Indexer created successfully ")
		return indexer, err
	}

	return nil, errors.New("something went wrong")
}


// FetchProducts
// fetch products of specific category and save it with special hash in the database
func (u *VgangUser) FetchProducts(accessToken string) error {
	url, err := url.Parse(fmt.Sprintf("%s/retailers/products", os.Getenv("VGANG_API_ENDPOINT")))
	if err != nil {
		return err
	}

	url.RawQuery = fmt.Sprintf("count=100&offset=0&category=%v&sort=Latest&dont_show_out_of_stock=1", ClothingCat)
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var respData prodD
	if err = json.Unmarshal(body, &respData); err != nil {
		return err
	}

	var products []*models.Product

	if resp.StatusCode == http.StatusOK {
		for _, prod := range respData.Products {
			products = append(products, &models.Product{
				ProductID: prod.Id,
				Hash:      utils.RandomURL(8),
			})
		}

		err := models.AddProducts(products)
		if err != nil {
			return err
		}

		return nil
	}

	return errors.New("something went wrong")
}

package http

import (
	"encoding/json"
	"fmt"
	"github.com/yukselcodingwithyou/gocoingecko/domain"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var baseAddress = "https://api.coingecko.com/api/v3"

type Client struct {
	httpClient *http.Client
}

func New(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{httpClient: httpClient}
}

func doRequest(req *http.Request, client *http.Client) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

func (c *Client) get(address string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, address, nil)

	if err != nil {
		return nil, err
	}
	resp, err := doRequest(req, c.httpClient)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// API

// Ping /ping endpoint
func (c *Client) Ping() (*domain.Ping, error) {
	address := fmt.Sprintf("%s/ping", baseAddress)
	resp, err := c.get(address)
	if err != nil {
		return nil, err
	}
	var data *domain.Ping
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// SimpleSinglePrice /simple/price  Single ID and Currency (ids, vs_currency)
func (c *Client) SimpleSinglePrice(id string, vsCurrency string) (*domain.SimpleSinglePrice, error) {
	idParam := []string{strings.ToLower(id)}
	vcParam := []string{strings.ToLower(vsCurrency)}

	t, err := c.SimplePrice(idParam, vcParam)
	if err != nil {
		return nil, err
	}
	curr := (*t)[id]
	if len(curr) == 0 {
		return nil, fmt.Errorf("id or vsCurrency not existed")
	}
	data := &domain.SimpleSinglePrice{ID: id, Currency: vsCurrency, MarketPrice: curr[vsCurrency]}
	return data, nil
}

// SimplePrice /simple/price Multiple ID and Currency (ids, vs_currencies)
func (c *Client) SimplePrice(ids []string, vsCurrencies []string) (*map[string]map[string]float32, error) {
	params := url.Values{}
	idsParam := strings.Join(ids[:], ",")
	vsCurrenciesParam := strings.Join(vsCurrencies[:], ",")

	params.Add("ids", idsParam)
	params.Add("vs_currencies", vsCurrenciesParam)

	address := fmt.Sprintf("%s/simple/price?%s", baseAddress, params.Encode())
	resp, err := c.get(address)
	if err != nil {
		return nil, err
	}

	t := make(map[string]map[string]float32)
	err = json.Unmarshal(resp, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// SimpleSupportedVSCurrencies /simple/supported_vs_currencies
func (c *Client) SimpleSupportedVSCurrencies() (*domain.SimpleSupportedVSCurrencies, error) {
	address := fmt.Sprintf("%s/simple/supported_vs_currencies", baseAddress)
	resp, err := c.get(address)
	if err != nil {
		return nil, err
	}
	var data *domain.SimpleSupportedVSCurrencies
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CoinsList /coins/list
func (c *Client) CoinsList() (*domain.CoinList, error) {
	address := fmt.Sprintf("%s/coins/list", baseAddress)
	resp, err := c.get(address)
	if err != nil {
		return nil, err
	}

	var data *domain.CoinList
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CoinsMarket /coins/market
func (c *Client) CoinsMarket(vsCurrency string, ids []string, order string, perPage int, page int, sparkline bool, priceChangePercentage []string) (*domain.CoinsMarket, error) {
	if len(vsCurrency) == 0 {
		return nil, fmt.Errorf("vs_currency is required")
	}
	params := url.Values{}
	// vs_currency
	params.Add("vs_currency", vsCurrency)
	// order
	if len(order) == 0 {
		order = domain.OrderTypeObject.MarketCapDesc
	}
	params.Add("order", order)
	// ids
	if len(ids) != 0 {
		idsParam := strings.Join(ids[:], ",")
		params.Add("ids", idsParam)
	}
	// per_page
	if perPage <= 0 || perPage > 250 {
		perPage = 100
	}
	params.Add("per_page", fmt.Sprintf("%v", perPage))
	params.Add("page", fmt.Sprintf("%v", page))
	// sparkline
	params.Add("sparkline", fmt.Sprintf("%v", sparkline))
	// price_change_percentage
	if len(priceChangePercentage) != 0 {
		priceChangePercentageParam := strings.Join(priceChangePercentage[:], ",")
		params.Add("price_change_percentage", priceChangePercentageParam)
	}
	address := fmt.Sprintf("%s/coins/markets?%s", baseAddress, params.Encode())
	resp, err := c.get(address)
	if err != nil {
		return nil, err
	}
	var data *domain.CoinsMarket
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CoinsID /coins/{id}
func (c *Client) CoinsID(id string, localization bool, tickers bool, marketData bool, communityData bool, developerData bool, sparkline bool) (*domain.CoinsID, error) {

	if len(id) == 0 {
		return nil, fmt.Errorf("id is required")
	}
	params := url.Values{}
	params.Add("localization", fmt.Sprintf("%v", localization))
	params.Add("tickers", fmt.Sprintf("%v", tickers))
	params.Add("market_data", fmt.Sprintf("%v", marketData))
	params.Add("community_data", fmt.Sprintf("%v", communityData))
	params.Add("developer_data", fmt.Sprintf("%v", developerData))
	params.Add("sparkline", fmt.Sprintf("%v", sparkline))
	address := fmt.Sprintf("%s/coins/%s?%s", baseAddress, id, params.Encode())
	resp, err := c.get(address)
	if err != nil {
		return nil, err
	}

	var data *domain.CoinsID
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CoinsIDTickers /coins/{id}/tickers
func (c *Client) CoinsIDTickers(id string, page int) (*domain.CoinsIDTickers, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("id is required")
	}
	params := url.Values{}
	if page > 0 {
		params.Add("page", fmt.Sprintf("%v", page))
	}
	address := fmt.Sprintf("%s/coins/%s/tickers?%s", baseAddress, id, params.Encode())
	resp, err := c.get(address)
	if err != nil {
		return nil, err
	}
	var data *domain.CoinsIDTickers
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CoinsIDHistory /coins/{id}/history?date={date}&localization=false
func (c *Client) CoinsIDHistory(id string, date string, localization bool) (*domain.CoinsIDHistory, error) {
	if len(id) == 0 || len(date) == 0 {
		return nil, fmt.Errorf("id and date is required")
	}
	params := url.Values{}
	params.Add("date", date)
	params.Add("localization", fmt.Sprintf("%v", localization))

	address := fmt.Sprintf("%s/coins/%s/history?%s", baseAddress, id, params.Encode())
	resp, err := c.get(address)
	if err != nil {
		return nil, err
	}
	var data *domain.CoinsIDHistory
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CoinsIDMarketChart /coins/{id}/market_chart?vs_currency={usd, eur, jpy, etc.}&days={1,14,30,max}
func (c *Client) CoinsIDMarketChart(id string, vsCurrency string, days string) (*domain.CoinsIDMarketChart, error) {
	if len(id) == 0 || len(vsCurrency) == 0 || len(days) == 0 {
		return nil, fmt.Errorf("id, vs_currency, and days is required")
	}

	params := url.Values{}
	params.Add("vs_currency", vsCurrency)
	params.Add("days", days)

	address := fmt.Sprintf("%s/coins/%s/market_chart?%s", baseAddress, id, params.Encode())
	resp, err := c.get(address)
	if err != nil {
		return nil, err
	}

	m := domain.CoinsIDMarketChart{}
	err = json.Unmarshal(resp, &m)
	if err != nil {
		return &m, err
	}

	return &m, nil
}

// EventsCountries https://api.coingecko.com/api/v3/events/countries
func (c *Client) EventsCountries() ([]domain.EventCountryItem, error) {
	address := fmt.Sprintf("%s/events/countries", baseAddress)
	resp, err := c.get(address)
	if err != nil {
		return nil, err
	}
	var data *domain.EventsCountries
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data.Data, nil

}

// ExchangeRates https://api.coingecko.com/api/v3/exchange_rates
func (c *Client) ExchangeRates() (*domain.ExchangeRatesItem, error) {
	address := fmt.Sprintf("%s/exchange_rates", baseAddress)
	resp, err := c.get(address)
	if err != nil {
		return nil, err
	}
	var data *domain.ExchangeRatesResponse
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data.Rates, nil
}

// Global https://api.coingecko.com/api/v3/global
func (c *Client) Global() (*domain.Global, error) {
	address := fmt.Sprintf("%s/global", baseAddress)
	resp, err := c.get(address)
	if err != nil {
		return nil, err
	}
	var data *domain.GlobalResponse
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data.Data, nil
}

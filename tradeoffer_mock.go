package steamapi

func getMockActiveStateIEconGetTradeOffer() string {
	return `
    {
      "response": {
        "offer": {
          "tradeofferid": "123456",
          "accountid_other": 1234,
          "message": "",
          "expiration_time": 1300000000,
          "trade_offer_state": 2,
          "items_to_receive": [
            {
              "appid": "123",
              "contextid": "2",
              "assetid": "1234553",
              "classid": "888881",
              "instanceid": "0",
              "amount": "1",
              "missing": false
            }
          ],
          "is_our_offer": true,
          "time_created": 1300000000,
          "time_updated": 1300000000,
          "from_real_time_trade": false
        },
        "descriptions": [
          {
            "appid": 123,
            "classid": "888881",
            "instanceid": "0",
            "currency": false,
            "background_color": "",
            "icon_url": "454309584309539klfksdjflksdjf",
            "icon_url_large": "09089dklngflndfmd",
            "descriptions": [
              {
                "type": "html",
                "value": "ttttt"
              },
              {
                "type": "html",
                "value": " "
              },
              {
                "type": "html",
                "value": "ttttt."
              },
              {
                "type": "html",
                "value": " "
              },
              {
                "type": "html",
                "value": "test",
                "color": "EEEE"
              },
              {
                "type": "html",
                "value": " "
              }
            ],
            "tradable": true,
            "actions": [
              {
                "link": "lnk",
                "name": ""
              }
            ],
            "name": "testtest",
            "name_color": "DDDD",
            "type": "test",
            "market_name": "testmkt",
            "market_hash_name": "testmkt",
            "market_actions": [
              {
                "link": "lnk",
                "name": ""
              }
            ],
            "commodity": false,
            "market_tradable_restriction": 0
          }
        ]
      }
    }`
}

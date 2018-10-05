# PhoneNumber

Phone numbers may be in any of the following formats:

- 333-333-3333
- (333) 333-3333
- 1-333-333-3333
- 333.333.3333
- 333-333-3333
- 333-333-3333 x3333
- (333) 333-3333 x3333
- 1-333-333-3333 x3333
- 333.333.3333 x3333

## Examples

```
faker.RandomCellPhone() => "(168) 454-2055"
faker.RandomPhoneNumber() => "397.693.1309"

// For Canada
faker.SetLocale("en-CA")

// For the 'US only' methods below, first you must do the following:
faker.SetLocale("en-US")

// US only
faker.Phone().AreaCode() => "201"

// US only
faker.Phone().ExchangeCode() => "208"

// length=4
faker.Phone().SubscriberNumber() => "3873"

// Optional parameter: length=2
faker.Phone().FixedSubscriberNumber(2) => "39"

faker.Phone().Extention() => "3764"
```

## support locales

support format TBD list:
- bg
- ca-CAT
- ca
- da-DK
- de-AT
- de-CH
- de
- ee
- en
- en-AU
- en-BORK
- en-CA
- en-GB
- en-IND
- en-MS
- en-NEP
- en-NG
- en-NZ
- en-PAK
- en-SG
- en-UG
- en-US
- en-ZA
- en-au-ocker
- en
- es-MX
- es
- fa
- fi-FI
- fr-CA
- fr-CH
- fr
- he
- id
- it
- ja
- ko
- lv
- nb-NO
- nl
- no
- pl
- pt-BR
- pt
- ru
- sk
- sv
- tr
- uk
- vi
- zh-CN
- zh-TW

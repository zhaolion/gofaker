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
gofaker.RandomCellPhone() => "(168) 454-2055"
gofaker.RandomPhoneNumber() => "397.693.1309"

// For Canada
gofaker.SetLocale("en-CA")

// For the 'US only' methods below, first you must do the following:
gofaker.SetLocale("en-US")

// US only
gofaker.Phone().AreaCode() => "201"

// US only
gofaker.Phone().ExchangeCode() => "208"

// length=4
gofaker.Phone().SubscriberNumber() => "3873"

// Optional parameter: length=2
gofaker.Phone().FixedSubscriberNumber(2) => "39"

gofaker.Phone().Extention() => "3764"
```

## support locales

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

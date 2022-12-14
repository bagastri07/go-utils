# go-utils

<a name="v1.29.0"></a>
## v1.29.0 - 2022-12-14
### Code Improvements
- generate the private key only once ([#25](https://github.com/kumparan/kumnats/issues/25))
- rename func from GetStoryIDFromStorySlug to GetIDFromSlug
- JoinURL to be more versatile and should be able to accept variadic variables ([#9](https://github.com/kumparan/kumnats/issues/9))

### Fixes
- fix marshal issue on gorm.DeletedAt empty value ([#32](https://github.com/kumparan/kumnats/issues/32))
- should handle error when marshal/unmarshal from gqlgen ([#29](https://github.com/kumparan/kumnats/issues/29))
- handle index out of bound ([#28](https://github.com/kumparan/kumnats/issues/28))
- add null string & handle unmarshal "null" values ([#27](https://github.com/kumparan/kumnats/issues/27))

### Fixes
- gqlgen cannot resolve the package ([#21](https://github.com/kumparan/kumnats/issues/21))
- missing RetryStopper constructor ([#6](https://github.com/kumparan/kumnats/issues/6))

### New Features
- add GitHub Action to run golang-ci lint, unit test, and go build
- add PaginateSlice ([#44](https://github.com/kumparan/kumnats/issues/44))
- add delete by value in slice
- add GetCronNextAt
- create ValueOrDefault function
- escape quote
- nanoseconds check on rfc3339nano time formatting ([#38](https://github.com/kumparan/kumnats/issues/38))
- parse duration with default ([#36](https://github.com/kumparan/kumnats/issues/36))
- add int64 pointer to int64 conversion ([#33](https://github.com/kumparan/kumnats/issues/33))
- generate ULID from time
- add some functions using go generic ([#34](https://github.com/kumparan/kumnats/issues/34))
- add constraint size gql directive ([#30](https://github.com/kumparan/kumnats/issues/30))
- add gqlgen NullInt64, NullInt64ID & NullTime ([#26](https://github.com/kumparan/kumnats/issues/26))
- add custom time for gqlgen ([#22](https://github.com/kumparan/kumnats/issues/22))
- add AESCryptor ([#23](https://github.com/kumparan/kumnats/issues/23))
- logrus sentry hook using sentry-go ([#18](https://github.com/kumparan/kumnats/issues/18))
- add custom scalara gqlgen Int64ID ([#20](https://github.com/kumparan/kumnats/issues/20))
- add GetDifferenceDaysForHumans ([#19](https://github.com/kumparan/kumnats/issues/19))
- strip HTML ([#16](https://github.com/kumparan/kumnats/issues/16))
- add int64 millis to time converter
- add string millis to time converter
- add money formatter for multiple currencies ([#13](https://github.com/kumparan/kumnats/issues/13))
- add formatter for indonesian money and date
- add function to truncate string by length ([#11](https://github.com/kumparan/kumnats/issues/11))
- get story id from story slug
- create GenerateRandomAlphanumeric
- add lower map key
- new method to know the caller of the method in the runtime
- generate media URL
- init go-utils

### Other Improvements
- upgrade bluemonday
- fix dependabot alert ([#37](https://github.com/kumparan/kumnats/issues/37))


[Unreleased]: https://github.com/kumparan/kumnats/compare/v1.29.0...HEAD
